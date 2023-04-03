package incentive

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(
	params Params,
	completedPrograms []IncentiveProgram,
	ongoingPrograms []IncentiveProgram,
	upcomingPrograms []IncentiveProgram,
	nextProgramID uint32,
	lastRewardTime int64,
	bonds []Bond,
	rewardTrackers []RewardTracker,
	rewardAccumulators []RewardAccumulator,
	accountUnbondings []AccountUnbondings,
) *GenesisState {
	return &GenesisState{
		Params:             params,
		CompletedPrograms:  completedPrograms,
		OngoingPrograms:    ongoingPrograms,
		UpcomingPrograms:   upcomingPrograms,
		NextProgramId:      nextProgramID,
		LastRewardsTime:    lastRewardTime,
		Bonds:              bonds,
		RewardTrackers:     rewardTrackers,
		RewardAccumulators: rewardAccumulators,
		AccountUnbondings:  accountUnbondings,
	}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:          DefaultParams(),
		NextProgramId:   1,
		LastRewardsTime: 0,
	}
}

// Validate checks a genesis state for basic issues
func (gs GenesisState) Validate() error {
	// TODO #1749
	return nil
}

// GetGenesisStateFromAppState returns x/incentive GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}

// NewIncentiveProgram creates the IncentiveProgram struct used in GenesisState
func NewIncentiveProgram(
	id uint32,
	startTime int64,
	duration int64,
	uDenom string,
	totalRewards, remainingRewards sdk.Coin,
	funded bool,
) IncentiveProgram {
	return IncentiveProgram{
		ID:               id,
		StartTime:        startTime,
		Duration:         duration,
		UToken:           uDenom,
		TotalRewards:     totalRewards,
		RemainingRewards: remainingRewards,
		Funded:           funded,
	}
}

// NewBond creates the Bond struct used in GenesisState
func NewBond(addr string, coin sdk.Coin) Bond {
	return Bond{
		Account: addr,
		Amount:  coin,
	}
}

// NewRewardTracker creates the RewardTracker struct used in GenesisState
func NewRewardTracker(addr, denom string, coins sdk.DecCoins) RewardTracker {
	return RewardTracker{
		Account: addr,
		Denom:   denom,
		Rewards: coins,
	}
}

// NewRewardAccumulator creates the RewardAccumulator struct used in GenesisState
func NewRewardAccumulator(denom string, exponent uint32, coins sdk.DecCoins) RewardAccumulator {
	return RewardAccumulator{
		Denom:    denom,
		Exponent: exponent,
		Rewards:  coins,
	}
}

// NewUnbonding creates the Unbonding struct used in GenesisState
func NewUnbonding(startTime, endTime int64, coin sdk.Coin) Unbonding {
	return Unbonding{
		Start:  startTime,
		End:    endTime,
		Amount: coin,
	}
}

// NewAccountUnbondings creates the AccountUnbondings struct used in GenesisState
func NewAccountUnbondings(addr, denom string, unbondings []Unbonding) AccountUnbondings {
	return AccountUnbondings{
		Account:    addr,
		Denom:      denom,
		Unbondings: unbondings,
	}
}
