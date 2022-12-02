package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/umee-network/umee/v3/util"
	"github.com/umee-network/umee/v3/x/oracle/types"
)

// HistoricMedians returns a list of a given denom's last numStamps medians.
func (k Keeper) HistoricMedians(
	ctx sdk.Context,
	denom string,
	numStamps uint64,
) []sdk.Dec {
	// calculate start block to iterate from
	blockPeriod := (numStamps - 1) * k.MedianStampPeriod(ctx)
	lastStampBlock := uint64(ctx.BlockHeight()) - (uint64(ctx.BlockHeight())%k.MedianStampPeriod(ctx) + 1)
	startBlock := lastStampBlock - blockPeriod

	return k.getMedians(ctx, denom, startBlock)
}

// CalcAndSetMedian uses all the historic prices of a given denom to
// calculate its median price at the current block and set it to the store.
// It will also call setMedianDeviation with the calculated median.
func (k Keeper) CalcAndSetMedian(
	ctx sdk.Context,
	denom string,
) {
	historicPrices := k.getHistoricPrices(ctx, denom)
	median := util.Median(historicPrices)
	block := uint64(ctx.BlockHeight())
	k.SetMedian(ctx, denom, block, median)
	k.calcAndSetMedianDeviation(ctx, denom, median, historicPrices)
}

func (k Keeper) SetMedian(
	ctx sdk.Context,
	denom string,
	blockNum uint64,
	median sdk.Dec,
) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdk.DecProto{Dec: median})
	store.Set(types.KeyMedian(denom, blockNum), bz)
}

// HistoricMedianDeviation returns a given denom's most recently stamped
// standard deviation around its median price at a given block.
func (k Keeper) HistoricMedianDeviation(
	ctx sdk.Context,
	denom string,
) (sdk.Dec, error) {
	store := ctx.KVStore(k.storeKey)
	blockDiff := uint64(ctx.BlockHeight())%k.MedianStampPeriod(ctx) + 1
	bz := store.Get(types.KeyMedianDeviation(denom, uint64(ctx.BlockHeight())-blockDiff))
	if bz == nil {
		return sdk.ZeroDec(), sdkerrors.Wrap(types.ErrNoMedianDeviation, fmt.Sprintf("denom: %s", denom))
	}

	decProto := sdk.DecProto{}
	k.cdc.MustUnmarshal(bz, &decProto)

	return decProto.Dec, nil
}

// WithinMedianDeviation returns whether or not a given price of a given
// denom is within the latest stamped Standard Deviation around the Median.
func (k Keeper) WithinMedianDeviation(
	ctx sdk.Context,
	denom string,
	price sdk.Dec,
) (bool, error) {
	// get latest median
	medians := k.HistoricMedians(
		ctx,
		denom,
		1,
	)
	if len(medians) == 0 {
		return false, sdkerrors.Wrap(types.ErrNoMedian, fmt.Sprintf("denom: %s", denom))
	}

	medianDeviation, err := k.HistoricMedianDeviation(ctx, denom)
	if err != nil {
		return false, err
	}

	return price.Sub(medians[0]).Abs().LTE(medianDeviation), nil
}

// setMedianDeviation sets a given denom's standard deviation around
// its median price in the current block.
func (k Keeper) calcAndSetMedianDeviation(
	ctx sdk.Context,
	denom string,
	median sdk.Dec,
	prices []sdk.Dec,
) {
	medianDeviation := util.MedianDeviation(median, prices)
	block := uint64(ctx.BlockHeight())
	k.SetMedianDeviation(ctx, denom, block, medianDeviation)
}

func (k Keeper) SetMedianDeviation(
	ctx sdk.Context,
	denom string,
	blockNum uint64,
	medianDeviation sdk.Dec,
) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdk.DecProto{Dec: medianDeviation})
	store.Set(types.KeyMedianDeviation(denom, blockNum), bz)
}

// MedianOfMedians calculates and returns the median of the last stampNum
// medians.
func (k Keeper) MedianOfMedians(
	ctx sdk.Context,
	denom string,
	numStamps uint64,
) sdk.Dec {
	medians := k.HistoricMedians(ctx, denom, numStamps)
	return util.Median(medians)
}

// AverageOfMedians calculates and returns the average of the last stampNum
// medians.
func (k Keeper) AverageOfMedians(
	ctx sdk.Context,
	denom string,
	numStamps uint64,
) sdk.Dec {
	medians := k.HistoricMedians(ctx, denom, numStamps)
	return util.Average(medians)
}

