package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v3/x/leverage/types"
)

// liquidateCollateral burns uToken collateral and sends the base token reward to the liquidator.
// This occurs during direct liquidation.
func (k Keeper) liquidateCollateral(ctx sdk.Context, borrower, liquidator sdk.AccAddress, uToken, token sdk.Coin,
) error {
	if err := k.burnCollateral(ctx, borrower, uToken); err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, liquidator, sdk.NewCoins(token))
}

// burnCollateral removes some uTokens from an account's collateral and burns them. This occurs
// during liquidations.
func (k Keeper) burnCollateral(ctx sdk.Context, addr sdk.AccAddress, uToken sdk.Coin) error {
	err := k.setCollateralAmount(ctx, addr, k.GetCollateralAmount(ctx, addr, uToken.Denom).Sub(uToken))
	if err != nil {
		return err
	}
	if err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(uToken)); err != nil {
		return err
	}
	return k.setUTokenSupply(ctx, k.GetUTokenSupply(ctx, uToken.Denom).Sub(uToken))
}

// decollateralize removes fromAddr's uTokens from the module and sends them to toAddr.
// It occurs when decollateralizing uTokens (in which case fromAddr and toAddr are the
// same) as well as during non-direct liquidations, where toAddr is the liquidator.
func (k Keeper) decollateralize(ctx sdk.Context, fromAddr, toAddr sdk.AccAddress, uToken sdk.Coin) error {
	err := k.setCollateralAmount(ctx, fromAddr, k.GetCollateralAmount(ctx, fromAddr, uToken.Denom).Sub(uToken))
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, toAddr, sdk.NewCoins(uToken))
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
	if err := collateral.Validate(); err != nil {
		return err
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
// the x/leverage module account currently holds as collateral. Non-uTokens return zero.
func (k Keeper) GetTotalCollateral(ctx sdk.Context, denom string) sdk.Coin {
	if !types.HasUTokenPrefix(denom) {
		// non-uTokens cannot be collateral
		return sdk.Coin{}
	}

	// uTokens in the module account are always from collateral
	return sdk.NewCoin(denom, k.ModuleBalance(ctx, denom))
}

// CalculateCollateralValue uses the price oracle to determine the value (in USD) provided by
// collateral sdk.Coins, using each token's uToken exchange rate.
// An error is returned if any input coins are not uTokens or if value calculation fails.
func (k Keeper) CalculateCollateralValue(ctx sdk.Context, collateral sdk.Coins) (sdk.Dec, error) {
	total := sdk.ZeroDec()

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
		total = total.Add(v)
	}

	return total, nil
}

// GetAllTotalCollateral returns total collateral across all uTokens.
func (k Keeper) GetAllTotalCollateral(ctx sdk.Context) sdk.Coins {
	total := sdk.NewCoins()

	tokens := k.GetAllRegisteredTokens(ctx)
	for _, t := range tokens {
		uDenom := types.ToUTokenDenom(t.BaseDenom)
		total = total.Add(k.GetTotalCollateral(ctx, uDenom))
	}
	return total
}

// CollateralLiquidity calculates the current collateral liquidity of a token denom,
// which is defined as the token's liquidity, divided by the base token equivalent
// of associated uToken's total collateral. Ranges from 0 to 1.0
func (k Keeper) CollateralLiquidity(ctx sdk.Context, denom string) sdk.Dec {
	totalCollateral := k.GetTotalCollateral(ctx, types.ToUTokenDenom(denom))
	exchangeRate := k.DeriveExchangeRate(ctx, denom)
	liquidity := k.AvailableLiquidity(ctx, denom)

	// Zero collateral will be interpreted as full liquidity. This encompasses two cases:
	// - 0/0 liquidity / collateral: Empty market, system is considered healthy by default
	// - x/0 liquidity / collateral: No collateral but plenty of liquidity, also considered healthy
	// In both cases, "all collateral is liquid" is technically true, given that there is no collateral.
	if totalCollateral.IsZero() {
		return sdk.OneDec()
	}

	collateralLiquidity := toDec(liquidity).Quo(exchangeRate.MulInt(totalCollateral.Amount))

	// Liquidity above 100% is ignored
	return sdk.MinDec(collateralLiquidity, sdk.OneDec())
}

// CollateralShare calculates the portion of overall collateral
// (measured in USD value) that a given uToken denom represents.
func (k *Keeper) CollateralShare(ctx sdk.Context, denom string) (sdk.Dec, error) {
	systemCollateral := k.GetAllTotalCollateral(ctx)
	thisCollateral := sdk.NewCoins(sdk.NewCoin(denom, systemCollateral.AmountOf(denom)))

	// get USD collateral value for all uTokens combined
	totalValue, err := k.CalculateCollateralValue(ctx, systemCollateral)
	if err != nil {
		return sdk.ZeroDec(), err
	}

	// get USD collateral value for this uToken only
	thisValue, err := k.CalculateCollateralValue(ctx, thisCollateral)
	if err != nil {
		return sdk.ZeroDec(), err
	}

	if !totalValue.IsPositive() {
		return sdk.ZeroDec(), nil
	}
	return thisValue.Quo(totalValue), nil
}

// checkCollateralLiquidity returns the appropriate error if a token denom's
// collateral liquidity is below its MinCollateralLiquidity
func (k Keeper) checkCollateralLiquidity(ctx sdk.Context, denom string) error {
	token, err := k.GetTokenSettings(ctx, denom)
	if err != nil {
		return err
	}

	collateralLiquidity := k.CollateralLiquidity(ctx, denom)
	if collateralLiquidity.LT(token.MinCollateralLiquidity) {
		return types.ErrMinCollateralLiquidity.Wrap(collateralLiquidity.String())
	}
	return nil
}

// checkCollateralShare returns an error if a given uToken is above its collateral share
func (k *Keeper) checkCollateralShare(ctx sdk.Context, denom string) error {
	token, err := k.GetTokenSettings(ctx, types.ToTokenDenom(denom))
	if err != nil {
		return err
	}

	share, err := k.CollateralShare(ctx, denom)
	if err != nil {
		return err
	}

	if share.GT(token.MaxCollateralShare) {
		return types.ErrMaxCollateralShare.Wrapf("%s share is %s", denom, share)
	}
	return nil
}
