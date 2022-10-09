package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
)

// SubmitCrossChainQuery Handling SubmitCrossChainQuery transaction
func (k Keeper) SubmitCrossChainQuery(goCtx context.Context, msg *types.MsgSubmitCrossChainQuery) (*types.MsgSubmitCrossChainQueryResponse, error) {
	// UnwrapSDKContext
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentTimestamp := uint64(ctx.BlockTime().UnixNano())
	currentHeight := clienttypes.GetSelfHeight(ctx)

	if msg.LocalTimeoutHeight.RevisionHeight == 0 && msg.LocalTimeoutStamp == 0 {
		return nil, types.ErrInvalidQuery
	}

	// Sanity-check that localTimeoutHeight.
	if msg.LocalTimeoutHeight.RevisionHeight > 0 && msg.LocalTimeoutHeight.RevisionHeight <= currentHeight.RevisionHeight {
		return nil, errorsmod.Wrapf(
			types.ErrTimeout,
			"localTimeoutHeight > 0 and current height >= localTimeoutHeight(%d >= %d)", currentHeight.RevisionHeight, msg.LocalTimeoutHeight.RevisionHeight,
		)
	}
	// Sanity-check that localTimeoutTimestamp
	if msg.LocalTimeoutStamp > 0 && msg.LocalTimeoutStamp <= currentTimestamp {
		return nil, errorsmod.Wrapf(
			types.ErrTimeout,
			"localTimeoutTimestamp > 0 and current timestamp >= localTimeoutTimestamp(%d >= %d)", currentTimestamp, msg.LocalTimeoutStamp,
		)
	}

	// call keeper function
	// keeper func save query in private store
	query := types.CrossChainQuery{
		Id:                    k.GenerateQueryIdentifier(ctx),
		Path:                  msg.Path,
		LocalTimeoutHeight:    msg.LocalTimeoutHeight,
		LocalTimeoutTimestamp: msg.LocalTimeoutStamp,
		QueryHeight:           msg.QueryHeight,
		ChainId:               msg.ChainId,
	}

	k.SetCrossChainQuery(ctx, query)

	// Log the query request
	k.Logger(ctx).Info("query sent", "query_id", query.GetId())

	queryCapability, err :=  k.scopedKeeper.NewCapability(ctx, query.Id)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "could not create capability for query ID %s", query.Id)
	}

	// emit event
	EmitQueryEvent(ctx, query)

	return &types.MsgSubmitCrossChainQueryResponse{Id: query.Id, CapKey: fmt.Sprint(queryCapability.Index)}, nil
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
		return nil, types.ErrCrossChainQueryNotFound
	}

	k.DeleteCrossChainQuery(ctx, queryResult.Id)

	// check Timeout by comparing the latest height of chain, latest timestamp
	selfHeight := clienttypes.GetSelfHeight(ctx)
	selfBlockTime := uint64(ctx.BlockTime().UnixNano())
	if query.LocalTimeoutHeight.RevisionHeight > 0 && selfHeight.GTE(query.LocalTimeoutHeight) {
		queryResult.Result = types.QueryResult_QUERY_RESULT_TIMEOUT
	}
	if query.LocalTimeoutTimestamp > 0 && selfBlockTime >= query.LocalTimeoutTimestamp {
		queryResult.Result = types.QueryResult_QUERY_RESULT_TIMEOUT
	}

	// store result in privateStore
	k.SetCrossChainQueryResult(ctx, queryResult)

	return &types.MsgSubmitCrossChainQueryResultResponse{}, nil
}

// SubmitPruneCrossChainQueryResult Handling SubmitPruneCrossChainQueryResult transaction
func (k Keeper) SubmitPruneCrossChainQueryResult(goCtx context.Context, msg *types.MsgSubmitPruneCrossChainQueryResult) (*types.MsgSubmitPruneCrossChainQueryResultResponse, error) {
	// UnwrapSDKContext
	ctx := sdk.UnwrapSDKContext(goCtx)

	queryResult, found := k.GetCrossChainQueryResult(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.ErrNotFound
	}

	if !k.scopedKeeper.AuthenticateCapability(ctx, msg.CapKey, msg.Id) {
		return nil, types.ErrInvalidCapability
	}

	k.DeleteCrossChainQueryResult(ctx, queryResult.Id)

	return &types.MsgSubmitPruneCrossChainQueryResultResponse{Id: queryResult.Id}, nil
}
