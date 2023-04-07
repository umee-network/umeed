package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v4/util"
	"github.com/umee-network/umee/v4/util/keys"
	"github.com/umee-network/umee/v4/util/store"
	"github.com/umee-network/umee/v4/x/incentive"
)

// getAllIncentivePrograms returns all incentive programs
// that have been passed by governance and have a particular status.
// The status of an incentive program is either Upcoming, Ongoing, or Completed.
func (k Keeper) getAllIncentivePrograms(ctx sdk.Context, status incentive.ProgramStatus,
) ([]incentive.IncentiveProgram, error) {
	programs := []incentive.IncentiveProgram{}

	var prefix []byte
	switch status {
	case incentive.ProgramStatusUpcoming:
		prefix = keyPrefixUpcomingIncentiveProgram
	case incentive.ProgramStatusOngoing:
		prefix = keyPrefixOngoingIncentiveProgram
	case incentive.ProgramStatusCompleted:
		prefix = keyPrefixCompletedIncentiveProgram
	default:
		return []incentive.IncentiveProgram{}, incentive.ErrInvalidProgramStatus
	}

	iterator := func(_, val []byte) error {
		var p incentive.IncentiveProgram
		k.cdc.MustUnmarshal(val, &p)

		programs = append(programs, p)
		return nil
	}

	err := store.Iterate(k.KVStore(ctx), prefix, iterator)
	return programs, err
}

// getPaginatedIncentivePrograms returns all incentive programs
// that have been passed by governance and have a particular status.
// The status of an incentive program is either Upcoming, Ongoing, or Completed.
// Accepts pagination parameters which specify the length of a page and which page to fetch.
func (k Keeper) getPaginatedIncentivePrograms(
	ctx sdk.Context, status incentive.ProgramStatus, page, limit uint64,
) ([]incentive.IncentiveProgram, error) {
	programs := []incentive.IncentiveProgram{}

	var prefix []byte
	switch status {
	case incentive.ProgramStatusUpcoming:
		prefix = keyPrefixUpcomingIncentiveProgram
	case incentive.ProgramStatusOngoing:
		prefix = keyPrefixOngoingIncentiveProgram
	case incentive.ProgramStatusCompleted:
		prefix = keyPrefixCompletedIncentiveProgram
	default:
		return []incentive.IncentiveProgram{}, incentive.ErrInvalidProgramStatus
	}

	iterator := func(_, val []byte) error {
		var p incentive.IncentiveProgram
		k.cdc.MustUnmarshal(val, &p)

		programs = append(programs, p)
		return nil
	}

	err := store.IteratePaginated(k.KVStore(ctx), prefix, uint(page), uint(limit), iterator)
	return programs, err
}

// getAllBondDenoms gets all uToken denoms for which an account has nonzero bonded amounts.
// useful for setting up queries which look at all of an account's bonds or unbondings.
func (k Keeper) getAllBondDenoms(ctx sdk.Context, addr sdk.AccAddress) []string {
	prefix := keyBondAmountNoDenom(addr)
	bonds := []string{}

	iterator := func(key, val []byte) error {
		_, denom, _, err := keys.LeadingAddressAndDenom(keyPrefixBondAmount, key)
		if err != nil {
			return err
		}
		bonds = append(bonds, denom)
		return nil
	}

	util.Panic(store.Iterate(k.KVStore(ctx), prefix, iterator))
	return bonds
}

// getAllBonds gets all bonds for all accounts (used during export genesis)
func (k Keeper) getAllBonds(ctx sdk.Context) []incentive.Bond {
	prefix := keyPrefixBondAmount
	bonds := []incentive.Bond{}

	iterator := func(key, val []byte) error {
		addr, denom, _, err := keys.LeadingAddressAndDenom(keyPrefixBondAmount, key)
		if err != nil {
			return err
		}
		amount := store.Int(val, "bond amount")
		bonds = append(bonds, incentive.NewBond(
			addr.String(),
			sdk.NewCoin(denom, amount),
		))

		return nil
	}

	util.Panic(store.Iterate(k.KVStore(ctx), prefix, iterator))
	return bonds
}

// getAllTotalBonded gets total bonded for all uTokens (used for a query)
func (k Keeper) getAllTotalBonded(ctx sdk.Context) sdk.Coins {
	prefix := keyPrefixTotalBonded
	total := sdk.NewCoins()

	iterator := func(key, val []byte) error {
		denom, _, err := keys.LeadingDenom(keyPrefixTotalBonded, key)
		if err != nil {
			return err
		}
		amount := store.Int(val, "total bonded")
		total = total.Add(sdk.NewCoin(denom, amount))
		return nil
	}

	util.Panic(store.Iterate(k.KVStore(ctx), prefix, iterator))
	return total
}

// getAllRewardTrackers gets all reward trackers for all accounts (used during export genesis)
func (k Keeper) getAllRewardTrackers(ctx sdk.Context) []incentive.RewardTracker {
	prefix := keyPrefixRewardTracker
	rewardTrackers := []incentive.RewardTracker{}

	iterator := func(_, val []byte) error {
		tracker := incentive.RewardTracker{}
		k.cdc.MustUnmarshal(val, &tracker)
		rewardTrackers = append(rewardTrackers, tracker)
		return nil
	}

	util.Panic(store.Iterate(k.KVStore(ctx), prefix, iterator))
	return rewardTrackers
}

// getAllRewardAccumulators gets all reward accumulators for all uTokens (used during export genesis)
func (k Keeper) getAllRewardAccumulators(ctx sdk.Context) []incentive.RewardAccumulator {
	prefix := keyPrefixRewardAccumulator
	rewardAccumulators := []incentive.RewardAccumulator{}

	iterator := func(_, val []byte) error {
		accumulator := incentive.RewardAccumulator{}
		k.cdc.MustUnmarshal(val, &accumulator)
		rewardAccumulators = append(rewardAccumulators, accumulator)
		return nil
	}

	util.Panic(store.Iterate(k.KVStore(ctx), prefix, iterator))
	return rewardAccumulators
}

// getAllAccountUnbondings gets all account unbondings for all accounts (used during export genesis)
func (k Keeper) getAllAccountUnbondings(ctx sdk.Context) []incentive.AccountUnbondings {
	prefix := keyPrefixUnbondings
	unbondings := []incentive.AccountUnbondings{}

	iterator := func(key, val []byte) error {
		au := incentive.AccountUnbondings{}
		k.cdc.MustUnmarshal(val, &au)
		unbondings = append(unbondings, au)
		return nil
	}

	util.Panic(store.Iterate(k.KVStore(ctx), prefix, iterator))
	return unbondings
}

// getAllTotalUnbonding gets total unbonding for all uTokens (used for a query)
func (k Keeper) getAllTotalUnbonding(ctx sdk.Context) sdk.Coins {
	prefix := keyPrefixTotalUnbonding
	total := sdk.NewCoins()

	iterator := func(key, val []byte) error {
		denom, _, err := keys.LeadingDenom(keyPrefixTotalUnbonding, key)
		if err != nil {
			return err
		}
		amount := store.Int(val, "total unbonding")
		total = total.Add(sdk.NewCoin(denom, amount))
		return nil
	}

	util.Panic(store.Iterate(k.KVStore(ctx), prefix, iterator))
	return total
}
