package simulation_test

import (
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	umeeapp "github.com/umee-network/umee/app"
	umeeappbeta "github.com/umee-network/umee/app/beta"
	"github.com/umee-network/umee/x/leverage"
	"github.com/umee-network/umee/x/leverage/simulation"
	"github.com/umee-network/umee/x/leverage/types"
)

// SimTestSuite wraps the test suite for running the simulations
type SimTestSuite struct {
	suite.Suite

	app *umeeappbeta.UmeeApp
	ctx sdk.Context
}

// SetupTest creates a new umee base app
func (s *SimTestSuite) SetupTest() {
	checkTx := false
	betaApp := umeeappbeta.Setup(s.T(), checkTx, 1)
	ctx := betaApp.NewContext(checkTx, tmproto.Header{})

	// Note: Setting umee collateral weight to 1.0 to allow lender to borrow heavily
	umeeToken := types.Token{
		BaseDenom:            umeeapp.BondDenom,
		ReserveFactor:        sdk.MustNewDecFromStr("0.25"),
		CollateralWeight:     sdk.MustNewDecFromStr("1.0"),
		BaseBorrowRate:       sdk.MustNewDecFromStr("0.02"),
		KinkBorrowRate:       sdk.MustNewDecFromStr("0.2"),
		MaxBorrowRate:        sdk.MustNewDecFromStr("1.0"),
		KinkUtilizationRate:  sdk.MustNewDecFromStr("0.8"),
		LiquidationIncentive: sdk.MustNewDecFromStr("0.1"),
	}
	atomIBCToken := types.Token{
		BaseDenom:            "ibc/CDC4587874B85BEA4FCEC3CEA5A1195139799A1FEE711A07D972537E18FDA39D",
		ReserveFactor:        sdk.MustNewDecFromStr("0.25"),
		CollateralWeight:     sdk.MustNewDecFromStr("0.1"),
		BaseBorrowRate:       sdk.MustNewDecFromStr("0.05"),
		KinkBorrowRate:       sdk.MustNewDecFromStr("0.3"),
		MaxBorrowRate:        sdk.MustNewDecFromStr("0.9"),
		KinkUtilizationRate:  sdk.MustNewDecFromStr("0.75"),
		LiquidationIncentive: sdk.MustNewDecFromStr("0.11"),
	}
	uabc := types.Token{
		BaseDenom:            "uabc",
		ReserveFactor:        sdk.MustNewDecFromStr("0"),
		CollateralWeight:     sdk.MustNewDecFromStr("0"),
		BaseBorrowRate:       sdk.MustNewDecFromStr("0.02"),
		KinkBorrowRate:       sdk.MustNewDecFromStr("0.22"),
		MaxBorrowRate:        sdk.MustNewDecFromStr("1.52"),
		KinkUtilizationRate:  sdk.MustNewDecFromStr("0.87"),
		LiquidationIncentive: sdk.MustNewDecFromStr("0.1"),
	}

	tokens := []types.Token{umeeToken, atomIBCToken, uabc}

	for _, token := range tokens {
		betaApp.LeverageKeeper.SetRegisteredToken(ctx, token)
		betaApp.LeverageKeeper.SetExchangeRate(ctx, token.BaseDenom, sdk.MustNewDecFromStr("1.1"))
	}

	betaApp.OracleKeeper.SetExchangeRate(ctx, umeeapp.DisplayDenom, sdk.OneDec())
	leverage.InitGenesis(ctx, betaApp.LeverageKeeper, *types.DefaultGenesis())

	s.app = betaApp
	s.ctx = ctx
}

