package incentive

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gotest.tools/v3/assert"
)

func TestDefaultParams(t *testing.T) {
	params := DefaultParams()
	assert.NilError(t, params.Validate())

	invalidMaxUnbondings := DefaultParams()
	invalidMaxUnbondings.MaxUnbondings = 0
	assert.ErrorContains(t, invalidMaxUnbondings.Validate(), "max unbondings cannot be zero")

	invalidUnbondingDuration := DefaultParams()
	invalidUnbondingDuration.UnbondingDuration = -1
	assert.ErrorContains(t, invalidUnbondingDuration.Validate(), "invalid unbonding duration")

	invalidEmergencyUnbondFee := DefaultParams()
	invalidEmergencyUnbondFee.EmergencyUnbondFee = sdk.OneDec()
	assert.ErrorContains(t, invalidEmergencyUnbondFee.Validate(), "invalid emergency unbonding fee")

	invalidCommunityFund := DefaultParams()
	invalidCommunityFund.CommunityFundAddress = "abcdefgh"
	assert.ErrorContains(t, invalidCommunityFund.Validate(), "decoding bech32 failed")
}
