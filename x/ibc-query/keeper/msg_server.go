package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/cosmos/interchain-queries/x/ibc-query/types"
)

//var _ types.MsgServer = Keeper{}

// SubmitCrossChainQuery Handling SubmitCrossChainQuery transaction
func (k Keeper) SubmitCrossChainQuery(goCtx context.Context, msg *types.MsgSubmitCrossChainQuery) (*types.MsgSubmitCrossChainQueryResponse, error) {
	// UnwrapSDKContext
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentTimestamp := uint64(ctx.BlockTime().UnixNano())
	currentHeight := clienttypes.GetSelfHeight(ctx)

	// Sanity-check that localTimeoutHeight.
	if msg.LocalTimeoutHeight.RevisionHeight > 0 && msg.LocalTimeoutHeight.RevisionHeight <= currentHeight.RevisionHeight {
		return nil, sdkerrors.Wrapf(
			types.ErrTimeout,
			"localTimeoutHeight > 0 and current height >= localTimeoutHeight(%d >= %d)", currentHeight.RevisionHeight, msg.LocalTimeoutHeight.RevisionHeight,
		)
	}
	// Sanity-check that localTimeoutTimestamp
	if msg.LocalTimeoutStamp > 0 && msg.LocalTimeoutStamp <= currentTimestamp {
		return nil, sdkerrors.Wrapf(
			types.ErrTimeout,
			"localTimeoutTimestamp > 0 and current timestamp >= localTimeoutTimestamp(%d >= %d)", currentTimestamp, msg.LocalTimeoutStamp,
		)
	}

	// call keeper function
	// keeper func save query in private store
	query := types.CrossChainQuery{
		Id:                    msg.Id,
		Path:                  msg.Path,
		LocalTimeoutHeight:    msg.LocalTimeoutHeight,
		LocalTimeoutTimestamp: msg.LocalTimeoutStamp,
		QueryHeight:           msg.QueryHeight,
		ClientId:              msg.Sender,
	}

	k.SetCrossChainQuery(ctx, query)

	if err := k.SendQuery(ctx, msg.SourcePort, msg.SourceChannel, query); err != nil {
		return nil, err
	}

	// Log the query request
	k.Logger(ctx).Info("query sent", "query_id", msg.GetId())

	// emit event
	EmitQueryEvent(ctx, msg)

	return &types.MsgSubmitCrossChainQueryResponse{QueryId: query.Id}, nil
}

func (k Keeper) SubmitCrossChainQueryResult(goCtx context.Context, msg *types.MsgSubmitCrossChainQueryResult) (*types.MsgSubmitCrossChainQueryResultResponse, error) {
	// UnwrapSDKContext
	ctx := sdk.UnwrapSDKContext(goCtx)

	queryResult := types.CrossChainQueryResult{
		Id:     msg.Id,
		Result: msg.Result,
		Data:   msg.Data,
	}

	query, found := k.GetCrossChainQuery(ctx, queryResult.Id)
	// if CrossChainQuery of queryId doesn't exist in store, other relayer already submitted CrossChainQueryResult
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCrossChainQueryNotFound, "cannot find ICS-31 cross chain query id %s", queryResult.Id)
	}

	k.DeleteCrossChainQuery(ctx, queryResult.Id)

	// check Timeout by comparing the latest height of chain, latest timestamp
	selfHeight := clienttypes.GetSelfHeight(ctx)
	selfBlockTime := uint64(ctx.BlockTime().UnixNano())
	if selfHeight.GTE(query.LocalTimeoutHeight) {
		queryResult.Result = types.QueryResult_QUERY_RESULT_TIMEOUT
	}
	if selfBlockTime >= query.LocalTimeoutTimestamp {
		queryResult.Result = types.QueryResult_QUERY_RESULT_TIMEOUT
	}

	// store result in privateStore
	k.SetCrossChainQueryResult(ctx, queryResult)

	return &types.MsgSubmitCrossChainQueryResultResponse{}, nil
}