// getTestingAccounts generates accounts with stake and umee balance
func (s *SimTestSuite) getTestingAccounts(r *rand.Rand, n int, cb func(fundedAccount simtypes.Account)) []simtypes.Account {
	accounts := simtypes.RandomAccounts(r, n)

	initAmt := sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
	bankCoins := sdk.NewCoins()
	accCoins := sdk.NewCoins()

	tokens, err := s.app.LeverageKeeper.GetAllRegisteredTokens(s.ctx)
	s.Require().NoError(err)

	for _, token := range tokens {
		bankCoins = bankCoins.Add(sdk.NewCoin(token.BaseDenom, initAmt))
		accCoins = accCoins.Add(sdk.NewCoin(token.BaseDenom, initAmt.Sub(sdk.NewInt(200))))
	}

	// add coins to the accounts
	for _, account := range accounts {
		acc := s.app.AccountKeeper.NewAccountWithAddress(s.ctx, account.Address)
		s.app.AccountKeeper.SetAccount(s.ctx, acc)
		s.Require().NoError(s.app.BankKeeper.MintCoins(s.ctx, minttypes.ModuleName, bankCoins))
		s.Require().NoError(s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, bankCoins))
		s.Require().NoError(
			s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, minttypes.ModuleName, acc.GetAddress(), accCoins),
		)
		cb(account)
	}

	return accounts
}

// TestWeightedOperations tests the weights of the operations.
func (s *SimTestSuite) TestWeightedOperations() {
	cdc := s.app.AppCodec()
	appParams := make(simtypes.AppParams)

	weightesOps := simulation.WeightedOperations(appParams, cdc, s.app.AccountKeeper, s.app.BankKeeper, s.app.LeverageKeeper)

	// setup 3 accounts
	r := rand.New(rand.NewSource(1))
	accs := s.getTestingAccounts(r, 3, func(acc simtypes.Account) {})

	expected := []struct {
		weight     int
		opMsgRoute string
		opMsgName  string
	}{
		{simulation.DefaultWeightMsgLendAsset, types.ModuleName, types.EventTypeLoanAsset},
		{simulation.DefaultWeightMsgWithdrawAsset, types.ModuleName, types.EventTypeWithdrawLoanedAsset},
		{simulation.DefaultWeightMsgBorrowAsset, types.ModuleName, types.EventTypeBorrowAsset},
		{simulation.DefaultWeightMsgSetCollateral, types.ModuleName, types.EventTypeSetCollateralSetting},
		{simulation.DefaultWeightMsgRepayAsset, types.ModuleName, types.EventTypeRepayBorrowedAsset},
		{simulation.DefaultWeightMsgLiquidate, types.ModuleName, types.EventTypeLiquidate},
	}

	for i, w := range weightesOps {
		operationMsg, _, _ := w.Op()(r, s.app.BaseApp, s.ctx, accs, "")
		// 	// the following checks are very much dependent from the ordering of the output given
		// 	// by WeightedOperations. if the ordering in WeightedOperations changes some tests
		// 	// will fail
		s.Require().Equal(expected[i].weight, w.Weight(), "weight should be the same")
		s.Require().Equal(expected[i].opMsgRoute, operationMsg.Route, "route should be the same")
		s.Require().Equal(expected[i].opMsgName, operationMsg.Name, "operation Msg name should be the same")
	}
}

func (s *SimTestSuite) TestSimulateMsgLendAsset() {
	r := rand.New(rand.NewSource(1))
	accs := s.getTestingAccounts(r, 3, func(fundedAccount simtypes.Account) {})

	s.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash}})

	op := simulation.SimulateMsgLendAsset(s.app.AccountKeeper, s.app.BankKeeper)
	operationMsg, futureOperations, err := op(r, s.app.BaseApp, s.ctx, accs, "")
	s.Require().NoError(err)

	var msg types.MsgLendAsset
	types.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	s.Require().True(operationMsg.OK)
	s.Require().Equal("umee1ghekyjucln7y67ntx7cf27m9dpuxxemn8w6h33", msg.Lender)
	s.Require().Equal(types.EventTypeLoanAsset, msg.Type())
	s.Require().Equal("185121068uumee", msg.Amount.String())
	s.Require().Len(futureOperations, 0)
}

