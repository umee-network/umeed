package keeper

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibcexported "github.com/cosmos/ibc-go/v5/modules/core/exported"

	"github.com/umee-network/umee/v4/x/uibc"
)

// SendPacket wraps IBC ChannelKeeper's SendPacket function
func (k Keeper) SendPacket(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI) error {
	funds, denom, err := k.GetFundsFromPacket(packet)
	if err != nil {
		return sdkerrors.Wrap(err, "bad packet in rate limit's SendPacket")
	}

	amount, ok := sdkmath.NewIntFromString(funds)
	if !ok {
		return uibc.ErrInvalidAmount.Wrapf("amount %s", funds)
	}

	if err := k.CheckAndUpdateQuota(ctx, denom, amount); err != nil {
		return sdkerrors.Wrap(err, "bad packet in rate limit's SendPacket")
	}

	return k.ics4Wrapper.SendPacket(ctx, chanCap, packet)
}

// WriteAcknowledgement wraps IBC ChannelKeeper's WriteAcknowledgement function
// ICS29 WriteAcknowledgement is used for asynchronous acknowledgements
func (k Keeper) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet ibcexported.PacketI,
	acknowledgement ibcexported.Acknowledgement,
) error {
	// ics4Wrapper may be core IBC or higher-level middleware
	return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, acknowledgement)
}

// GetAppVersion returns the underlying application version.
func (k Keeper) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return k.ics4Wrapper.GetAppVersion(ctx, portID, channelID)
}
