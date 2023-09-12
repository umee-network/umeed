package metoken

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v6/util/coin"
)

func EmptyIndexPrices(index Index) IndexPrices {
	return IndexPrices{
		Denom:    index.Denom,
		Price:    sdk.Dec{},
		Exponent: index.Exponent,
		Assets:   make([]AssetPrice, 0),
	}
}

// PriceByBaseDenom returns a AssetPrice given a specific base_denom.
func (ip IndexPrices) PriceByBaseDenom(denom string) (AssetPrice, error) {
	for _, ap := range ip.Assets {
		if ap.BaseDenom == denom {
			return ap, nil
		}
	}
	return AssetPrice{}, sdkerrors.ErrNotFound.Wrapf("price not found for denom %s", denom)
}

// SetPrice to the IndexPrices.
func (ip *IndexPrices) SetPrice(price AssetPrice) {
	for i, ap := range ip.Assets {
		if ap.BaseDenom == price.BaseDenom {
			ip.Assets[i] = price
			return
		}
	}

	ip.Assets = append(ip.Assets, price)

}

// SwapRate converts an asset to meToken applying exchange_rate and normalizing the exponent.
func (ip IndexPrices) SwapRate(from sdk.Coin) (sdkmath.Int, error) {
	fromPrice, err := ip.PriceByBaseDenom(from.Denom)
	if err != nil {
		return sdkmath.Int{}, err
	}

	exchangeRate := fromPrice.Price.Quo(ip.Price)

	exponentFactor, err := ExponentFactor(fromPrice.Exponent, ip.Exponent)
	if err != nil {
		return sdkmath.Int{}, err
	}

	return exchangeRate.MulInt(from.Amount).Mul(exponentFactor).TruncateInt(), nil
}

// RedeemRate converts meToken to an asset applying exchange_rate and normalizing the exponent.
func (ip IndexPrices) RedeemRate(from sdk.Coin, to string) (sdkmath.Int, error) {
	toPrice, err := ip.PriceByBaseDenom(to)
	if err != nil {
		return sdkmath.Int{}, err
	}

	exchangeRate := ip.Price.Quo(toPrice.Price)

	exponentFactor, err := ExponentFactor(ip.Exponent, toPrice.Exponent)
	if err != nil {
		return sdkmath.Int{}, err
	}

	return exchangeRate.MulInt(from.Amount).Mul(exponentFactor).TruncateInt(), nil
}

// QueryExport completes the structure with missing data for the query.
func (ip IndexPrices) QueryExport() IndexPrices {
	assets := make([]AssetPrice, len(ip.Assets))
	for i := 0; i < len(ip.Assets); i++ {
		asset := ip.Assets[i]
		asset.SwapRate = asset.Price.Quo(ip.Price)
		asset.RedeemRate = ip.Price.Quo(asset.Price)
		assets[i] = asset
	}

	return IndexPrices{
		Denom:    ip.Denom,
		Price:    ip.Price,
		Exponent: ip.Exponent,
		Assets:   assets,
	}
}

// ExponentFactor calculates the factor to multiply by which the assets with different exponents.
// If there is no such difference the result will be 1.
func ExponentFactor(initialExponent, resultExponent uint32) (sdk.Dec, error) {
	exponentDiff := int(resultExponent) - int(initialExponent)
	multiplier, ok := coin.Exponents[exponentDiff]
	if !ok {
		return sdk.Dec{}, fmt.Errorf("multiplier not found for exponentDiff %d", exponentDiff)
	}

	return multiplier, nil
}
