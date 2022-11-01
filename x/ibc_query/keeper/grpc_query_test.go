package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	testkeeper "github.com/cosmos/interchain-queries/testutil/keeper"
	ibcquerykeeper "github.com/cosmos/interchain-queries/x/ibc_query/keeper"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// TestCrossChainQueryResult validates CrossChainQueryResult against the spec.
// Retrieve the query result from the private store using the query's identifier.
func TestCrossChainQueryResult(t *testing.T) {
	_, _, senderAddr := testdata.KeyTestPubAddr()
	
    var (
		msgSubmitQuery  types.MsgSubmitCrossChainQuery
		msgQueryResult  types.MsgSubmitCrossChainQueryResult
		IBCQueryKeeper  ibcquerykeeper.Keeper
		ctx             sdk.Context
		mocks           testkeeper.MockedKeepers
    )

    testCases := []struct {
        msg             string
		expPass         bool
        setUp           func()
		queryMsg        types.QueryCrossChainQueryResult
		expQueryResult  types.QueryCrossChainQueryResultResponse
    }{
        {
            "success: query valid query id", true,
			func () {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr.String())
				msgQueryResult = testkeeper.GetMsgSubmitCrossChainQueryResult(ctx, "query-0", senderAddr.String())
			},
			types.QueryCrossChainQueryResult{
				Id: "query-0",
			},
			types.QueryCrossChainQueryResultResponse{
				Id:       "query-0",
				Result:   types.QueryResult_QUERY_RESULT_SUCCESS,
				Data:     []byte("test result data"),
			},
        },
		{
            "success: query valid query id", true,
			func () {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr.String())
				msgQueryResult = testkeeper.GetMsgSubmitCrossChainQueryResult(ctx, "query-0", senderAddr.String())
				gomock.InOrder(testkeeper.GetMocksForSubmitCrossChainQuery(ctx, &mocks, "query-0", msgSubmitQuery.Sender)...)
				IBCQueryKeeper.SubmitCrossChainQuery(ctx, &msgSubmitQuery)   
				IBCQueryKeeper.SubmitCrossChainQueryResult(ctx, &msgQueryResult)

				msgQueryResult = testkeeper.GetMsgSubmitCrossChainQueryResult(ctx, "query-1", senderAddr.String())
			},
			types.QueryCrossChainQueryResult{
				Id: "query-1",
			},
			types.QueryCrossChainQueryResultResponse{
				Id:       "query-1",
				Result:   types.QueryResult_QUERY_RESULT_SUCCESS,
				Data:     []byte("test result data"),
			},
        },
		{
            "fail: query invalid query id", false,
			func () {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr.String())
				msgQueryResult = testkeeper.GetMsgSubmitCrossChainQueryResult(ctx, "query-0", senderAddr.String())
			},
			types.QueryCrossChainQueryResult{
				Id: "query-1",
			},
			types.QueryCrossChainQueryResultResponse{},
        },
    }

    for _, tc := range testCases {
		tc.setUp()

		// Run SubmitCrossChainQuery
		gomock.InOrder(testkeeper.GetMocksForSubmitCrossChainQuery(ctx, &mocks, testkeeper.GetQueryId(IBCQueryKeeper, ctx), msgSubmitQuery.Sender)...)
		IBCQueryKeeper.SubmitCrossChainQuery(ctx, &msgSubmitQuery)
		IBCQueryKeeper.SubmitCrossChainQueryResult(ctx, &msgQueryResult)   

		// Run query CrossChainQueryResult
		queryResultRes, err := IBCQueryKeeper.CrossChainQueryResult(ctx, &tc.queryMsg)
		if tc.expPass {
			require.NoError(t, err)
			require.Equal(t, msgQueryResult.Id, queryResultRes.Id)
			require.Equal(t, msgQueryResult.Data, queryResultRes.Data)
			require.Equal(t, msgQueryResult.Result, queryResultRes.Result)
		} else {
			require.Error(t, err)
		}
        
    }
}

