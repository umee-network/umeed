package types

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// VoteForTally is a convenience wrapper to reduce redundant lookup cost.
type VoteForTally struct {
	Denom        string
	ExchangeRate sdk.Dec
	Voter        sdk.ValAddress
	Power        int64
}

// NewVoteForTally returns a new VoteForTally instance.
func NewVoteForTally(rate sdk.Dec, denom string, voter sdk.ValAddress, power int64) VoteForTally {
	return VoteForTally{
		ExchangeRate: rate,
		Denom:        denom,
		Voter:        voter,
		Power:        power,
	}
}

// ExchangeRateBallot is a convenience wrapper around a ExchangeRateVote slice.
type ExchangeRateBallot []VoteForTally

// ToMap return organized exchange rate map by validator.
func (pb ExchangeRateBallot) ToMap() map[string]sdk.Dec {
	exchangeRateMap := make(map[string]sdk.Dec)
	for _, vote := range pb {
		if vote.ExchangeRate.IsPositive() {
			exchangeRateMap[vote.Voter.String()] = vote.ExchangeRate
		}
	}

	return exchangeRateMap
}

// ToCrossRate return cross_rate(base/exchange_rate) ballot.
func (pb ExchangeRateBallot) ToCrossRate(bases map[string]sdk.Dec) (cb ExchangeRateBallot) {
	for i := range pb {
		vote := pb[i]

		if exchangeRateRT, ok := bases[vote.Voter.String()]; ok && vote.ExchangeRate.IsPositive() {
			vote.ExchangeRate = exchangeRateRT.Quo(vote.ExchangeRate)
		} else {
			// If we can't get reference terra exchange rate, we just convert the vote as abstain vote
			vote.ExchangeRate = sdk.ZeroDec()
			vote.Power = 0
		}

		cb = append(cb, vote)
	}

	return
}

// Power returns the total amount of voting power in the ballot.
func (pb ExchangeRateBallot) Power() int64 {
	var totalPower int64
	for _, vote := range pb {
		totalPower += vote.Power
	}

	return totalPower
}

// WeightedMedian returns the median weighted by the power of the ExchangeRateVote.
// CONTRACT: The ballot must be sorted.
func (pb ExchangeRateBallot) WeightedMedian() sdk.Dec {
	totalPower := pb.Power()
	if pb.Len() > 0 {
		var pivot int64
		for _, v := range pb {
			votePower := v.Power

			pivot += votePower
			if pivot >= (totalPower / 2) {
				return v.ExchangeRate
			}
		}
	}

	return sdk.ZeroDec()
}

// StandardDeviation returns the standard deviation by the power of the ExchangeRateVote.
func (pb ExchangeRateBallot) StandardDeviation() (standardDeviation sdk.Dec, err error) {
	if len(pb) == 0 {
		return sdk.ZeroDec(), nil
	}

	median := pb.WeightedMedian()

	sum := sdk.ZeroDec()
	for _, v := range pb {
		deviation := v.ExchangeRate.Sub(median)
		sum = sum.Add(deviation.Mul(deviation))
	}

	variance := sum.QuoInt64(int64(len(pb)))

	floatNum, err := strconv.ParseFloat(variance.String(), 64)
	if err != nil {
		return sdk.Dec{}, err
	}

	floatNum = math.Sqrt(floatNum)
	standardDeviation, _ = sdk.NewDecFromStr(fmt.Sprintf("%f", floatNum))

	return
}

// Len implements sort.Interface
func (pb ExchangeRateBallot) Len() int {
	return len(pb)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (pb ExchangeRateBallot) Less(i, j int) bool {
	return pb[i].ExchangeRate.LT(pb[j].ExchangeRate)
}

// Swap implements sort.Interface.
func (pb ExchangeRateBallot) Swap(i, j int) {
	pb[i], pb[j] = pb[j], pb[i]
}

// BallotDenom is a convenience wrapper for setting rates deterministically.
type BallotDenom struct {
	Ballot ExchangeRateBallot
	Denom  string
}

// BallotMapToSlice returns an array of sorted exchange rate ballots.
func BallotMapToSlice(votes map[string]ExchangeRateBallot) (b []BallotDenom) {
	for denom, ballot := range votes {
		b = append(b, BallotDenom{
			Denom:  denom,
			Ballot: ballot,
		})
	}
	sort.Slice(b, func(i, j int) bool {
		return b[i].Denom < b[j].Denom
	})
	return b
}

// Claim is an interface that directs its rewards to an attached bank account.
type Claim struct {
	Power     int64
	Weight    int64
	WinCount  int64
	Recipient sdk.ValAddress
}

// NewClaim generates a Claim instance.
func NewClaim(power, weight, winCount int64, recipient sdk.ValAddress) Claim {
	return Claim{
		Power:     power,
		Weight:    weight,
		WinCount:  winCount,
		Recipient: recipient,
	}
}
