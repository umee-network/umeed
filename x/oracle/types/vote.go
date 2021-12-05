package types

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewAggregateExchangeRatePrevote(
	hash AggregateVoteHash,
	voter sdk.ValAddress,
	submitBlock uint64,
) AggregateExchangeRatePrevote {

	return AggregateExchangeRatePrevote{
		Hash:        hash.String(),
		Voter:       voter.String(),
		SubmitBlock: submitBlock,
	}
}

// String implement stringify
func (v AggregateExchangeRatePrevote) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

func NewAggregateExchangeRateVote(
	exchangeRateTuples ExchangeRateTuples,
	voter sdk.ValAddress,
) AggregateExchangeRateVote {

	return AggregateExchangeRateVote{
		ExchangeRateTuples: exchangeRateTuples,
		Voter:              voter.String(),
	}
}

// String implement stringify
func (v AggregateExchangeRateVote) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

// NewExchangeRateTuple creates a ExchangeRateTuple instance
func NewExchangeRateTuple(denom string, exchangeRate sdk.Dec) ExchangeRateTuple {
	return ExchangeRateTuple{
		denom,
		exchangeRate,
	}
}

// String implement stringify
func (v ExchangeRateTuple) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

// ExchangeRateTuples - array of ExchangeRateTuple
type ExchangeRateTuples []ExchangeRateTuple

// String implements fmt.Stringer interface
func (tuples ExchangeRateTuples) String() string {
	out, _ := yaml.Marshal(tuples)
	return string(out)
}

// ParseExchangeRateTuples ExchangeRateTuple parser
func ParseExchangeRateTuples(tuplesStr string) (ExchangeRateTuples, error) {
	tuplesStr = strings.TrimSpace(tuplesStr)
	if len(tuplesStr) == 0 {
		return nil, nil
	}

	tupleStrs := strings.Split(tuplesStr, ",")
	tuples := make(ExchangeRateTuples, len(tupleStrs))

	duplicateCheckMap := make(map[string]bool)
	for i, tupleStr := range tupleStrs {
		// Get denom and amount
		denomAmountStr := strings.Split(tupleStr, ":")
		decCoin := sdk.MustNewDecFromStr(denomAmountStr[1])
		denom := denomAmountStr[0]

		tuples[i] = ExchangeRateTuple{
			Denom:        denom,
			ExchangeRate: decCoin,
		}

		if _, ok := duplicateCheckMap[denom]; ok {
			return nil, fmt.Errorf("duplicated denom %s", denom)
		}

		duplicateCheckMap[denom] = true
	}

	return tuples, nil
}
