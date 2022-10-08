package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	testkeeper "github.com/cosmos/interchain-queries/testutil/keeper"
	ibcquerykeeper "github.com/cosmos/interchain-queries/x/ibc_query/keeper"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
	"github.com/stretchr/testify/require"
)



func TestSubmitCrossChainQuery(t *testing.T) {
	var (
		msg              types.MsgSubmitCrossChainQuery
		IBCQueryKeeper   ibcquerykeeper.Keeper
		ctx              sdk.Context
	)

	testCases := []struct {
		name     string
		expPass  bool
		malleate func()
	}{
		{
			"success",
			true,
			func() {
				IBCQueryKeeper, ctx, _, _ = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msg = types.MsgSubmitCrossChainQuery{
					Path: "test/query_path", 
					LocalTimeoutHeight: clienttypes.NewHeight(0, uint64(ctx.BlockHeight()+50)), 
					LocalTimeoutStamp: uint64(ctx.BlockTime().UnixNano() + 5000000), 
					QueryHeight: uint64(ctx.BlockHeight()), 
					ChainId: "ibc-query", 
					Sender: "testAddress", 
				}
			},
		},
	}
	
	for _, tc := range testCases {
		
		tc.malleate()
		res, err := IBCQueryKeeper.SubmitCrossChainQuery(ctx, &msg)
		if tc.expPass {
			require.NoError(t, err)
			require.NotNil(t, res)
			queryResult, found := IBCQueryKeeper.GetCrossChainQuery(ctx, res.QueryId)
			require.True(t, found)
			require.Equal(t, res.QueryId, queryResult.Id)
			require.Equal(t, msg.Path, queryResult.Path)
			require.Equal(t, msg.LocalTimeoutHeight.RevisionHeight, queryResult.LocalTimeoutHeight.RevisionHeight)
			require.Equal(t, msg.LocalTimeoutStamp, queryResult.LocalTimeoutTimestamp)
			require.Equal(t, msg.QueryHeight, queryResult.QueryHeight)
			require.Equal(t, msg.ChainId, queryResult.ChainId)
		}
	}
}