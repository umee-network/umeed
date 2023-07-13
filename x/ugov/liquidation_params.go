package ugov

import (
	fmt "fmt"
	time "time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	appparams "github.com/umee-network/umee/v5/app/params"
)

var (
	DefaultMaxSupply              = sdk.NewCoin(appparams.BondDenom, sdk.NewInt(12_000000000))
	DefaultInflationCycleDuration = time.Second * time.Duration(60*60*24*365)
	DefaultInflationReductionRate = sdk.MustNewDecFromStr("0.01")
)

func DefaultLiquidationParams() LiquidationParams {
	return LiquidationParams{
		MaxSupply:              DefaultMaxSupply,
		InflationCycleDuration: DefaultInflationCycleDuration,
		InflationReductionRate: DefaultInflationReductionRate,
	}
}

func (lp LiquidationParams) Validate() error {
	if lp.MaxSupply.Amount.LT(math.NewInt(0)) {
		return fmt.Errorf("%s must be not negative: %s", "max_supply", lp.MaxSupply.Amount.String())
	}

	if lp.InflationReductionRate.LT(sdk.ZeroDec()) {
		return fmt.Errorf("%s must be not negative: %s", "inflation reduction rate ", lp.InflationReductionRate.String())
	}

	return nil
}
