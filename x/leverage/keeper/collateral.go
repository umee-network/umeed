package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/umee-network/umee/v2/x/leverage/types"
)

// liquidateCollateral burns utoken collateral and sends the base token reward to the liquidator.
// This occurs during direct liquidation.
func (k Keeper) liquidateCollateral(ctx sdk.Context, borrower, liquidator sdk.AccAddress, utoken, token sdk.Coin) error {
	err := k.setCollateralAmount(ctx, borrower, k.GetCollateralAmount(ctx, borrower, utoken.Denom).Sub(utoken))
	if err != nil {
		return err
	}
	if err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(utoken)); err != nil {
		return err
	}
	if err = k.setUTokenSupply(ctx, k.GetUTokenSupply(ctx, utoken.Denom).Sub(utoken)); err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, liquidator, sdk.NewCoins(token),
	)
}

// decollateralize removes fromAddr's uTokens from the module and sends them to toAddr.
// It occurs when decollateralizing uTokens (in which case fromAddr and toAddr are the
// same) as well as during liquidations, where toAddr is the liquidator.
func (k Keeper) decollateralize(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, coin sdk.Coin) error {
	err := k.setCollateralAmount(ctx, fromAddr, k.GetCollateralAmount(ctx, fromAddr, coin.Denom).Sub(coin))
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, sdk.NewCoins(coin))
}

// GetCollateralAmount returns an sdk.Coin representing how much of a given denom the
// x/leverage module account currently holds as collateral for a given borrower.
func (k Keeper) GetCollateralAmount(ctx sdk.Context, borrowerAddr sdk.AccAddress, denom string) sdk.Coin {
	store := ctx.KVStore(k.storeKey)
	collateral := sdk.NewCoin(denom, sdk.ZeroInt())
	key := types.CreateCollateralAmountKey(borrowerAddr, denom)

	if bz := store.Get(key); bz != nil {
		err := collateral.Amount.Unmarshal(bz)
		if err != nil {
			panic(err)
		}
	}

	return collateral
}

// setCollateralAmount sets the amount of a given denom the x/leverage module account
// currently holds as collateral for a given borrower. If the amount is zero, any
// stored value is cleared. A negative amount or invalid coin causes an error.
// This function does not move coins to or from the module account.
func (k Keeper) setCollateralAmount(ctx sdk.Context, borrowerAddr sdk.AccAddress, collateral sdk.Coin) error {
	if !collateral.IsValid() {
		return sdkerrors.Wrap(types.ErrInvalidAsset, collateral.String())
	}

	if borrowerAddr.Empty() {
		return types.ErrEmptyAddress
	}

	bz, err := collateral.Amount.Marshal()
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := types.CreateCollateralAmountKey(borrowerAddr, collateral.Denom)

	if collateral.Amount.IsZero() {
		store.Delete(key)
	} else {
		store.Set(key, bz)
	}
	return nil
}

// GetTotalCollateral returns an sdk.Coin representing how much of a given uToken
// the x/leverage module account currently holds as collateral. Non-uTokens return
// zero.
func (k Keeper) GetTotalCollateral(ctx sdk.Context, denom string) sdk.Int {
	if !types.HasUTokenPrefix(denom) {
		// non-uTokens cannot be collateral
		return sdk.ZeroInt()
	}

	// uTokens in the module account are always from collateral
	return k.ModuleBalance(ctx, denom)
}

// CalculateCollateralValue uses the price oracle to determine the value (in USD) provided by
// collateral sdk.Coins, using each token's uToken exchange rate.
// An error is returned if any input coins are not uTokens or if value calculation fails.
func (k Keeper) CalculateCollateralValue(ctx sdk.Context, collateral sdk.Coins) (sdk.Dec, error) {
	limit := sdk.ZeroDec()

	for _, coin := range collateral {
		// convert uToken collateral to base assets
		baseAsset, err := k.ExchangeUToken(ctx, coin)
		if err != nil {
			return sdk.ZeroDec(), err
		}

		// get USD value of base assets
		v, err := k.TokenValue(ctx, baseAsset)
		if err != nil {
			return sdk.ZeroDec(), err
		}

		// add each collateral coin's weighted value to borrow limit
		limit = limit.Add(v)
	}

	return limit, nil
}