// MaxMedian calculates and returns the maximum value of the last stampNum
// medians.
func (k Keeper) MaxMedian(
	ctx sdk.Context,
	denom string,
	numStamps uint64,
) sdk.Dec {
	medians := k.HistoricMedians(ctx, denom, numStamps)
	return util.Max(medians)
}

// MinMedian calculates and returns the minimum value of the last stampNum
// medians.
func (k Keeper) MinMedian(
	ctx sdk.Context,
	denom string,
	numStamps uint64,
) sdk.Dec {
	medians := k.HistoricMedians(ctx, denom, numStamps)
	return util.Min(medians)
}

// getHistoricPrices returns all the historic prices of a given denom.
func (k Keeper) getHistoricPrices(
	ctx sdk.Context,
	denom string,
) []sdk.Dec {
	historicPrices := []sdk.Dec{}

	k.IterateHistoricPrices(ctx, denom, func(exchangeRate sdk.Dec) bool {
		historicPrices = append(historicPrices, exchangeRate)
		return false
	})

	return historicPrices
}

// IterateHistoricPrices iterates over historic prices of a given
// denom in the store.
// Iterator stops when exhausting the source, or when the handler returns `true`.
func (k Keeper) IterateHistoricPrices(
	ctx sdk.Context,
	denom string,
	handler func(sdk.Dec) bool,
) {
	store := ctx.KVStore(k.storeKey)

	// make sure we have one zero byte to correctly separate denoms
	prefix := util.ConcatBytes(1, types.KeyPrefixHistoricPrice, []byte(denom))
	iter := sdk.KVStorePrefixIterator(store, prefix)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		decProto := sdk.DecProto{}
		k.cdc.MustUnmarshal(iter.Value(), &decProto)
		if handler(decProto.Dec) {
			break
		}
	}
}

// getMedians returns median prices of a given denom in the store
// starting at startblock until the latest block.
func (k Keeper) getMedians(
	ctx sdk.Context,
	denom string,
	startBlock uint64,
) []sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	medians := []sdk.Dec{}
	currBlock := startBlock

	for currBlock < uint64(ctx.BlockHeight()) {
		bz := store.Get(types.KeyMedian(denom, currBlock))
		currBlock += k.MedianStampPeriod(ctx)

		if bz == nil {
			continue
		}
		decProto := sdk.DecProto{}
		k.cdc.MustUnmarshal(bz, &decProto)
		medians = append(medians, decProto.Dec)
	}

	return medians
}

// AddHistoricPrice adds the historic price of a denom at the current
// block height.
func (k Keeper) AddHistoricPrice(
	ctx sdk.Context,
	denom string,
	exchangeRate sdk.Dec,
) {
	block := uint64(ctx.BlockHeight())
	k.SetHistoricPrice(ctx, denom, block, exchangeRate)
}

func (k Keeper) SetHistoricPrice(
	ctx sdk.Context,
	denom string,
	blockNum uint64,
	exchangeRate sdk.Dec,
) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&sdk.DecProto{Dec: exchangeRate})
	store.Set(types.KeyHistoricPrice(denom, blockNum), bz)
}

// DeleteHistoricPrice deletes the historic price of a denom at a
// given block.
func (k Keeper) DeleteHistoricPrice(
	ctx sdk.Context,
	denom string,
	blockNum uint64,
) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyHistoricPrice(denom, blockNum))
}

// DeleteMedian deletes a given denom's median price at a given block.
func (k Keeper) DeleteMedian(
	ctx sdk.Context,
	denom string,
	blockNum uint64,
) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyMedian(denom, blockNum))
}

// DeleteMedianDeviation deletes a given denom's standard deviation around
// its median price at a given block.
func (k Keeper) DeleteMedianDeviation(
	ctx sdk.Context,
	denom string,
	blockNum uint64,
) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.KeyMedianDeviation(denom, blockNum))
}

// ClearMedians iterates through all medians in the store and deletes them.
func (k Keeper) ClearMedians(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixMedian)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}

// ClearMedianDeviations iterates through all median deviations in the store
// and deletes them.
func (k Keeper) ClearMedianDeviations(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.KeyPrefixMedianDeviation)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		store.Delete(iter.Key())
	}
}
