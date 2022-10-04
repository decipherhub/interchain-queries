package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
)

// OnRecvPacket IBCQueryModule doesn't receive packet
func (k Keeper) OnRecvPacket(ctx sdk.Context, packet channeltypes.Packet) error {
	return nil
}

func (k Keeper) OnTimeoutPacket(ctx sdk.Context) error {
	return sdkerrors.Wrapf(types.ErrTimeout, "query packet timeout")
}
