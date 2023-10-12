package e2e

import (
	"fmt"
	"time"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	appparams "github.com/umee-network/umee/v6/app/params"
	setup "github.com/umee-network/umee/v6/tests/e2e/setup"
	"github.com/umee-network/umee/v6/tests/grpc"
	"github.com/umee-network/umee/v6/x/uibc"
)

const (
	// ibc hash of gaia stake token on umee
	stakeIBCHash = "ibc/C053D637CCA2A2BA030E2C5EE1B28A16F71CCB0E45E8BE52766DC1B241B77878"
	// ibc hash of uumee token on gaia
	umeeIBCHash = "ibc/9F53D255F5320A4BE124FF20C29D46406E126CE8A09B00CA8D3CFF7905119728"
	// ibc hash of uatom token on umee
	uatomIBCHash = "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
)

var powerReduction = sdk.MustNewDecFromStr("10").Power(6)

func (s *E2ETest) checkOutflowByPercentage(endpoint, excDenom string, outflow, amount, perDiff sdk.Dec) {
	// get historic average price for denom (SYMBOL_DENOM)
	histoAvgPrice, err := s.QueryHistAvgPrice(endpoint, excDenom)
	s.Require().NoError(err)
	totalPrice := amount.Quo(powerReduction).Mul(histoAvgPrice)
	s.T().Log("exchangeRate total price ", totalPrice.String(), "outflow value", outflow.String())
	percentageDiff := totalPrice.Mul(perDiff)
	// Note: checking outflow >= total_price with percentageDiff
	// either total_price >= outflow with percentageDiff
	s.Require().True(outflow.GTE(totalPrice.Sub(percentageDiff)) || totalPrice.GTE(outflow.Sub(percentageDiff)))
}

func (s *E2ETest) checkOutflows(umeeAPIEndpoint, denom string, checkWithExcRate bool, amount sdk.Dec, excDenom string) {
	s.Require().Eventually(
		func() bool {
			a, err := s.QueryOutflows(umeeAPIEndpoint, denom)
			if err != nil {
				return false
			}
			if checkWithExcRate {
				s.checkOutflowByPercentage(umeeAPIEndpoint, excDenom, a, amount, sdk.MustNewDecFromStr("0.01"))
			}
			return true
		},
		30*time.Second,
		500*time.Millisecond,
	)
}

func (s *E2ETest) checkSupply(endpoint, ibcDenom string, amount math.Int) {
	s.Require().Eventually(
		func() bool {
			supply, err := s.QueryTotalSupply(endpoint)
			if err != nil {
				return false
			}

			return supply.AmountOf(ibcDenom).Equal(amount)
		},
		2*time.Minute,
		time.Second,
		fmt.Sprintf("check supply: %s (expected %s)", ibcDenom, amount),
	)
}

