package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	leveragetypes "github.com/umee-network/umee/v4/x/leverage/types"
	"github.com/umee-network/umee/v4/x/oracle/types"
)

// Hooks defines a structure around the x/oracle Keeper that implements various
// Hooks interface defined by other modules such as x/leverage.
type Hooks struct {
	k Keeper
}

var _ leveragetypes.Hooks = Hooks{}

// Hooks returns a new Hooks instance that wraps the x/oracle keeper.
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// AfterTokenRegistered implements the x/leverage Hooks interface. Specifically,
// it checks if the provided Token should be added to the existing accepted list
// of assets for the x/oracle module.
func (h Hooks) AfterTokenRegistered(ctx sdk.Context, token leveragetypes.Token) {
	if token.Blacklist {
		// Blacklisted tokens should not trigger updates to the oracle accept list
		return
	}

	acceptList := h.k.AcceptList(ctx)

	var tokenExists bool
	for _, t := range acceptList {
		// On case-insensitive match of symbol denom, update base denom and exponent
		// to match the leverage registry. Does not modify existing symbol denom.
		if strings.EqualFold(t.SymbolDenom, token.SymbolDenom) {
			tokenExists = true
			t.BaseDenom = token.BaseDenom
			t.Exponent = token.Exponent
			break
		}
	}

	if !tokenExists {
		acceptList = append(acceptList, types.Denom{
			BaseDenom:   token.BaseDenom,
			SymbolDenom: token.SymbolDenom,
			Exponent:    token.Exponent,
		})
	}

	h.k.SetAcceptList(ctx, acceptList)
}

// AfterRegisteredTokenRemoved implements the x/leverage Hooks interface.
// Currently, it performs a no-op, however, we may want to remove tokens from
// the accept list in the future when they're removed from the x/leverage Token
// registry.
//
// We don't remove the token from the accept list because it might prove to be
// useful to still have price data for assets outside of the scope of the
// x/leverage registry. If assets need to be removed, they can always be purged
// via param change proposals.
func (h Hooks) AfterRegisteredTokenRemoved(sdk.Context, leveragetypes.Token) {}
