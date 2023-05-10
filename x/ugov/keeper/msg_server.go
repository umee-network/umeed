package keeper

import (
	"context"

	"github.com/umee-network/umee/v4/util/sdkutil"
	"github.com/umee-network/umee/v4/x/ugov"
)

type msgServer struct {
	kb Builder
}

// NewMsgServer returns an implementation of uibc.MsgServer
func NewMsgServer(kb Builder) ugov.MsgServer {
	return msgServer{kb: kb}
}

// GovUpdateMinFees sets protocol controlled tx min fees.
func (m msgServer) GovUpdateMinGasPrice(ctx context.Context, msg *ugov.MsgGovUpdateMinGasPrice) (*ugov.MsgGovUpdateMinGasPriceResponse, error) {
	sdkCtx, err := sdkutil.StartMsg(ctx, msg)
	if err != nil {
		return nil, err
	}

	k := m.kb.Keeper(&sdkCtx)
	k.SetMinGasPrice(&msg.MinGasPrice)
	return &ugov.MsgGovUpdateMinGasPriceResponse{}, nil
}
