package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// UTokenPrefix defines the uToken denomination prefix for all uToken types.
	UTokenPrefix = "u/"
)

// HasUTokenPrefix detects the uToken prefix on a denom.
func HasUTokenPrefix(denom string) bool {
	return strings.HasPrefix(denom, UTokenPrefix)
}

// ToUTokenDenom adds the uToken prefix to a denom if not already present.
func ToUTokenDenom(denom string) string {
	if HasUTokenPrefix(denom) {
		return denom
	}
	return UTokenPrefix + denom
}

// ToTokenDenom strips the uToken prefix from a denom if present.
func ToTokenDenom(denom string) string {
	return strings.TrimPrefix(denom, UTokenPrefix)
}

// Validate performs validation on an Token type returning an error if the Token
// is invalid.
func (t Token) Validate() error {
	if err := sdk.ValidateDenom(t.BaseDenom); err != nil {
		return err
	}
	if HasUTokenPrefix(t.BaseDenom) {
		// prevent base asset denoms that start with "u/"
		return sdkerrors.Wrap(ErrInvalidAsset, t.BaseDenom)
	}

	if err := sdk.ValidateDenom(t.SymbolDenom); err != nil {
		return err
	}
	if HasUTokenPrefix(t.SymbolDenom) {
		// prevent symbol denoms that start with "u/"
		return sdkerrors.Wrap(ErrInvalidAsset, t.SymbolDenom)
	}

	// Reserve factor and collateral weight range between 0 and 1, inclusive.
	if t.ReserveFactor.IsNegative() || t.ReserveFactor.GT(sdk.OneDec()) {
		return fmt.Errorf("invalid reserve factor: %s", t.ReserveFactor)
	}

	if t.CollateralWeight.IsNegative() || t.CollateralWeight.GT(sdk.OneDec()) {
		return fmt.Errorf("invalid collateral rate: %s", t.CollateralWeight)
	}

	// Liquidation threshold ranges between collateral weight and 1, inclusive.
	if t.LiquidationThreshold.LT(t.CollateralWeight) || t.LiquidationThreshold.GT(sdk.OneDec()) {
		return fmt.Errorf("invalid liquidation threshold: %s", t.LiquidationThreshold)
	}

	// Kink utilization rate ranges between 0 and 1, exclusive. This prevents
	// multiple interest rates being defined at exactly 0% or 100% utilization
	// e.g. kink at 0%, 2% base borrow rate, 4% borrow rate at kink.
	if !t.KinkUtilization.IsPositive() || t.KinkUtilization.GTE(sdk.OneDec()) {
		return fmt.Errorf("invalid kink utilization rate: %s", t.KinkUtilization)
	}

	// interest rates are non-negative; they do not need to have a maximum value
	if t.BaseBorrowRate.IsNegative() {
		return fmt.Errorf("invalid base borrow rate: %s", t.BaseBorrowRate)
	}
	if t.KinkBorrowRate.IsNegative() {
		return fmt.Errorf("invalid kink borrow rate: %s", t.KinkBorrowRate)
	}
	if t.MaxBorrowRate.IsNegative() {
		return fmt.Errorf("invalid max borrow rate: %s", t.MaxBorrowRate)
	}

	// Liquidation incentive is non-negative
	if t.LiquidationIncentive.IsNegative() {
		return fmt.Errorf("invalid liquidation incentive: %s", t.LiquidationIncentive)
	}

	// Blacklisted assets cannot have borrow or supply enabled
	if t.Blacklist {
		if t.EnableMsgBorrow {
			return fmt.Errorf("blacklisted assets cannot have borrowing enabled")
		}
		if t.EnableMsgSupply {
			return fmt.Errorf("blacklisted assets cannot have supplying enabled")
		}
	}

	if t.MaxCollateralShare.IsNegative() || t.MaxCollateralShare.GT(sdk.OneDec()) {
		return sdkerrors.ErrInvalidRequest.Wrap("Token.MaxCollateralShare must be between 0 and 1")
	}

	if t.MaxSupplyUtilization.IsNegative() || t.MaxSupplyUtilization.GT(sdk.OneDec()) {
		return sdkerrors.ErrInvalidRequest.Wrap("Token.MaxSupplyUtilization must be between 0 and 1")
	}

	if t.MinCollateralLiquidity.IsNegative() || t.MaxSupplyUtilization.GT(sdk.OneDec()) {
		return sdkerrors.ErrInvalidRequest.Wrap("Token.MinCollateralLiquidity be between 0 and 1")
	}

	return nil
}

// AssertSupplyEnabled returns an error if a Token cannot be supplied.
func (t Token) AssertSupplyEnabled() error {
	if !t.EnableMsgSupply {
		return sdkerrors.Wrap(ErrSupplyNotAllowed, t.BaseDenom)
	}
	return nil
}

// AssertBorrowEnabled returns an error if a Token cannot be borrowed.
func (t Token) AssertBorrowEnabled() error {
	if !t.EnableMsgBorrow {
		return sdkerrors.Wrap(ErrBorrowNotAllowed, t.BaseDenom)
	}
	return nil
}

// AssertNotBlacklisted returns an error if a Token is blacklisted.
func (t Token) AssertNotBlacklisted() error {
	if t.Blacklist {
		return sdkerrors.Wrap(ErrBlacklisted, t.BaseDenom)
	}
	return nil
}
