package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/x/leverage/types"
)

const (
	routeExchangeRates = "exchange-rates"
	routeReserveAmount = "reserve-amount"
	routeBorrowAmount  = "borrow-amount"
)

// RegisterInvariants registers the leverage module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, routeExchangeRates, ExchangeRatesInvariant(k))
	ir.RegisterRoute(types.ModuleName, routeReserveAmount, ReserveAmountInvariant(k))
	ir.RegisterRoute(types.ModuleName, routeBorrowAmount, BorrowAmountInvariant(k))
}

// AllInvariants runs all invariants of the x/leverage module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := ExchangeRatesInvariant(k)(ctx)
		if stop {
			return res, stop
		}

		res, stop = ReserveAmountInvariant(k)(ctx)
		if stop {
			return res, stop
		}

		return BorrowAmountInvariant(k)(ctx)
	}
}

// ExchangeRatesInvariant checks that all exchante rate denom are bigger or equal to 1
func ExchangeRatesInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var (
			msg   string
			count int
		)

		store := ctx.KVStore(k.storeKey)
		exchangeRatePrefix := types.CreateExchangeRateKeyNoDenom()

		iter := sdk.KVStorePrefixIterator(store, exchangeRatePrefix)
		defer iter.Close()

		// Iterate throught all denoms which have an exchange rate stored
		// in the keeper. If a token is registered but its exchange rate is
		// lower than 1.0 or it has some error doing the unmarshal it
		// adds the denom invariant count and message description
		for ; iter.Valid(); iter.Next() {
			// key is exchangeRatePrefix | denom | 0x00
			key, val := iter.Key(), iter.Value()

			// remove exchangeRatePrefix and null-terminator
			denom := types.DenomFromKey(key, exchangeRatePrefix)

			amount := sdk.ZeroDec()
			if err := amount.Unmarshal(val); err != nil {
				count++
				msg += fmt.Sprintf("\t%s received an error while Unmarshal the byte %+v\n", denom, val)
				continue
			}

			if amount.LT(sdk.OneDec()) {
				count++
				msg += fmt.Sprintf("\t%s exchange rate %s is lower than one\n", denom, amount.String())
			}
		}

		broken := count != 0

		return sdk.FormatInvariant(
			types.ModuleName, routeExchangeRates,
			fmt.Sprintf("amount of exchange rate lower than one %d\n%s", count, msg),
		), broken
	}
}

// ReserveAmountInvariant checks that reserve amounts have non-negative balances
func ReserveAmountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var (
			msg   string
			count int
		)

		store := ctx.KVStore(k.storeKey)
		reserveAmountPrefix := types.CreateReserveAmountKeyNoDenom()

		iter := sdk.KVStorePrefixIterator(store, reserveAmountPrefix)
		defer iter.Close()

		// Iterate throught all denoms which have an reserve amount stored
		// in the keeper. If a token is registered but its reserve amount is
		// negative or it has some error doing the unmarshal it
		// adds the denom invariant count and message description
		for ; iter.Valid(); iter.Next() {
			// key is reserveAmountPrefix | denom | 0x00
			key, val := iter.Key(), iter.Value()

			// remove reserveAmountPrefix and null-terminator
			denom := types.DenomFromKey(key, reserveAmountPrefix)

			amount := sdk.ZeroInt()
			if err := amount.Unmarshal(val); err != nil {
				count++
				msg += fmt.Sprintf("\t%s received an error while Unmarshal the byte %+v\n", denom, val)
				continue
			}

			if amount.IsNegative() {
				count++
				msg += fmt.Sprintf("\t%s reserve amount %s is negative\n", denom, amount.String())
			}
		}

		broken := count != 0

		return sdk.FormatInvariant(
			types.ModuleName, routeReserveAmount,
			fmt.Sprintf("number of negative reserve amount found %d\n%s", count, msg),
		), broken
	}
}

// BorrowAmountInvariant checks that borrow amounts have all positive values
func BorrowAmountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var (
			msg   string
			count int
		)

		store := ctx.KVStore(k.storeKey)
		loanKeyPrefix := types.CreateLoanKeyNoAddress()

		iter := sdk.KVStorePrefixIterator(store, loanKeyPrefix)
		defer iter.Close()

		// Iterate throught all denoms which have an borrow amount stored
		// in the keeper. If a token is registered but its borrow amount is
		// not positive or it has some error doing the unmarshal it
		// adds the denom invariant count and message description
		for ; iter.Valid(); iter.Next() {
			// key is prefix | lengthPrefixed(borrowerAddr) | denom | 0x00
			key, val := iter.Key(), iter.Value()

			// remove prefix | lengthPrefixed(addr) and null-terminator
			denom := types.DenomFromKeyWithAddress(key, loanKeyPrefix)

			amount := sdk.ZeroInt()
			if err := amount.Unmarshal(val); err != nil {
				count++
				msg += fmt.Sprintf("\t%s received an error while Unmarshal the byte %+v\n", denom, val)
				continue
			}

			if !amount.IsPositive() {
				count++
				msg += fmt.Sprintf("\t%s borrow amount %s is not positive\n", denom, amount.String())
			}
		}

		broken := count != 0

		return sdk.FormatInvariant(
			types.ModuleName, routeBorrowAmount,
			fmt.Sprintf("number of not positive borrow amount found %d\n%s", count, msg),
		), broken
	}
}