func (s *SimTestSuite) TestSimulateMsgWithdrawAsset() {
	r := rand.New(rand.NewSource(1))
	lendValue := sdk.NewCoin(umeeapp.BondDenom, sdk.NewInt(100))

	accs := s.getTestingAccounts(r, 3, func(fundedAccount simtypes.Account) {
		uToken, err := s.app.LeverageKeeper.ExchangeToken(s.ctx, lendValue)
		if err != nil {
			s.Require().NoError(err)
		}

		s.app.LeverageKeeper.SetCollateralSetting(s.ctx, fundedAccount.Address, uToken.Denom, true)
		s.app.LeverageKeeper.LendAsset(s.ctx, fundedAccount.Address, lendValue)
	})

	s.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash}})

	op := simulation.SimulateMsgWithdrawAsset(s.app.AccountKeeper, s.app.BankKeeper, s.app.LeverageKeeper)
	operationMsg, futureOperations, err := op(r, s.app.BaseApp, s.ctx, accs, "")
	s.Require().NoError(err)

	var msg types.MsgWithdrawAsset
	types.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	s.Require().True(operationMsg.OK)
	s.Require().Equal("umee1ghekyjucln7y67ntx7cf27m9dpuxxemn8w6h33", msg.Lender)
	s.Require().Equal(types.EventTypeWithdrawLoanedAsset, msg.Type())
	s.Require().Equal("90u/uumee", msg.Amount.String())
	s.Require().Len(futureOperations, 0)
}

func (s *SimTestSuite) TestSimulateMsgBorrowAsset() {
	r := rand.New(rand.NewSource(8))
	lendValue := sdk.NewCoin(umeeapp.BondDenom, sdk.NewInt(100))

	accs := s.getTestingAccounts(r, 3, func(fundedAccount simtypes.Account) {
		uToken, err := s.app.LeverageKeeper.ExchangeToken(s.ctx, lendValue)
		if err != nil {
			s.Require().NoError(err)
		}

		s.app.LeverageKeeper.SetCollateralSetting(s.ctx, fundedAccount.Address, uToken.Denom, true)
		s.app.LeverageKeeper.LendAsset(s.ctx, fundedAccount.Address, lendValue)
	})

	s.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash}})

	op := simulation.SimulateMsgBorrowAsset(s.app.AccountKeeper, s.app.BankKeeper, s.app.LeverageKeeper)
	operationMsg, futureOperations, err := op(r, s.app.BaseApp, s.ctx, accs, "")
	s.Require().NoError(err)

	var msg types.MsgBorrowAsset
	types.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	s.Require().True(operationMsg.OK)
	s.Require().Equal("umee1qnclgkcxtuledc8xhle4lqly2q0z96uqkks60s", msg.Borrower)
	s.Require().Equal(types.EventTypeBorrowAsset, msg.Type())
	s.Require().Equal("67uumee", msg.Amount.String())
	s.Require().Len(futureOperations, 0)
}

func (s *SimTestSuite) TestSimulateMsgSetCollateralSetting() {
	r := rand.New(rand.NewSource(1))

	accs := s.getTestingAccounts(r, 3, func(fundedAccount simtypes.Account) {})

	s.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash}})

	op := simulation.SimulateMsgSetCollateralSetting(s.app.AccountKeeper, s.app.BankKeeper, s.app.LeverageKeeper)
	operationMsg, futureOperations, err := op(r, s.app.BaseApp, s.ctx, accs, "")
	s.Require().NoError(err)

	var msg types.MsgSetCollateral
	types.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	s.Require().True(operationMsg.OK)
	s.Require().Equal("umee1ghekyjucln7y67ntx7cf27m9dpuxxemn8w6h33", msg.Borrower)
	s.Require().Equal(types.EventTypeSetCollateralSetting, msg.Type())
	s.Require().Equal(true, msg.Enable)
	s.Require().Len(futureOperations, 0)
}

