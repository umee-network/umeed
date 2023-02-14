//go:build experimental
// +build experimental

package uibc

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"gotest.tools/v3/assert"
)

func TestGenesisValidation(t *testing.T) {
	gs := DefaultGenesisState()
	err := gs.Validate()
	assert.NilError(t, err)

	gs.TotalOutflowSum = sdk.NewDec(-123123)
	err = gs.Validate()
	assert.ErrorContains(t, err, "total outflow sum cannot be negative")

	gs.Quotas = []sdk.DecCoin{{Denom: "umee", Amount: sdk.NewDec(-11123123)}}
	err = gs.Validate()
	assert.ErrorContains(t, err, "amount cannot be negative")
}
