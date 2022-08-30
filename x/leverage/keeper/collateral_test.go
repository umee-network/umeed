package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v3/x/leverage/types"
)

func (s *IntegrationTestSuite) TestGetCollateralAmount() {
	ctx, require := s.ctx, s.Require()
	uDenom := types.ToUTokenDenom(umeeDenom)

	// get u/umee collateral amount of empty account address
	collateral := s.tk.GetCollateralAmount(ctx, sdk.AccAddress{}, uDenom)
	require.Equal(coin(uDenom, 0), collateral)

	// fund an account
	addr := s.newAccount(coin(umeeDenom, 1000))

	// get u/umee collateral amount of non-empty account address
	collateral = s.tk.GetCollateralAmount(ctx, addr, uDenom)
	require.Equal(coin(uDenom, 0), collateral)

	// supply 1000 u/uumee but do not collateralize
	s.supply(addr, coin(umeeDenom, 1000))

	// confirm collateral amount is 0 u/uumee
	collateral = s.tk.GetCollateralAmount(ctx, addr, uDenom)
	require.Equal(coin(uDenom, 0), collateral)

	// enable u/umee as collateral
	s.collateralize(addr, coin(uDenom, 1000))

	// confirm collateral amount is 1000 u/uumee
	collateral = s.tk.GetCollateralAmount(ctx, addr, uDenom)
	require.Equal(coin(uDenom, 1000), collateral)

	// collateral amount of non-utoken denom (zero)
	collateral = s.tk.GetCollateralAmount(ctx, addr, umeeDenom)
	require.Equal(coin(umeeDenom, 0), collateral)

	// collateral amount of unregistered denom (zero)
	collateral = s.tk.GetCollateralAmount(ctx, addr, "abcd")
	require.Equal(coin("abcd", 0), collateral)

	// disable u/umee as collateral
	s.decollateralize(addr, coin(uDenom, 1000))

	// confirm collateral amount is 0 u/uumee
	collateral = s.tk.GetCollateralAmount(ctx, addr, uDenom)
	require.Equal(coin(uDenom, 0), collateral)

	// we do not test empty denom, as that will cause a panic
}

func (s *IntegrationTestSuite) TestSetCollateralAmount() {
	ctx, require := s.ctx, s.Require()
	uDenom := types.ToUTokenDenom(umeeDenom)

	// set u/umee collateral amount of empty account address (error)
	err := s.tk.SetCollateralAmount(ctx, sdk.AccAddress{}, sdk.NewInt64Coin(uDenom, 0))
	require.EqualError(err, "empty address")

	addr := s.newAccount()

	// force invalid denom
	err = s.tk.SetCollateralAmount(ctx, addr, sdk.Coin{Denom: "", Amount: sdk.ZeroInt()})
	require.ErrorContains(err, "invalid denom")

	// set u/umee collateral amount
	require.NoError(s.tk.SetCollateralAmount(ctx, addr, sdk.NewInt64Coin(uDenom, 10)))

	// confirm effect
	collateral := s.tk.GetCollateralAmount(ctx, addr, uDenom)
	require.Equal(sdk.NewInt64Coin(uDenom, 10), collateral)

	// set u/umee collateral amount to zero
	require.NoError(s.tk.SetCollateralAmount(ctx, addr, sdk.NewInt64Coin(uDenom, 0)))

	// confirm effect
	collateral = s.tk.GetCollateralAmount(ctx, addr, uDenom)
	require.Equal(sdk.NewInt64Coin(uDenom, 0), collateral)

	// set unregistered token collateral amount
	require.NoError(s.tk.SetCollateralAmount(ctx, addr, sdk.NewInt64Coin("abcd", 10)))

	// confirm effect
	collateral = s.tk.GetCollateralAmount(ctx, addr, "abcd")
	require.Equal(sdk.NewInt64Coin("abcd", 10), collateral)

	// set unregistered token collateral amount to zero
	require.NoError(s.tk.SetCollateralAmount(ctx, addr, sdk.NewInt64Coin("abcd", 0)))

	// confirm effect
	collateral = s.tk.GetCollateralAmount(ctx, addr, "abcd")
	require.Equal(sdk.NewInt64Coin("abcd", 0), collateral)

	// we do not test empty denom, as that will cause a panic
}

func (s *IntegrationTestSuite) TestTotalCollateral() {
	app, ctx, require := s.app, s.ctx, s.Require()

	// Test zero collateral
	uDenom := types.ToUTokenDenom(umeeDenom)
	collateral := app.LeverageKeeper.GetTotalCollateral(ctx, uDenom)
	require.Equal(sdk.ZeroInt(), collateral, "zero collateral")

	// create a supplier which will have 100 u/UMEE collateral
	addr := s.newAccount(coin(umeeDenom, 100_000000))
	s.supply(addr, coin(umeeDenom, 100_000000))
	s.collateralize(addr, coin(uDenom, 100_000000))

	// Test nonzero collateral
	collateral = app.LeverageKeeper.GetTotalCollateral(ctx, uDenom)
	require.Equal(sdk.NewInt(100_000000), collateral, "nonzero collateral")
}
