package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	testkeeper "github.com/cosmos/interchain-queries/testutil/keeper"
	ibcquerykeeper "github.com/cosmos/interchain-queries/x/ibc_query/keeper"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// TestSumitCrossChainQuery validates SubmitCrossChainQuery implementation against the spec.
// Submit query message to quering chain and store the query in the local, private store.
func TestSubmitCrossChainQuery(t *testing.T) {
	senderAddr := "cosmos37dtl0mjt3t77dpoyz2edqzjpszulwhgnga7ljs"

	var (
		IBCQueryKeeper     ibcquerykeeper.Keeper
		ctx                sdk.Context
		mocks              testkeeper.MockedKeepers
	)

	testCases := []struct {
		name              string
		expPass           bool
		setUp             func()
		query             types.MsgSubmitCrossChainQuery
		expQuery          types.CrossChainQuery
		mockExpectations  func(sdk.Context, testkeeper.MockedKeepers) []*gomock.Call
	}{
		{
			"success",true, func() {
				IBCQueryKeeper, ctx, _,mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
			},
			types.MsgSubmitCrossChainQuery{
				Path:               "test/query_path",
				LocalTimeoutHeight: clienttypes.NewHeight(0, uint64(ctx.BlockHeight() + 50)),
				LocalTimeoutStamp:  uint64(ctx.BlockTime().UnixNano() + 5000000),
				QueryHeight:        uint64(ctx.BlockHeight()),
				ChainId:            "ibc-query",
				Sender:             senderAddr,
			},
			types.CrossChainQuery{
				Id:                     "query-0",
				Path:                   "test/query_path",
				LocalTimeoutHeight:     clienttypes.NewHeight(0, uint64(ctx.BlockHeight() + 50)),
				LocalTimeoutTimestamp:  uint64(ctx.BlockTime().UnixNano() + 5000000),
				QueryHeight:            uint64(ctx.BlockHeight()),
				ChainId:                "ibc-query",
			},
			func (sdk.Context, testkeeper.MockedKeepers) []*gomock.Call {	
				return testkeeper.GetMocksForSubmitCrossChainQuery(ctx, &mocks, testkeeper.GetQueryId(IBCQueryKeeper, ctx), senderAddr)
			},
		},
		{
			"fail: invalid local timestamp", false, func() {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
			},
			types.MsgSubmitCrossChainQuery{
				Path:               "test/query_path",
				LocalTimeoutHeight: clienttypes.NewHeight(0, uint64(ctx.BlockHeight() + 50)),
				LocalTimeoutStamp:  uint64(ctx.BlockTime().UnixNano()),
				QueryHeight:        uint64(ctx.BlockHeight()),
				ChainId:            "ibc-query",
				Sender:             senderAddr,
			},
			types.CrossChainQuery{},
			func (sdk.Context, testkeeper.MockedKeepers) []*gomock.Call {	
				return []*gomock.Call{}
			},
		},
	}

	for _, tc := range testCases {
		tc.setUp()

		// Run SubmitCrossChainQuery
		gomock.InOrder(tc.mockExpectations(ctx, mocks)...)
		res, err := IBCQueryKeeper.SubmitCrossChainQuery(ctx, &tc.query)
		
		if tc.expPass {
			queryResult, found := IBCQueryKeeper.GetCrossChainQuery(ctx, types.QueryPath(res.Id))
			require.True(t, found)
			require.NoError(t, err)
			require.NotNil(t, res)
			require.Equal(t, res.Id, queryResult.Id)
			require.Equal(t, tc.expQuery.Path, queryResult.Path)
			require.Equal(t, tc.expQuery.LocalTimeoutHeight.RevisionHeight, queryResult.LocalTimeoutHeight.RevisionHeight)
			require.Equal(t, tc.expQuery.LocalTimeoutTimestamp, queryResult.LocalTimeoutTimestamp)
			require.Equal(t, tc.expQuery.QueryHeight, queryResult.QueryHeight)
			require.Equal(t, tc.expQuery.ChainId, queryResult.ChainId)
		} else {
			require.Error(t, err)
		}
	}
}


