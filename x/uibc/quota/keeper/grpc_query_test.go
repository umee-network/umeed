package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/v4/x/uibc"
	"gotest.tools/v3/assert"
)

func TestGRPCQueryParams(t *testing.T) {
	s := initKeeperTestSuite(t)
	ctx, client := s.ctx, s.queryClient
	tests := []struct {
		name        string
		req         uibc.QueryParams
		errExpected bool
	}{
		{
			name:        "valid",
			req:         uibc.QueryParams{},
			errExpected: false,
		},
	}

	for _, tc := range tests {
		_, err := client.Params(ctx, &tc.req)
		if tc.errExpected {
			assert.Error(t, err, "")
		} else {
			assert.NilError(t, err)
		}
	}
}

func TestGRPCGetQuota(t *testing.T) {
	t.Parallel()
	suite := initKeeperTestSuite(t)
	ctx, client := suite.ctx, suite.queryClient
	tests := []struct {
		name   string
		req    uibc.QueryOutflows
		errMsg string
	}{
		{
			name:   "valid: total outflows",
			req:    uibc.QueryOutflows{},
			errMsg: "",
		}, {
			name:   "valid req: OutflowSum zero because ibc-transfer not hapeen",
			req:    uibc.QueryOutflows{Denom: "umee"},
			errMsg: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := client.Outflows(ctx, &tc.req)
			if tc.errMsg == "" {
				assert.NilError(t, err)
				// TODO: add more data to check non zero values
				assert.DeepEqual(t, sdk.NewDec(0), resp.Amount)
			} else {
				assert.Error(t, err, tc.errMsg)
			}
		})
	}
}
