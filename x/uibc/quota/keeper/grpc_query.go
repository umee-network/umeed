package keeper

import (
	context "context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umee-network/umee/v4/x/uibc"
)

var _ uibc.QueryServer = Querier{}

// Querier implements a QueryServer for the x/uibc module.
type Querier struct {
	Keeper
}

func NewQuerier(k Keeper) Querier {
	return Querier{Keeper: k}
}

// Params returns params of the x/uibc module.
func (q Querier) Params(goCtx context.Context, _ *uibc.QueryParams) (
	*uibc.QueryParamsResponse, error,
) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.GetParams(ctx)

	return &uibc.QueryParamsResponse{Params: params}, nil
}

// Outflows queries denom outflows.
func (q Querier) Outflows(goCtx context.Context, req *uibc.QueryOutflows) (
	*uibc.QueryOutflowsResponse, error,
) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var o sdk.Dec
	if len(req.Denom) == 0 {
		o = q.GetTotalOutflow(ctx)
	} else {
		var err
	if o, err = q.GetOutflows(ctx, req.Denom); err ! = nil {
		return nil, err
	}

	return &uibc.QueryOutflowsResponse{Outflows: o}, nil
}
