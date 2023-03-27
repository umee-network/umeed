package ics20

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ibctransfer "github.com/cosmos/ibc-go/v5/modules/apps/transfer"
	ibctransfertypes "github.com/cosmos/ibc-go/v5/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"

	ibcutil "github.com/umee-network/umee/v4/util/ibc"
	ltypes "github.com/umee-network/umee/v4/x/leverage/types"
	"github.com/umee-network/umee/v4/x/uibc/ics20/keeper"
)

// IBCModule embeds the ICS-20 transfer IBCModule where we only override specific
// methods.
type IBCModule struct {
	// leverage keeper
	lkeeper keeper.LeverageKeeper
	// embed the ICS-20 transfer's AppModule
	ibctransfer.IBCModule

	keeper keeper.Keeper
}

func NewIBCModule(leverageKeeper keeper.LeverageKeeper, am ibctransfer.IBCModule, k keeper.Keeper) IBCModule {
	return IBCModule{
		lkeeper:   leverageKeeper,
		IBCModule: am,
		keeper:    k,
	}
}

// OnRecvPacket delegates the OnRecvPacket call to the embedded ICS-20 transfer
// IBCModule and updates metadata if successful.
func (am IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	// Allowing only registered token
	var data ibctransfertypes.FungibleTokenPacketData
	if err := ibctransfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		ackErr := sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot unmarshal ICS-20 transfer packet data")
		return channeltypes.NewErrorAcknowledgement(ackErr)
	}

	_, denom, err := ibcutil.GetFundsFromPacket(packet)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	_, err = am.lkeeper.GetTokenSettings(ctx, denom)
	if err != nil && ltypes.ErrNotRegisteredToken.Is(err) {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	ack := am.IBCModule.OnRecvPacket(ctx, packet, relayer)
	if ack.Success() {
		var data ibctransfertypes.FungibleTokenPacketData
		if err := ibctransfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err == nil {
			// track metadata
			am.keeper.PostOnRecvPacket(ctx, packet, data)
		}
	}

	return ack
}
