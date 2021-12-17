package keeper_test

import (
	"encoding/hex"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/x/oracle/types"
	oracletypes "github.com/umee-network/umee/x/oracle/types"
)

// GenerateSalt generates a random salt, size length/2,  as a HEX encoded string.
func GenerateSalt(length int) (string, error) {
	if length == 0 {
		return "", fmt.Errorf("failed to generate salt: zero length")
	}

	n := length / 2
	bytes := make([]byte, n)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func (s *IntegrationTestSuite) TestMsgServer_AggregateExchangeRatePrevote() {
	ctx := s.ctx

	exchangeRatesStr := "123.2:UMEE"
	salt, err := GenerateSalt(2)
	s.Require().NoError(err)

	hash := oracletypes.GetAggregateVoteHash(salt, exchangeRatesStr, valAddr)
	msg := &types.MsgAggregateExchangeRatePrevote{
		Hash:      hash.String(),
		Feeder:    addr.String(),
		Validator: valAddr.String(),
	}
	_, err = s.msgServer.AggregateExchangeRatePrevote(sdk.WrapSDKContext(ctx), msg)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestMsgServer_AggregateExchangeRateVote() {
	ctx := s.ctx

	exchangeRatesStr := "123.2:UMEE"
	salt, err := GenerateSalt(2)
	s.Require().NoError(err)

	hash := oracletypes.GetAggregateVoteHash(salt, exchangeRatesStr, valAddr)
	prevoteMsg := &types.MsgAggregateExchangeRatePrevote{
		Hash:      hash.String(),
		Feeder:    addr.String(),
		Validator: valAddr.String(),
	}
	_, err = s.msgServer.AggregateExchangeRatePrevote(sdk.WrapSDKContext(ctx), prevoteMsg)
	s.Require().NoError(err)

	voteMsg := &types.MsgAggregateExchangeRateVote{
		Feeder:        addr.String(),
		Validator:     valAddr.String(),
		Salt:          salt,
		ExchangeRates: exchangeRatesStr,
	}
	_, err = s.msgServer.AggregateExchangeRateVote(sdk.WrapSDKContext(ctx), voteMsg)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TestMsgServer_DelegateFeedConsent() {
	app, ctx := s.app, s.ctx

	feederAddr := sdk.AccAddress([]byte("addr________________"))
	feederAcc := app.AccountKeeper.NewAccountWithAddress(ctx, feederAddr)
	app.AccountKeeper.SetAccount(ctx, feederAcc)

	_, err := s.msgServer.DelegateFeedConsent(sdk.WrapSDKContext(ctx), &types.MsgDelegateFeedConsent{
		Operator: valAddr.String(),
		Delegate: feederAddr.String(),
	})
	s.Require().NoError(err)
}
