package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/cosmos/interchain-queries/testutil/keeper"
	ibcquerykeeper "github.com/cosmos/interchain-queries/x/ibc_query/keeper"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
	"github.com/stretchr/testify/require"
)

// TestCrossChainQueryResult validates CrossChainQueryResult against the spec.
// Retrieve the query result from the private store using the query's identifier.
func TestCrossChainQueryResult(t *testing.T) {
    var (
		queryResult    types.MsgSubmitCrossChainQueryResult
		IBCQueryKeeper ibcquerykeeper.Keeper
		ctx            sdk.Context
		mocks          testkeeper.MockedKeepers
    )

    testCases := []struct {
        msg          string
		expPass      bool
        setUp        func() (types.MsgSubmitCrossChainQueryResult)
		queryMsg     types.QueryCrossChainQueryResult
    }{
        {
            "success: query valid query id", true,
			func () (types.MsgSubmitCrossChainQueryResult) {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				IBCQueryKeeper, ctx, mocks, queryResult = testkeeper.GetSubmitCrossChainQueryResultKeeperAndCtx(IBCQueryKeeper, ctx, mocks)
				return queryResult
			},
			types.QueryCrossChainQueryResult{
				Id: "query-0",
			},
        },
		{
            "success: query valid query id", true,
			func () (types.MsgSubmitCrossChainQueryResult) {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				testkeeper.GetSubmitCrossChainQueryResultKeeperAndCtx(IBCQueryKeeper, ctx, mocks)
				IBCQueryKeeper, ctx, mocks, queryResult = testkeeper.GetSubmitCrossChainQueryResultKeeperAndCtx(IBCQueryKeeper, ctx, mocks)
				return queryResult
			},
			types.QueryCrossChainQueryResult{
				Id: "query-1",
			},
        },
		{
            "fail: query invalid query id", false,
			func () (types.MsgSubmitCrossChainQueryResult) {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				IBCQueryKeeper, ctx, mocks, queryResult = testkeeper.GetSubmitCrossChainQueryResultKeeperAndCtx(IBCQueryKeeper, ctx, mocks)
				return queryResult
			},
			types.QueryCrossChainQueryResult{
				Id: "query-1",
			},
        },
    }

    for _, tc := range testCases {
		queryResult := tc.setUp()
		queryResultRes, err := IBCQueryKeeper.CrossChainQueryResult(ctx, &tc.queryMsg)

		if tc.expPass {
			require.NoError(t, err)
			require.Equal(t, queryResult.Id, queryResultRes.Id)
			require.Equal(t, queryResult.Data, queryResultRes.Data)
			require.Equal(t, queryResult.Result, queryResultRes.Result)
		} else {
			require.Error(t, err)
		}
        
    }
}

