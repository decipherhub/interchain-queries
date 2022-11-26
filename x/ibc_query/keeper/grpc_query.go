package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) CrossChainQuery(c context.Context, req *types.QueryCrossChainQuery) (*types.QueryCrossChainQueryResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	query, found := k.GetCrossChainQuery(ctx, types.QueryPath(req.Id))
	if !found {
		return nil, sdkerrors.ErrNotFound
	}

	return &types.QueryCrossChainQueryResponse{Result: &query}, nil
}

// CrossChainQueryResult implements the Query/CrossChainQueryResult gRPC method
func (k Keeper) CrossChainQueryResult(c context.Context, req *types.QueryCrossChainQueryResult) (*types.QueryCrossChainQueryResultResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)
	queryResult, found := k.GetCrossChainQueryResult(ctx, types.ResultQueryPath(req.Id))
	if !found {
		return nil, sdkerrors.ErrNotFound
	}

	return &types.QueryCrossChainQueryResultResponse{Id: req.Id, Result: queryResult.Result, Data: queryResult.Data}, nil
}
