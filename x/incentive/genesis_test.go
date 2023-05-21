package incentive

import (
	"testing"

	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v4/util/coin"
	leveragetypes "github.com/umee-network/umee/v4/x/leverage/types"
)

func TestValidateGenesis(t *testing.T) {
	validAddr := "umee1s84d29zk3k20xk9f0hvczkax90l9t94g72n6wm"

	genesis := DefaultGenesis()
	assert.NilError(t, genesis.Validate())

	invalidParams := DefaultGenesis()
	invalidParams.Params.EmergencyUnbondFee = sdk.MustNewDecFromStr("-0.01")
	assert.ErrorContains(t, invalidParams.Validate(), "invalid emergency unbonding fee")

	zeroID := DefaultGenesis()
	zeroID.NextProgramId = 0
	assert.ErrorIs(t, zeroID.Validate(), ErrInvalidProgramID)

	negativeTime := DefaultGenesis()
	negativeTime.LastRewardsTime = -1
	assert.ErrorIs(t, negativeTime.Validate(), ErrDecreaseLastRewardTime)

	invalidRewardTracker := DefaultGenesis()
	invalidRewardTracker.RewardTrackers = []RewardTracker{{}}
	assert.ErrorContains(t, invalidRewardTracker.Validate(), "empty address string is not allowed")

	rt := RewardTracker{
		Account: validAddr,
		UToken:  "u/uumee",
		Rewards: sdk.NewDecCoins(),
	}
	duplicateRewardTracker := DefaultGenesis()
	duplicateRewardTracker.RewardTrackers = []RewardTracker{rt, rt}
	assert.ErrorContains(t, duplicateRewardTracker.Validate(), "duplicate reward trackers")

	invalidRewardAccumulator := DefaultGenesis()
	invalidRewardAccumulator.RewardAccumulators = []RewardAccumulator{{}}
	assert.ErrorIs(t, invalidRewardAccumulator.Validate(), leveragetypes.ErrNotUToken)

	ra := RewardAccumulator{
		UToken:  "u/uumee",
		Rewards: sdk.NewDecCoins(),
	}
	duplicateRewardAccumulator := DefaultGenesis()
	duplicateRewardAccumulator.RewardAccumulators = []RewardAccumulator{ra, ra}
	assert.ErrorContains(t, duplicateRewardAccumulator.Validate(), "duplicate reward accumulators")

	invalidProgram := IncentiveProgram{}
	validProgram := NewIncentiveProgram(1, 1, 1, "u/uumee", sdk.NewInt64Coin("uumee", 1), coin.Zero("uumee"), false)

	invalidUpcomingProgram := DefaultGenesis()
	invalidUpcomingProgram.UpcomingPrograms = []IncentiveProgram{invalidProgram}
	assert.ErrorIs(t, invalidUpcomingProgram.Validate(), ErrInvalidProgramID)

	duplicateUpcomingProgram := DefaultGenesis()
	duplicateUpcomingProgram.UpcomingPrograms = []IncentiveProgram{validProgram, validProgram}
	assert.ErrorContains(t, duplicateUpcomingProgram.Validate(), "duplicate upcoming incentive programs")

	invalidOngoingProgram := DefaultGenesis()
	invalidOngoingProgram.OngoingPrograms = []IncentiveProgram{invalidProgram}
	assert.ErrorIs(t, invalidOngoingProgram.Validate(), ErrInvalidProgramID)

	duplicateOngoingProgram := DefaultGenesis()
	duplicateOngoingProgram.UpcomingPrograms = []IncentiveProgram{validProgram}
	duplicateOngoingProgram.OngoingPrograms = []IncentiveProgram{validProgram}
	assert.ErrorContains(t, duplicateOngoingProgram.Validate(), "duplicate ongoing incentive programs")

	invalidCompletedProgram := DefaultGenesis()
	invalidCompletedProgram.CompletedPrograms = []IncentiveProgram{invalidProgram}
	assert.ErrorIs(t, invalidCompletedProgram.Validate(), ErrInvalidProgramID)

	duplicateCompletedProgram := DefaultGenesis()
	duplicateCompletedProgram.UpcomingPrograms = []IncentiveProgram{validProgram}
	duplicateCompletedProgram.CompletedPrograms = []IncentiveProgram{validProgram}
	assert.ErrorContains(t, duplicateCompletedProgram.Validate(), "duplicate completed incentive programs")

	invalidBond := DefaultGenesis()
	invalidBond.Bonds = []Bond{{}}
	assert.ErrorContains(t, invalidBond.Validate(), "empty address string is not allowed")

	b := Bond{
		Account: validAddr,
		UToken:  sdk.NewInt64Coin("uumee", 1),
	}

	duplicateBond := DefaultGenesis()
	duplicateBond.Bonds = []Bond{b, b}
	assert.ErrorContains(t, duplicateBond.Validate(), "duplicate bonds")

	invalidAccountUnbondings := DefaultGenesis()
	invalidAccountUnbondings.AccountUnbondings = []AccountUnbondings{{}}
	assert.ErrorContains(t, invalidAccountUnbondings.Validate(), "empty address string is not allowed")

	au := AccountUnbondings{
		Account:    validAddr,
		UToken:     "u/uumee",
		Unbondings: []Unbonding{},
	}
	duplicateAccountUnbonding := DefaultGenesis()
	duplicateAccountUnbonding.AccountUnbondings = []AccountUnbondings{au, au}
	assert.ErrorContains(t, duplicateAccountUnbonding.Validate(), "duplicate account unbondings")
}
