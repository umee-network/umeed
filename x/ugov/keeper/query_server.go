package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/umee-network/umee/v6/x/ugov"
)

var _ ugov.QueryServer = Querier{}

// Querier implements a QueryServer for the x/uibc module.
type Querier struct {
	Builder
}

func NewQuerier(kb Builder) Querier {
	return Querier{kb}
}

// MinTxFees returns minimum transaction fees.
func (q Querier) MinGasPrice(ctx context.Context, _ *ugov.QueryMinGasPrice) (*ugov.QueryMinGasPriceResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return &ugov.QueryMinGasPriceResponse{MinGasPrice: q.Keeper(&sdkCtx).MinGasPrice()}, nil
}

func (q Querier) EmergencyGroup(ctx context.Context, _ *ugov.QueryEmergencyGroup,
) (*ugov.QueryEmergencyGroupResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return &ugov.QueryEmergencyGroupResponse{
			EmergencyGroup: q.Keeper(&sdkCtx).EmergencyGroup().String(),
		},
		nil
}

// InflationParams returns inflation rate change params
func (q Querier) InflationParams(ctx context.Context, _ *ugov.QueryInflationParams) (
	*ugov.QueryInflationParamsResponse, error,
) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	return &ugov.QueryInflationParamsResponse{Params: q.Keeper(&sdkCtx).InflationParams()}, nil
}

// InflationCycleEnd return when the inflation cycle will be ended.
func (q Querier) InflationCycleEnd(ctx context.Context, _ *ugov.QueryInflationCycleEnd) (
	*ugov.QueryInflationCycleEndResponse, error,
) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	cycleEndTime := q.Keeper(&sdkCtx).InflationCycleEnd()
	return &ugov.QueryInflationCycleEndResponse{End: &cycleEndTime}, nil
}

// TokenBalances implements ugov.QueryServer.
func (q Querier) TokenBalances(ctx context.Context, req *ugov.QueryTokenBalances) (*ugov.QueryTokenBalancesResponse,
	error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	if req.Height != 0 {
		sdkCtx = sdkCtx.WithBlockHeight(req.Height)
	}
	resp, err := q.BankKeeper.DenomOwners(sdk.WrapSDKContext(sdkCtx), &banktypes.QueryDenomOwnersRequest{
		Denom:      req.Denom,
		Pagination: req.Pagination,
	})
	if err != nil {
		return nil, err
	}

	denomsOwners := make([]*ugov.DenomOwner, 0)
	for _, v := range resp.DenomOwners {
		denomsOwners = append(denomsOwners, &ugov.DenomOwner{
			Address: v.Address,
			Balance: v.Balance,
		})
	}

	return &ugov.QueryTokenBalancesResponse{Pagination: resp.Pagination, DenomOwners: denomsOwners}, nil
}