// TestSubmitCrossChainQeuryResult validates SubmitCrossChainQueryResult implementation against the spec.
// If relayer gets query data from queried chain, relayer calls SubmitCrossChainQueryResult function.
// Delete the query from the local, private store and create a query result record.
func TestSubmitCrossChainQueryResult(t *testing.T) {
	_, _, senderAddr1 := testdata.KeyTestPubAddr()
	_, _, senderAddr2 := testdata.KeyTestPubAddr()

    var (
		msgSubmitQuery  types.MsgSubmitCrossChainQuery
		IBCQueryKeeper  ibcquerykeeper.Keeper
		ctx             sdk.Context
		mocks           testkeeper.MockedKeepers
    )

    testCases := []struct {
        name           string
		expPass        bool
        setUp          func()
		queryResult    types.MsgSubmitCrossChainQueryResult
		expQueryResult types.CrossChainQueryResult
    }{
        {
            "success: valid query id", true, func (){
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr1.String())
			},
			types.MsgSubmitCrossChainQueryResult{
				Id:           "query-0",
				QueryHeight:  uint64(ctx.BlockHeight()),
				Result:       types.QueryResult_QUERY_RESULT_SUCCESS,
				Data:         []byte("test result data"),
				Sender:       senderAddr1.String(),
			},
			types.CrossChainQueryResult{
				Id:      "query-0",
				Result:  types.QueryResult_QUERY_RESULT_SUCCESS,
				Data:    []byte("test result data"),
				Sender:  senderAddr1.String(),
			},
        },
		{
            "fail: invalid query id", false, func () {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr1.String())
			},
			types.MsgSubmitCrossChainQueryResult{
				Id:           "query-1",
				QueryHeight:  uint64(ctx.BlockHeight()),
				Result:       types.QueryResult_QUERY_RESULT_SUCCESS,
				Data:         []byte("test result data"),
				Sender:       senderAddr1.String(),
			},
			types.CrossChainQueryResult{},
        },
		{
            "success: send query sender1 and sender2", true, func () {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr1.String())
				gomock.InOrder(testkeeper.GetMocksForSubmitCrossChainQuery(ctx, &mocks, testkeeper.GetQueryId(IBCQueryKeeper, ctx), msgSubmitQuery.Sender)...)
				IBCQueryKeeper.SubmitCrossChainQuery(ctx, &msgSubmitQuery)
				
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr2.String())
			},
			types.MsgSubmitCrossChainQueryResult{
				Id:           "query-1",
				QueryHeight:  uint64(ctx.BlockHeight()),
				Result:       types.QueryResult_QUERY_RESULT_SUCCESS,
				Data:         []byte("test result data"),
				Sender:       senderAddr2.String(),
			},
			types.CrossChainQueryResult{
				Id:      "query-1",
				Result:  types.QueryResult_QUERY_RESULT_SUCCESS,
				Data:    []byte("test result data"),
				Sender:  senderAddr2.String(),
			},
        },
	}

	for _, tc := range testCases {
		tc.setUp()
		
		// Run SubmitCrossChainQuery
		gomock.InOrder(testkeeper.GetMocksForSubmitCrossChainQuery(ctx, &mocks, testkeeper.GetQueryId(IBCQueryKeeper, ctx), msgSubmitQuery.Sender)...)
		query, _ := IBCQueryKeeper.SubmitCrossChainQuery(ctx, &msgSubmitQuery)   

		// Run SubmitCrossChainQueryResult
		_, err := IBCQueryKeeper.SubmitCrossChainQueryResult(ctx, &tc.queryResult)
		if tc.expPass {
			require.NoError(t, err)
			// After a relayer sumits query result, query should be deleted from store
			_, found := IBCQueryKeeper.GetCrossChainQuery(ctx, types.QueryPath(query.Id))
			require.False(t, found)

			queryResult, found := IBCQueryKeeper.GetCrossChainQueryResult(ctx, types.ResultQueryPath(query.Id))
			require.True(t, found)
			require.Equal(t, tc.expQueryResult.Id, queryResult.Id)
			require.Equal(t, tc.expQueryResult.Data, tc.queryResult.Data)
			require.Equal(t, tc.expQueryResult.Result, tc.queryResult.Result)
			require.Equal(t, tc.expQueryResult.Sender, queryResult.Sender)
		} else {
			require.Error(t, err)
		}
    }
}


