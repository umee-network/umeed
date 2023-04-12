package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/umee-network/umee/v4/util/store"
	"github.com/umee-network/umee/v4/x/incentive"
)

func (k Keeper) GetParams(ctx sdk.Context) incentive.Params {
	params := incentive.Params{}
	ok := store.GetObject(k.KVStore(ctx), k.cdc, keyPrefixParams, &params, "params")
	if !ok {
		// on missing module parameters, return defaults rather than panicking or returning zero values
		return incentive.DefaultParams()
	}
	return params
}

// setParams validates and sets the incentive module parameters
func (k Keeper) setParams(ctx sdk.Context, params incentive.Params) error {
	kvs := k.KVStore(ctx)
	if err := params.Validate(); err != nil {
		return err
	}
	return store.SetObject(kvs, k.cdc, keyPrefixParams, &params, "params")
}

// getIncentiveProgram gets an incentive program by ID, regardless of the program's status.
// Returns the program's status if found, or an error if it does not exist.
func (k Keeper) getIncentiveProgram(ctx sdk.Context, id uint32) (
	incentive.IncentiveProgram, incentive.ProgramStatus, error,
) {
	statuses := []incentive.ProgramStatus{
		incentive.ProgramStatusUpcoming,
		incentive.ProgramStatusOngoing,
		incentive.ProgramStatusCompleted,
	}

	kvStore := k.KVStore(ctx)
	program := incentive.IncentiveProgram{}

	// Looks for an incentive program with the specified ID in upcoming, ongoing, then completed program lists.
	for _, status := range statuses {
		ok := store.GetObject(kvStore, k.cdc, keyIncentiveProgram(id, status), &program, "incentive program")
		if ok {
			if program.ID != id {
				return incentive.IncentiveProgram{}, 0, incentive.ErrInvalidProgramID
			}
			return program, status, nil
		}
	}

	// If the program was not found in any of the three lists
	return incentive.IncentiveProgram{}, 0, sdkerrors.ErrNotFound
}

// setIncentiveProgram stores an incentive program in either the upcoming, ongoing, or completed program lists.
// does not validate the incentive program struct or the validity of its status change (e.g. upcoming -> complete)
func (k Keeper) setIncentiveProgram(ctx sdk.Context,
	program incentive.IncentiveProgram, status incentive.ProgramStatus,
) error {
	keys := [][]byte{
		keyIncentiveProgram(program.ID, incentive.ProgramStatusUpcoming),
		keyIncentiveProgram(program.ID, incentive.ProgramStatusOngoing),
		keyIncentiveProgram(program.ID, incentive.ProgramStatusCompleted),
	}

	kvStore := k.KVStore(ctx)
	for _, key := range keys {
		// always clear the program from the status it was prevously stored under
		kvStore.Delete(key)
	}

	key := keyIncentiveProgram(program.ID, status)
	return store.SetObject(kvStore, k.cdc, key, &program, "incentive program")
}

// getNextProgramID gets the ID that will be assigned to the next incentive program passed by governance.
func (k Keeper) getNextProgramID(ctx sdk.Context) uint32 {
	return store.GetUint32(k.KVStore(ctx), keyPrefixNextProgramID, "next program ID")
}

// setNextProgramID sets the ID that will be assigned to the next incentive program passed by governance.
func (k Keeper) setNextProgramID(ctx sdk.Context, id uint32) error {
	prev := k.getNextProgramID(ctx)
	if id < prev {
		return incentive.ErrDecreaseNextProgramID.Wrapf("%d to %d", id, prev)
	}
	return store.SetUint32(k.KVStore(ctx), keyPrefixNextProgramID, id, "next program ID")
}

// getLastRewardsTime gets the last unix time incentive rewards were computed globally by EndBlocker.
// panics if it would return a negative value.
func (k Keeper) GetLastRewardsTime(ctx sdk.Context) int64 {
	t := store.GetInt64(k.KVStore(ctx), keyPrefixLastRewardsTime, "last reward time")
	if t < 0 {
		panic("negative last reward time")
	}
	return t
}