func (s *SimTestSuite) TestSimulateMsgRepayAsset() {
	r := rand.New(rand.NewSource(1))
	borrowValue := sdk.NewCoin(umeeapp.BondDenom, sdk.NewInt(190))

	accs := s.getTestingAccounts(r, 3, func(fundedAccount simtypes.Account) {
		uToken, err := s.app.LeverageKeeper.ExchangeToken(s.ctx, borrowValue)
		if err != nil {
			s.Require().NoError(err)
		}

		s.Require().NoError(s.app.LeverageKeeper.SetCollateralSetting(s.ctx, fundedAccount.Address, uToken.Denom, true))
		s.app.LeverageKeeper.LendAsset(s.ctx, fundedAccount.Address, borrowValue.AddAmount(sdk.NewInt(50)))
		s.Require().NoError(s.app.LeverageKeeper.BorrowAsset(s.ctx, fundedAccount.Address, borrowValue))
	})

	s.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash}})

	op := simulation.SimulateMsgRepayAsset(s.app.AccountKeeper, s.app.BankKeeper, s.app.LeverageKeeper)
	operationMsg, futureOperations, err := op(r, s.app.BaseApp, s.ctx, accs, "")
	s.Require().NoError(err)

	var msg types.MsgRepayAsset
	types.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	s.Require().True(operationMsg.OK)
	s.Require().Equal("umee1ghekyjucln7y67ntx7cf27m9dpuxxemn8w6h33", msg.Borrower)
	s.Require().Equal(types.EventTypeRepayBorrowedAsset, msg.Type())
	s.Require().Equal("190uumee", msg.Amount.String())
	s.Require().Len(futureOperations, 0)
}

func (s *SimTestSuite) TestSimulateMsgLiquidate() {
	r := rand.New(rand.NewSource(1))
	borrowValue := sdk.NewCoin(umeeapp.BondDenom, sdk.NewInt(190))

	accs := s.getTestingAccounts(r, 3, func(fundedAccount simtypes.Account) {
		uToken, err := s.app.LeverageKeeper.ExchangeToken(s.ctx, borrowValue)
		if err != nil {
			s.Require().NoError(err)
		}

		s.Require().NoError(s.app.LeverageKeeper.SetCollateralSetting(s.ctx, fundedAccount.Address, uToken.Denom, true))
		s.Require().NoError(s.app.LeverageKeeper.LendAsset(s.ctx, fundedAccount.Address, borrowValue))
		s.Require().NoError(s.app.LeverageKeeper.SetBorrow(s.ctx, fundedAccount.Address, borrowValue))
	})

	s.app.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: s.app.LastBlockHeight() + 1, AppHash: s.app.LastCommitID().Hash}})

	op := simulation.SimulateMsgLiquidate(s.app.AccountKeeper, s.app.BankKeeper, s.app.LeverageKeeper)
	operationMsg, futureOperations, err := op(r, s.app.BaseApp, s.ctx, accs, "")
	s.Require().NoError(err)

	var msg types.MsgLiquidate
	types.ModuleCdc.UnmarshalJSON(operationMsg.Msg, &msg)

	s.Require().True(operationMsg.OK)
	s.Require().Equal("umee1p8wcgrjr4pjju90xg6u9cgq55dxwq8j7wrm6ea", msg.Liquidator)
	s.Require().Equal("umee1p8wcgrjr4pjju90xg6u9cgq55dxwq8j7wrm6ea", msg.Borrower)
	s.Require().Equal("u/uumee", msg.RewardDenom)
	s.Require().Equal(types.EventTypeLiquidate, msg.Type())
	s.Require().Equal("190uumee", msg.Repayment.String())
	s.Require().Len(futureOperations, 0)
}

func TestSimTestSuite(t *testing.T) {
	suite.Run(t, new(SimTestSuite))
}
