package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/umee-network/umee/v5/x/metoken"

	"github.com/umee-network/umee/v5/x/leverage/types"
)

// TestKeeper is a keeper with some normally
// unexported methods exposed for testing.
type TestKeeper struct {
	*Keeper
}

// NewTestKeeper returns a new leverage keeper, and
// an additional TestKeeper that exposes normally
// unexported methods for testing.
func NewTestKeeper(
	cdc codec.Codec,
	storeKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	bk types.BankKeeper,
	ok types.OracleKeeper,
	enableLiquidatorQuery bool,
) (Keeper, TestKeeper) {
	k := NewKeeper(
		cdc,
		storeKey,
		paramSpace,
		bk,
		ok,
		enableLiquidatorQuery,
		authtypes.NewModuleAddress(metoken.ModuleName),
	)
	return k, TestKeeper{&k}
}

func (tk *TestKeeper) SetBadDebtAddress(ctx sdk.Context, addr sdk.AccAddress, denom string, hasDebt bool) error {
	return tk.Keeper.setBadDebtAddress(ctx, addr, denom, hasDebt)
}

func (tk *TestKeeper) SetBorrow(ctx sdk.Context, addr sdk.AccAddress, amount sdk.Coin) error {
	return tk.Keeper.setBorrow(ctx, addr, amount)
}

func (tk *TestKeeper) SetCollateral(ctx sdk.Context, addr sdk.AccAddress, collateral sdk.Coin) error {
	return tk.Keeper.setCollateral(ctx, addr, collateral)
}

func (tk *TestKeeper) SetInterestScalar(ctx sdk.Context, denom string, scalar sdk.Dec) error {
	return tk.Keeper.setInterestScalar(ctx, denom, scalar)
}

func (tk *TestKeeper) SetReserveAmount(ctx sdk.Context, coin sdk.Coin) error {
	return tk.Keeper.setReserves(ctx, coin)
}
