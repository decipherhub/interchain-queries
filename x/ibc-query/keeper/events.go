package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/interchain-queries/x/ibc-query/types"
)

// EmitCreateClientEvent emits a create client event
func EmitQueryEvent(ctx sdk.Context, query *types.MsgSubmitCrossChainQuery) {
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventSendQuery,
			sdk.NewAttribute(types.AttributeKeyQueryID, query.GetId()),
			sdk.NewAttribute(types.AttributeKeyQueryHeight, fmt.Sprintf("%d", query.GetQueryHeight())),
			sdk.NewAttribute(types.AttributeKeyQueryPath, fmt.Sprintf("%x", query.GetQueryPath())),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	})
}
