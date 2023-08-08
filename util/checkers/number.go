package checkers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/constraints"
)

func IntegerMaxDiff[T constraints.Integer](a, b, maxDiff T, note string) error {
	var diff T
	if a > b {
		diff = a - b
	} else {
		diff = b - a
	}
	if diff > maxDiff {
		return fmt.Errorf("%s, diff (=%v) is too big", note, diff)
	}
	return nil
}

func DecMaxDiff(a, b, maxDiff sdk.Dec, note string) error {
	diff := a.Sub(b).Abs()
	if diff.GT(maxDiff) {
		return fmt.Errorf("%s, diff (=%v) is too big", note, diff)
	}
	return nil
}