// TestSubmitPruneCrossChainQueryResult validates PruneCrossChainQueryResult against the spec.
// If the caller has the right to clean the query result, the function delete the query result from the the local, private store.
func TestSubmitPruneCrossChainQueryResult(t *testing.T) {
	_, _, invalidSenderAddr := testdata.KeyTestPubAddr()
	_, _, senderAddr        := testdata.KeyTestPubAddr()
	dummyCap                := &capabilitytypes.Capability{}

    var (
		msgSubmitQuery  types.MsgSubmitCrossChainQuery
		queryResult     types.MsgSubmitCrossChainQueryResult
		IBCQueryKeeper  ibcquerykeeper.Keeper
		ctx             sdk.Context
		mocks           testkeeper.MockedKeepers
    )

	testCases := []struct {
        name               string
		expPass            bool
        setUp              func()
		pruneMsg           types.MsgSubmitPruneCrossChainQueryResult
		mockExpectations   func(sdk.Context, testkeeper.MockedKeepers) []*gomock.Call
    }{
        {
            "success: valid query id", true,
			func () {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr.String())
				queryResult = testkeeper.GetMsgSubmitCrossChainQueryResult(ctx, "query-0", senderAddr.String())
			},
			types.MsgSubmitPruneCrossChainQueryResult {
				Id: "query-0",       
				Sender: senderAddr.String(),
			},
			func(sdk.Context, testkeeper.MockedKeepers) []*gomock.Call {
				return testkeeper.GetMocksForPruneCrossChainQueryResult(ctx, &mocks, "query-0", senderAddr.String())
			},
        },
		{
            "fail: invalid query id", false,
			func () {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr.String())
				queryResult = testkeeper.GetMsgSubmitCrossChainQueryResult(ctx, "query-0", senderAddr.String())
			},
			types.MsgSubmitPruneCrossChainQueryResult {
				Id: "query-1",
				Sender: senderAddr.String(),
			},
			func(sdk.Context, testkeeper.MockedKeepers) []*gomock.Call{
				return []*gomock.Call{}
			},
        },
		{
            "success: valid query id", true,
			func () {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr.String())
				queryResult = testkeeper.GetMsgSubmitCrossChainQueryResult(ctx, "query-0", senderAddr.String())
				// submit query-0 query 
				gomock.InOrder(testkeeper.GetMocksForSubmitCrossChainQuery(ctx, &mocks, "query-0", msgSubmitQuery.Sender)...)
				IBCQueryKeeper.SubmitCrossChainQuery(ctx, &msgSubmitQuery)   
				IBCQueryKeeper.SubmitCrossChainQueryResult(ctx, &queryResult)

				queryResult = testkeeper.GetMsgSubmitCrossChainQueryResult(ctx, "query-1", senderAddr.String())
			},
			types.MsgSubmitPruneCrossChainQueryResult {
				Id: "query-1",       
				Sender: senderAddr.String(),
			},
			func(sdk.Context, testkeeper.MockedKeepers) []*gomock.Call {
				return testkeeper.GetMocksForPruneCrossChainQueryResult(ctx, &mocks, "query-1", senderAddr.String())
			},
        },
		{
            "fail: invalid capability", false,
			func () {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				msgSubmitQuery = testkeeper.GetMsgSubmitCrossChainQuery(ctx, "test/path", senderAddr.String())
				queryResult = testkeeper.GetMsgSubmitCrossChainQueryResult(ctx, "query-0", senderAddr.String())
			},
			types.MsgSubmitPruneCrossChainQueryResult {
				Id: "query-0",
				Sender: invalidSenderAddr.String(),
			},
			func(ctx sdk.Context, mocks testkeeper.MockedKeepers) []*gomock.Call{			
				return []*gomock.Call{mocks.MockScopedKeeper.EXPECT().GetCapability(
					ctx, types.FormatQueryCapabilityIdentifier("query-0", invalidSenderAddr.String()),
					).Return(dummyCap, false).Times(1),
				}
			},
        },
	}

	for _, tc := range testCases {
		tc.setUp()

		// Run SubmitCrossChainQuery
		gomock.InOrder(testkeeper.GetMocksForSubmitCrossChainQuery(ctx, &mocks, testkeeper.GetQueryId(IBCQueryKeeper, ctx), msgSubmitQuery.Sender)...)
		IBCQueryKeeper.SubmitCrossChainQuery(ctx, &msgSubmitQuery)
		IBCQueryKeeper.SubmitCrossChainQueryResult(ctx, &queryResult)   
 	
		// Run SubmitPruneCrossChainQueryResult
		gomock.InOrder(tc.mockExpectations(ctx, mocks)...)
		pruneResult, err := IBCQueryKeeper.SubmitPruneCrossChainQueryResult(ctx, &tc.pruneMsg)	
		if tc.expPass {
			require.NoError(t, err)
			require.Equal(t, queryResult.Id, pruneResult.Id)
			require.Equal(t, queryResult.Data, pruneResult.Data)
			require.Equal(t, queryResult.Result, pruneResult.Result)

			// Unable to read query results
			_, err := IBCQueryKeeper.CrossChainQueryResult(ctx, &types.QueryCrossChainQueryResult{Id: pruneResult.Id})
			require.Error(t, err)
		} else {
			require.Error(t, err)
		}
	}

}