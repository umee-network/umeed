package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v2/x/leverage/types"
)

// validateLendAsset validates an sdk.Coin and ensures its Denom is a Token with EnableMsgLend
func (k Keeper) validateLendAsset(ctx sdk.Context, loan sdk.Coin) error {
	if !loan.IsValid() {
		return types.ErrInvalidAsset.Wrap(loan.String())
	}
	token, err := k.GetTokenSettings(ctx, loan.Denom)
	if err != nil {
		return err
	}
	return token.AssertLendEnabled()
}

// validateBorrowAsset validates an sdk.Coin and ensures its Denom is a Token with EnableMsgBorrow
func (k Keeper) validateBorrowAsset(ctx sdk.Context, borrow sdk.Coin) error {
	if !borrow.IsValid() {
		return types.ErrInvalidAsset.Wrap(borrow.String())
	}

	token, err := k.GetTokenSettings(ctx, borrow.Denom)
	if err != nil {
		return err
	}
	return token.AssertBorrowEnabled()
}

// validateCollateralAsset validates an sdk.Coin and ensures its Denom is a Token with EnableMsgLend
// and CollateralWeight > 0
func (k Keeper) validateCollateralAsset(ctx sdk.Context, collateral sdk.Coin) error {
	if !collateral.IsValid() {
		return types.ErrInvalidAsset.Wrap(collateral.String())
	}

	tokenDenom := k.FromUTokenToTokenDenom(ctx, collateral.Denom)
	token, err := k.GetTokenSettings(ctx, tokenDenom)
	if err != nil {
		return err
	}
	if token.CollateralWeight.IsZero() {
		return types.ErrCollateralWeightZero
	}
	return token.AssertLendEnabled()
}
