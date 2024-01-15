package keeper

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/umee-network/umee/v6/util/coin"
	"github.com/umee-network/umee/v6/x/metoken"
	otypes "github.com/umee-network/umee/v6/x/oracle/types"
)

var usdExponent = uint32(0)

// Prices calculates meToken price as an avg of median prices of all index accepted assets.
func (k Keeper) Prices(index metoken.Index) (metoken.IndexPrices, error) {
	indexPrices := metoken.EmptyIndexPrices(index)

	supply, err := k.IndexBalances(index.Denom)
	if err != nil {
		return indexPrices, err
	}

	allPrices := k.oracleKeeper.AllMedianPrices(*k.ctx)

	// calculate the total assets value in the index balances
	totalAssetsUSDValue := sdkmath.LegacyZeroDec()
	for _, aa := range index.AcceptedAssets {
		// get token settings from leverageKeeper to use the symbol_denom
		tokenSettings, err := k.leverageKeeper.GetTokenSettings(*k.ctx, aa.Denom)
		if err != nil {
			return indexPrices, err
		}

		assetPrice, err := latestPrice(allPrices, tokenSettings.SymbolDenom)
		if err != nil {
			return indexPrices, err
		}

		indexPrices.SetPrice(
			metoken.AssetPrice{
				BaseDenom:   aa.Denom,
				SymbolDenom: tokenSettings.SymbolDenom,
				Price:       assetPrice,
				Exponent:    tokenSettings.Exponent,
			},
		)

		// if no meTokens were minted, the totalAssetValue is the sum of all the assets prices.
		// otherwise is the sum of the value of all the assets in the index.
		if supply.MetokenSupply.IsZero() {
			totalAssetsUSDValue = totalAssetsUSDValue.Add(assetPrice)
		} else {
			balance, i := supply.AssetBalance(aa.Denom)
			if i < 0 {
				return indexPrices, sdkerrors.ErrNotFound.Wrapf("balance for denom %s not found", aa.Denom)
			}

			assetUSDValue, err := valueInUSD(balance.AvailableSupply(), assetPrice, tokenSettings.Exponent)
			if err != nil {
				return indexPrices, err
			}
			totalAssetsUSDValue = totalAssetsUSDValue.Add(assetUSDValue)
		}
	}

	if supply.MetokenSupply.IsZero() {
		// if no meTokens were minted, the meTokenPrice is totalAssetsUSDValue divided by accepted assets quantity
		indexPrices.Price = totalAssetsUSDValue.QuoInt(sdkmath.NewInt(int64(len(index.AcceptedAssets))))
	} else {
		// otherwise, the meTokenPrice is totalAssetsUSDValue divided by meTokens minted quantity
		meTokenPrice, err := priceInUSD(supply.MetokenSupply.Amount, totalAssetsUSDValue, index.Exponent)
		if err != nil {
			return indexPrices, err
		}
		indexPrices.Price = meTokenPrice
	}

	for i := 0; i < len(indexPrices.Assets); i++ {
		asset := indexPrices.Assets[i]
		swapRate, err := metoken.Rate(asset.Price, indexPrices.Price, asset.Exponent, indexPrices.Exponent)
		if err != nil {
			return indexPrices, err
		}
		swapFee, _, err := k.swapFee(index, indexPrices, coin.One(asset.BaseDenom))
		if err != nil {
			return indexPrices, err
		}

		redeemRate, err := metoken.Rate(indexPrices.Price, asset.Price, indexPrices.Exponent, asset.Exponent)
		if err != nil {
			return indexPrices, err
		}
		redeemFee, _, err := k.redeemFee(index, indexPrices, coin.One(asset.BaseDenom))
		if err != nil {
			return indexPrices, err
		}

		indexPrices.Assets[i].SwapRate = swapRate
		indexPrices.Assets[i].SwapFee = swapFee
		indexPrices.Assets[i].RedeemRate = redeemRate
		indexPrices.Assets[i].RedeemFee = redeemFee
	}

	return indexPrices, nil
}

// SetPricesToOracle of every registered index.
func (k Keeper) SetPricesToOracle() error {
	indexes := k.GetAllRegisteredIndexes()
	for _, index := range indexes {
		iPrice, err := k.Prices(index)
		if err != nil {
			k.Logger().Error(
				"setting price to oracle: couldn't calculate the price",
				"denom", index.Denom,
				"error", err.Error(),
				"block_time", k.ctx.BlockTime(),
			)
			continue
		}

		indexToken, err := k.leverageKeeper.GetTokenSettings(*k.ctx, index.Denom)
		if err != nil {
			k.Logger().Error(
				"setting price to oracle: couldn't get token settings",
				"denom", index.Denom,
				"error", err.Error(),
				"block_time", k.ctx.BlockTime(),
			)
			continue
		}

		k.oracleKeeper.SetExchangeRate(*k.ctx, indexToken.SymbolDenom, iPrice.Price)
	}

	return nil
}

// latestPrice from the list of medians, based on the block number.
func latestPrice(prices otypes.Prices, symbolDenom string) (sdkmath.LegacyDec, error) {
	latestPrice := otypes.Price{}
	for _, price := range prices {
		if price.ExchangeRateTuple.Denom == symbolDenom && price.BlockNum > latestPrice.BlockNum {
			latestPrice = price
		}
	}

	if latestPrice.BlockNum == 0 {
		return sdkmath.LegacyDec{}, fmt.Errorf("price not found in oracle for denom %s", symbolDenom)
	}

	return latestPrice.ExchangeRateTuple.ExchangeRate, nil
}

// valueInUSD given a specific amount, price and exponent
func valueInUSD(amount sdkmath.Int, assetPrice sdkmath.LegacyDec, assetExponent uint32) (sdkmath.LegacyDec, error) {
	exponentFactor, err := metoken.ExponentFactor(assetExponent, usdExponent)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}
	return exponentFactor.MulInt(amount).Mul(assetPrice), nil
}

// priceInUSD given a specific amount, totalValue and exponent
func priceInUSD(amount sdkmath.Int, totalValue sdkmath.LegacyDec, assetExponent uint32) (sdkmath.LegacyDec, error) {
	exponentFactor, err := metoken.ExponentFactor(assetExponent, usdExponent)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}

	return totalValue.Quo(exponentFactor.MulInt(amount)), nil
}
