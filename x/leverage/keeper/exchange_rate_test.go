package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	umeeapp "github.com/umee-network/umee/v2/app"
)

func (s *IntegrationTestSuite) TestDeriveExchangeRate() {
	// creates account which has supplied and collateralized 1000 uumee
	addr := s.newAccount(coin(umeeDenom, 1000))
	s.supply(addr, coin(umeeDenom, 1000))
	s.collateralize(addr, coin("u/"+umeeDenom, 1000))

	// artificially increase total borrows (by affecting a single address)
	err := s.tk.SetBorrow(s.ctx, addr, coin(umeeDenom, 2000))
	s.Require().NoError(err)

	// artificially set reserves
	err = s.tk.SetReserveAmount(s.ctx, coin(umeeDenom, 300))
	s.Require().NoError(err)

	// expected token:uToken exchange rate
	//    = (total borrows + module balance - reserves) / utoken supply
	//    = 2000 + 1000 - 300 / 1000
	//    = 2.7

	// get derived exchange rate
	rate := s.app.LeverageKeeper.DeriveExchangeRate(s.ctx, umeeapp.BondDenom)
	s.Require().Equal(sdk.MustNewDecFromStr("2.7"), rate)
}