func (s *E2ETest) TestIBCTokenTransfer() {
	// IBC inbound transfer of non x/leverage registered tokens must fail, because
	// because we won't have price for it.
	s.Run("send_stake_to_umee", func() {
		// require the recipient account receives the IBC tokens (IBC packets ACKd)
		umeeAPIEndpoint := s.UmeeREST()
		recipient := s.AccountAddr(0).String()

		token := sdk.NewInt64Coin("stake", 3300000000) // 3300stake
		s.SendIBC(setup.GaiaChainID, s.Chain.ID, recipient, token, false, "")
		// Zero, since not a registered token
		s.checkSupply(umeeAPIEndpoint, stakeIBCHash, sdk.ZeroInt())
	})

	s.Run("ibc_transfer_quota", func() {
		// require the recipient account receives the IBC tokens (IBC packets ACKd)
		gaiaAPIEndpoint := s.GaiaREST()
		umeeAPIEndpoint := s.UmeeREST()
		atomSymbol := "ATOM"
		umeeSymbol := "UMEE"
		totalQuota := math.NewInt(120)
		tokenQuota := math.NewInt(100)

		// send uatom from gaia to umee
		// Note : gaia -> umee (ibc_quota will not check token limit)
		atomPrice, err := s.QueryHistAvgPrice(umeeAPIEndpoint, atomSymbol)
		s.Require().NoError(err)
		s.Require().True(atomPrice.GT(sdk.OneDec()),
			"atom price should be non zero, and expecting higher than 1, got: %s", atomPrice)
		emOfAtom := sdk.NewDecFromInt(totalQuota).Quo(atomPrice)
		c := sdk.NewInt64Coin("uatom", emOfAtom.Mul(powerReduction).RoundInt64())
		s.Require().True(c.Amount.GT(sdk.NewInt(2_000_000)),
			"amount should be non zero, and expecting much higher than 2 atom = 2e6 uatom, got: %s", c.Amount)
		s.SendIBC(setup.GaiaChainID, s.Chain.ID, "", c, false, "uatom from gaia to umee")
		s.checkSupply(umeeAPIEndpoint, uatomIBCHash, c.Amount)

		// sending more tokens than token_quota limit of umee (token_quota is 100$)
		histoAvgPriceOfUmee, err := s.QueryHistAvgPrice(umeeAPIEndpoint, umeeSymbol)
		s.Require().NoError(err)
		s.Require().True(histoAvgPriceOfUmee.GT(sdk.MustNewDecFromStr("0.001")),
			"umee price should be non zero, and expecting higher than 0.001, got: %s", histoAvgPriceOfUmee)
		exceedAmountOfUmee := sdk.NewDecFromInt(totalQuota).Quo(histoAvgPriceOfUmee)
		exceedAmountCoin := sdk.NewInt64Coin(appparams.BondDenom, exceedAmountOfUmee.Mul(powerReduction).RoundInt64())
		s.SendIBC(s.Chain.ID, setup.GaiaChainID, "", exceedAmountCoin, true, fmt.Sprintf(
			"sending %s amount %s more than %s", umeeSymbol, exceedAmountOfUmee.String(), totalQuota.String(),
		))
		// check the ibc (umee) quota after ibc txs
		s.checkSupply(gaiaAPIEndpoint, umeeIBCHash, math.ZeroInt())

		// send $90 UMEE from umee to gaia
		// Note: receiver is null means hermes will default send to key_name (from config) of target chain (gaia)
		// umee -> gaia (ibc_quota will check)
		umeeInitialQuota := math.NewInt(90)
		belowTokenQuota := sdk.NewDecFromInt(umeeInitialQuota).Quo(histoAvgPriceOfUmee)
		token := sdk.NewInt64Coin(appparams.BondDenom, belowTokenQuota.Mul(powerReduction).RoundInt64())
		s.SendIBC(s.Chain.ID, setup.GaiaChainID, "", token, false, fmt.Sprintf(
			"sending %s amount %s less than token quota %s", "UMEE", belowTokenQuota.String(), tokenQuota.String()))
		s.checkOutflows(umeeAPIEndpoint, appparams.BondDenom, true, sdk.NewDecFromInt(token.Amount), appparams.Name)
		s.checkSupply(gaiaAPIEndpoint, umeeIBCHash, token.Amount)

		// send uatom (ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2) from umee to gaia
		uatomIBCToken := sdk.NewInt64Coin(uatomIBCHash, c.Amount.Int64())
		// supply will be not be decreased because sending uatomIBCToken amount is more than token quota so it will fail
		s.SendIBC(s.Chain.ID, setup.GaiaChainID, "", uatomIBCToken, true, "uatom from umee to gaia")
		s.checkSupply(umeeAPIEndpoint, uatomIBCHash, uatomIBCToken.Amount)

		// send uatom below the token quota
		/*
			umee -> gaia
			umee token_quot = 90$
			total_quota = 120$
		*/
		belowTokenQuotabutNotBelowTotalQuota := sdk.NewDecFromInt(math.NewInt(90)).Quo(atomPrice)
		uatomIBCToken.Amount = math.NewInt(belowTokenQuotabutNotBelowTotalQuota.Mul(powerReduction).RoundInt64())
		s.SendIBC(s.Chain.ID, setup.GaiaChainID, "", uatomIBCToken, true, "below token quota but not total quota")
		// supply will be not be decreased because sending more than total quota from umee to gaia
		s.checkSupply(umeeAPIEndpoint, uatomIBCHash, c.Amount)
		// making sure below the total quota
		belowTokenQuotaInUSD := totalQuota.Sub(umeeInitialQuota).Sub(math.NewInt(2))
		belowTokenQuotaforAtom := sdk.NewDecFromInt(belowTokenQuotaInUSD).Quo(atomPrice)
		uatomIBCToken.Amount = math.NewInt(belowTokenQuotaforAtom.Mul(powerReduction).RoundInt64())
		s.SendIBC(s.Chain.ID, setup.GaiaChainID, "", uatomIBCToken, false, "below both quotas")
		// remaing supply still exists for uatom in umee
		s.checkSupply(umeeAPIEndpoint, uatomIBCHash, c.Amount.Sub(uatomIBCToken.Amount))
		s.checkOutflows(umeeAPIEndpoint, uatomIBCHash, true, sdk.NewDecFromInt(uatomIBCToken.Amount), atomSymbol)

		// sending more tokens then token_quota limit of umee
		s.SendIBC(s.Chain.ID, setup.GaiaChainID, "", exceedAmountCoin, true, "above umee quota")
		// check the ibc (umee) supply after ibc txs, it will same as previous because it will fail because to quota limit exceed
		s.checkSupply(gaiaAPIEndpoint, umeeIBCHash, token.Amount)

		/* sending back some amount from receiver to sender (ibc/XXX)
		gaia -> umee
		*/
		s.SendIBC(setup.GaiaChainID, s.Chain.ID, "", sdk.NewInt64Coin(umeeIBCHash, 1000), false, "send back some")
		s.checkSupply(gaiaAPIEndpoint, umeeIBCHash, token.Amount.Sub(math.NewInt(1000)))
		// sending back remaining ibc amount from receiver to sender (ibc/XXX)
		s.SendIBC(setup.GaiaChainID, s.Chain.ID, "",
			sdk.NewInt64Coin(umeeIBCHash, token.Amount.Sub(math.NewInt(1000)).Int64()), false, "send back remaining")
		s.checkSupply(gaiaAPIEndpoint, umeeIBCHash, math.ZeroInt())

		// reset the outflows
		s.T().Logf("waiting until quota reset, basically it will take around 300 seconds to do quota reset")
		s.Require().Eventually(
			func() bool {
				amount, err := s.QueryOutflows(umeeAPIEndpoint, appparams.BondDenom)
				if err != nil {
					return false
				}
				if amount.IsZero() {
					s.T().Logf("quota is reset : %s is 0", appparams.BondDenom)
					return true
				}
				return false
			},
			4*time.Minute,
			3*time.Second,
		)

		/****
			IBC_Status : disble (making ibc_transfer quota check disabled)
			No Outflows will updated
		***/
		// Make gov proposal to disable the quota check on ibc-transfer

		err = grpc.UIBCIBCTransferStatusUpdate(
			s.AccountClient(0),
			uibc.IBCTransferStatus_IBC_TRANSFER_STATUS_QUOTA_DISABLED,
		)
		s.Require().NoError(err)
		// Get the uibc params for quota checking
		uibcParams, err := s.AccountClient(0).QueryUIBCParams()
		s.Require().NoError(err)
		s.Require().Equal(uibcParams.IbcStatus, uibc.IBCTransferStatus_IBC_TRANSFER_STATUS_QUOTA_DISABLED)
		token = sdk.NewInt64Coin("uumee", 100000000) // 100 Umee

		// sending the umee tokens
		s.SendIBC(s.Chain.ID, setup.GaiaChainID, "", token, false, "sending umee")
		// Check the outflows
		s.checkSupply(gaiaAPIEndpoint, umeeIBCHash, token.Amount)
		s.Require().Eventually(
			func() bool {
				a, err := s.QueryOutflows(umeeAPIEndpoint, appparams.BondDenom)
				if err != nil {
					return false
				}
				return a.Equal(sdk.ZeroDec())
			},
			30*time.Second,
			1*time.Second,
		)
		// resend the umee token from gaia to umee
		s.SendIBC(setup.GaiaChainID, s.Chain.ID, "", sdk.NewInt64Coin(umeeIBCHash, token.Amount.Int64()), false, "resend")
		s.checkSupply(gaiaAPIEndpoint, umeeIBCHash, sdk.ZeroInt())
	})
}