// setLastRewardsTime sets the last unix time incentive rewards were computed globally by EndBlocker.
// does not accept negative unix time.
func (k Keeper) setLastRewardsTime(ctx sdk.Context, time int64) error {
	prev := k.GetLastRewardsTime(ctx)
	if time < 0 || time < prev {
		return incentive.ErrDecreaseLastRewardTime.Wrapf("%d to %d", time, prev)
	}
	return store.SetInt64(k.KVStore(ctx), keyPrefixLastRewardsTime, time, "last reward time")
}

// getTotalBonded retrieves the total amount of uTokens of a given denom which are bonded to the incentive module
func (k Keeper) getTotalBonded(ctx sdk.Context, denom string) sdk.Coin {
	key := keyTotalBonded(denom)
	amount := store.GetInt(k.KVStore(ctx), key, "total bonded")
	return sdk.NewCoin(denom, amount)
}

// GetBonded retrieves the amount of uTokens of a given denom which are bonded by an account
func (k Keeper) GetBonded(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	key := keyBondAmount(addr, denom)
	amount := store.GetInt(k.KVStore(ctx), key, "bonded amount")
	return sdk.NewCoin(denom, amount)
}

// setBonded sets the amount of uTokens of a given denom which are bonded by an account.
// Automatically updates TotalBonded as well.
//
// REQUIREMENT: This is the only function which is allowed to set TotalBonded.
func (k Keeper) setBonded(ctx sdk.Context, addr sdk.AccAddress, uToken sdk.Coin) error {
	// compute the change in bonded amount (can be negative when bond decreases)
	delta := uToken.Amount.Sub(k.GetBonded(ctx, addr, uToken.Denom).Amount)

	// Set bond amount
	key := keyBondAmount(addr, uToken.Denom)
	if err := store.SetInt(k.KVStore(ctx), key, uToken.Amount, "bonded amount"); err != nil {
		return err
	}

	// Update total bonded for this utoken denom using the computed change
	total := k.getTotalBonded(ctx, uToken.Denom)
	totalkey := keyTotalBonded(uToken.Denom)
	return store.SetInt(k.KVStore(ctx), totalkey, total.Amount.Add(delta), "total bonded")
}

