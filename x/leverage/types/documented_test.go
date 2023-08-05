package types_test

import (
	"testing"

	"gotest.tools/v3/assert"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v5/util/coin"
	"github.com/umee-network/umee/v5/x/leverage/types"
)

func TestMaxBorrowScenarioA(t *testing.T) {
	// This scenario reproduces the final table of "MaxBorrow Scenario A" from x/leverage/EXAMPLES.md
	borrowPosition := types.NewAccountPosition(
		[]types.Token{
			testToken("AAAA", "0.4", "0.5"),
			testToken("BBBB", "0.3", "0.5"),
			testToken("CCCC", "0.2", "0.5"),
			testToken("DDDD", "0.1", "0.5"),
		},
		[]types.SpecialAssetPair{
			testPair("AAAA", "BBBB", "0.5", "0.5"),
			testPair("BBBB", "AAAA", "0.5", "0.5"),

			testPair("AAAA", "CCCC", "0.4", "0.4"),
			testPair("CCCC", "AAAA", "0.4", "0.4"),
		},
		sdk.NewDecCoins(
			coin.Dec("AAAA", "100"),
			coin.Dec("DDDD", "300"),
		),
		sdk.NewDecCoins(
			coin.Dec("BBBB", "55"),
			coin.Dec("CCCC", "20"),
			coin.Dec("DDDD", "5"),
		),
		false,
	)
	assert.Equal(t,
		"special:\n"+
			"  100.000000000000000000AAAA, 50.000000000000000000BBBB, 0.500000000000000000\n"+
			"  0.000000000000000000BBBB, 0.000000000000000000AAAA, 0.500000000000000000\n"+
			"  0.000000000000000000CCCC, 0.000000000000000000AAAA, 0.400000000000000000\n"+
			"normal:\n"+
			"  {50.000000000000000000DDDD 0.100000000000000000}, {5.000000000000000000BBBB 0.300000000000000000}\n"+
			"  {200.000000000000000000DDDD 0.100000000000000000}, {20.000000000000000000CCCC 0.200000000000000000}\n"+
			"  {50.000000000000000000DDDD 0.100000000000000000}, {5.000000000000000000DDDD 0.100000000000000000}\n",
		borrowPosition.String(),
	)
}
