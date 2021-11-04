package e2e

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *IntegrationTestSuite) TestUmeeTokenTransfers() {
	// // deploy umee ERC20 token contract
	// var umeeERC20Addr string
	// s.Run("deploy_umee_erc20", func() {
	// 	umeeERC20Addr = s.deployERC20Token("uumee")
	// 	s.Require().NotEmpty(umeeERC20Addr)

	// 	_, err := hexutil.Decode(umeeERC20Addr)
	// 	s.Require().NoError(err)
	// })

	// // send 300 umee tokens from Umee to Ethereum
	// s.Run("send_uumee_tokens_to_eth", func() {
	// 	ethRecipient := s.chain.validators[1].ethereumKey.address
	// 	s.sendFromUmeeToEth(0, ethRecipient, "300uumee", "10photon", "7uumee")

	// 	endpoint := fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))
	// 	fromAddr := s.chain.validators[0].keyInfo.GetAddress()

	// 	balance, err := queryUmeeDenomBalance(endpoint, fromAddr.String(), "uumee")
	// 	s.Require().NoError(err)
	// 	s.Require().Equal(9999999693, balance)

	// 	expEthBalance := 300

	// 	// require the Ethereum recipient balance increased
	// 	s.Require().Eventually(
	// 		func() bool {
	// 			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// 			defer cancel()

	// 			b, err := queryEthTokenBalance(ctx, s.ethClient, umeeERC20Addr, ethRecipient)
	// 			if err != nil {
	// 				return false
	// 			}

	// 			return b == expEthBalance
	// 		},
	// 		2*time.Minute,
	// 		5*time.Second,
	// 	)
	// })

	// // send 300 umee tokens from Ethereum back to Umee
	// s.Run("send_uumee_tokens_from_eth", func() {
	// 	s.sendFromEthToUmee(1, umeeERC20Addr, s.chain.validators[0].keyInfo.GetAddress().String(), "300")

	// 	umeeEndpoint := fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))
	// 	toAddr := s.chain.validators[0].keyInfo.GetAddress()
	// 	expBalance := 9999999993

	// 	// require the original sender's (validator) balance increased
	// 	s.Require().Eventually(
	// 		func() bool {
	// 			b, err := queryUmeeDenomBalance(umeeEndpoint, toAddr.String(), "uumee")
	// 			if err != nil {
	// 				return false
	// 			}

	// 			return b == expBalance
	// 		},
	// 		2*time.Minute,
	// 		5*time.Second,
	// 	)
	// })
}

func (s *IntegrationTestSuite) TestIBCTokenTransfer() {
	var ibcStakeDenom string
	s.Run("send_stake_to_umee", func() {
		recipient := s.chain.validators[0].keyInfo.GetAddress().String()
		token := sdk.NewInt64Coin("stake", 3300000000) // 3300stake
		s.sendIBC(gaiaChainID, s.chain.id, recipient, token)

		umeeAPIEndpoint := fmt.Sprintf("http://%s", s.valResources[0].GetHostPort("1317/tcp"))

		// require the recipient account receives the IBC tokens (IBC packets ACKd)
		var (
			balances sdk.Coins
			err      error
		)
		s.Require().Eventually(
			func() bool {
				balances, err = queryUmeeAllBalances(umeeAPIEndpoint, recipient)
				s.Require().NoError(err)

				return balances.Len() == 3
			},
			time.Minute,
			5*time.Second,
		)

		for _, c := range balances {
			if strings.Contains(c.Denom, "ibc/") {
				ibcStakeDenom = c.Denom
				break
			}
		}

		s.Require().NotEmpty(ibcStakeDenom)
	})
}