// getUnbonding retrieves the amount of uTokens of a given denom which are unbonding by an account
func (k Keeper) getUnbondingAmount(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin {
	key := keyUnbondAmount(addr, denom)
	amount := store.GetInt(k.KVStore(ctx), key, "unbonding amount")
	return sdk.NewCoin(denom, amount)
}

// getTotalUnbonding retrieves the total amount of uTokens of a given denom which are unbonding from
// the incentive module
func (k Keeper) getTotalUnbonding(ctx sdk.Context, denom string) sdk.Coin {
	key := keyTotalUnbonding(denom)
	amount := store.GetInt(k.KVStore(ctx), key, "total unbonding")
	return sdk.NewCoin(denom, amount)
}

// getUnbondings gets all unbondings currently associated with an account and bonded denom.
func (k Keeper) getUnbondings(ctx sdk.Context, addr sdk.AccAddress, denom string) []incentive.Unbonding {
	key := keyUnbondings(addr, denom)
	kvStore := k.KVStore(ctx)

	accUnbondings := incentive.AccountUnbondings{}
	ok := store.GetObject(kvStore, k.cdc, key, &accUnbondings, "account unbondings")
	if !ok {
		return []incentive.Unbonding{}
	}
	return accUnbondings.Unbondings
}

// setUnbondings stores the list of unbondings currently associated with an account and denom.
// It also updates the account's stored unbonding amounts, and thus the module's total unbonding as well.
// Any zero-amount unbondings are automatically filtered out before storage.
//
// REQUIREMENT: This is the only function which is allowed to set unbonding amounts and total unbonding.
func (k Keeper) setUnbondings(ctx sdk.Context, unbondings incentive.AccountUnbondings) error {
	kvStore := k.KVStore(ctx)
	addr, err := sdk.AccAddressFromBech32(unbondings.Account)
	if err != nil {
		// catches invalid and empty addresses
		return err
	}
	denom := unbondings.Denom

	// remove any zero-amount unbondings before setting
	nonzeroUnbondings := []incentive.Unbonding{}
	for _, u := range unbondings.Unbondings {
		if u.Amount.Amount.IsPositive() {
			nonzeroUnbondings = append(nonzeroUnbondings, u)
		}
	}
	unbondings.Unbondings = nonzeroUnbondings

	// compute the new total unbonding specific to this account and denom.
	newUnbonding := sdk.ZeroInt()
	for _, u := range unbondings.Unbondings {
		newUnbonding = newUnbonding.Add(u.Amount.Amount)
	}
	// compute the change in unbonding amount (can be negative when unbonding decreases)
	delta := newUnbonding.Sub(k.getUnbondingAmount(ctx, addr, denom).Amount)

	// Update unbonding amount
	amountKey := keyUnbondAmount(addr, denom)
	if err := store.SetInt(k.KVStore(ctx), amountKey, newUnbonding, "unbonding amount"); err != nil {
		return err
	}

	// Update total unbonding for this utoken denom using the computed change
	total := k.getTotalUnbonding(ctx, denom)
	totalkey := keyTotalUnbonding(denom)
	if err := store.SetInt(k.KVStore(ctx), totalkey, total.Amount.Add(delta), "total unbonding"); err != nil {
		return err
	}

	// set list of unbondings
	key := keyUnbondings(addr, unbondings.Denom)
	if len(unbondings.Unbondings) == 0 {
		// clear store on no unbondings remaining
		kvStore.Delete(key)
	}
	return store.SetObject(kvStore, k.cdc, key, &unbondings, "account unbondings")
}

// getRewardAccumulator retrieves the reward accumulator of all reward tokens for a single bonded uToken -
// for example, how much UMEE, ATOM, etc (reward) would have been earned by 1 ATOM (bonded) since genesis.
func (k Keeper) getRewardAccumulator(ctx sdk.Context, bondDenom string) incentive.RewardAccumulator {
	key := keyRewardAccumulator(bondDenom)
	accumulator := incentive.RewardAccumulator{}
	ok := store.GetObject(k.KVStore(ctx), k.cdc, key, &accumulator, "reward accumulator")
	if ok {
		return accumulator
	}
	return incentive.NewRewardAccumulator(bondDenom, 0, sdk.NewDecCoins())
}

// setRewardAccumulator sets the full reward accumulator for a single bonded uToken.
func (k Keeper) setRewardAccumulator(ctx sdk.Context, accumulator incentive.RewardAccumulator) error {
	key := keyRewardAccumulator(accumulator.Denom)
	return store.SetObject(k.KVStore(ctx), k.cdc, key, &accumulator, "reward accumulator")
}

// GetRewardTracker retrieves the reward tracker for a single bonded uToken on one account - this is the value of
// the reward accumulator for that specific denom the last time this account performed any action that required
// a reward tracker update (i.e. Bond, Claim, BeginUnbonding, or being Liquidated).
func (k Keeper) getRewardTracker(ctx sdk.Context, addr sdk.AccAddress, bondDenom string) incentive.RewardTracker {
	key := keyRewardTracker(addr, bondDenom)
	tracker := incentive.RewardTracker{}
	ok := store.GetObject(k.KVStore(ctx), k.cdc, key, &tracker, "reward tracker")
	if ok {
		return tracker
	}
	return incentive.NewRewardTracker(addr.String(), bondDenom, sdk.NewDecCoins())
}

// setRewardTracker sets the reward tracker for a single bonded uToken for an address.
// automatically clear the entry if rewards are all zero.
func (k Keeper) setRewardTracker(ctx sdk.Context, tracker incentive.RewardTracker) error {
	key := keyRewardTracker(sdk.MustAccAddressFromBech32(tracker.Account), tracker.Denom)
	if tracker.Rewards.IsZero() {
		k.clearRewardTracker(ctx, sdk.MustAccAddressFromBech32(tracker.Account), tracker.Denom)
		return nil
	}
	return store.SetObject(k.KVStore(ctx), k.cdc, key, &tracker, "reward tracker")
}

// clearRewardTracker clears a reward tracker matching a specific account + bonded uToken from the store
func (k Keeper) clearRewardTracker(ctx sdk.Context, addr sdk.AccAddress, bondDenom string) {
	key := keyRewardTracker(addr, bondDenom)
	k.KVStore(ctx).Delete(key)
}
