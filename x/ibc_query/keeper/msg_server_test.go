package keeper_test

import (
	"fmt"
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
		IBCQueryKeeper ibcquerykeeper.Keeper
		ctx            sdk.Context
		mocks          testkeeper.MockedKeepers
		errFromNewCap  error
	)

	testCases := []struct {
		name     string
		expPass  bool
		setUp func()
		query      types.MsgSubmitCrossChainQuery
	}{
		{
			"success",true, func() {
				IBCQueryKeeper, ctx, _,mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				dummyCap := &capabilitytypes.Capability{}
				gomock.InOrder(			
					mocks.MockScopedKeeper.EXPECT().NewCapability(
						ctx, types.FormatQueryCapabilityIdentifier("query-0", senderAddr),
					).Return(dummyCap, errFromNewCap),
				)
			},
			types.MsgSubmitCrossChainQuery{
				Path:               "test/query_path",
				LocalTimeoutHeight: clienttypes.NewHeight(0, uint64(ctx.BlockHeight() + 50)),
				LocalTimeoutStamp:  uint64(ctx.BlockTime().UnixNano() + 5000000),
				QueryHeight:        uint64(ctx.BlockHeight()),
				ChainId:            "ibc-query",
				Sender:             senderAddr,
			},
		},
		{
			"fail", false, func() {
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
		},
	}

	for _, tc := range testCases {
		tc.setUp()

		// Run SubmitCrossChainQuery
		res, err := IBCQueryKeeper.SubmitCrossChainQuery(ctx, &tc.query)
		
		if tc.expPass {
			queryResult, found := IBCQueryKeeper.GetCrossChainQuery(ctx, types.QueryPath(res.Id))
			require.True(t, found)
			require.NoError(t, err)
			require.NotNil(t, res)
			require.Equal(t, res.Id, queryResult.Id)
			require.Equal(t, tc.query.Path, queryResult.Path)
			require.Equal(t, tc.query.LocalTimeoutHeight.RevisionHeight, queryResult.LocalTimeoutHeight.RevisionHeight)
			require.Equal(t, tc.query.LocalTimeoutStamp, queryResult.LocalTimeoutTimestamp)
			require.Equal(t, tc.query.QueryHeight, queryResult.QueryHeight)
			require.Equal(t, tc.query.ChainId, queryResult.ChainId)
		} else {
			require.Error(t, err)
		}
	}
}


// TestSubmitCrossChainQeuryResult validates SubmitCrossChainQueryResult implementation against the spec.
// If relayer gets query data from queried chain, relayer calls SubmitCrossChainQueryResult function.
// Delete the query from the local, private store and create a query result record.
func TestSubmitCrossChainQueryResult(t *testing.T) {
	_, _, senderAddr := testdata.KeyTestPubAddr()

    var (
		query          *types.MsgSubmitCrossChainQueryResponse
		IBCQueryKeeper ibcquerykeeper.Keeper
		ctx            sdk.Context
		mocks          testkeeper.MockedKeepers
    )

    testCases := []struct {
        name          string
		expPass      bool
        setUp        func() (*types.MsgSubmitCrossChainQueryResponse)
		queryResult  types.MsgSubmitCrossChainQueryResult
    }{
        {
            "success: valid query id", true, func () (*types.MsgSubmitCrossChainQueryResponse) {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				query = testkeeper.SetupForSubmitCrossChainQueryState(IBCQueryKeeper, ctx, mocks, senderAddr.String())
				return query
			},
			types.MsgSubmitCrossChainQueryResult{
				Id:           "query-0",
				QueryHeight:  uint64(ctx.BlockHeight()),
				Result:       types.QueryResult_QUERY_RESULT_SUCCESS,
				Data:         []byte("test result data"),
				Sender:       senderAddr.String(),
			},
        },
		{
            "fail: invalid query id", false, func () (*types.MsgSubmitCrossChainQueryResponse) {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				query = testkeeper.SetupForSubmitCrossChainQueryState(IBCQueryKeeper, ctx, mocks, senderAddr.String())
				return query
			},
			types.MsgSubmitCrossChainQueryResult{
				Id:           "query-1",
				QueryHeight:  uint64(ctx.BlockHeight()),
				Result:       types.QueryResult_QUERY_RESULT_SUCCESS,
				Data:         []byte("test result data"),
				Sender:       senderAddr.String(),
			},
        },
	}

	for _, tc := range testCases {
		query := tc.setUp()
		_, found := IBCQueryKeeper.GetCrossChainQuery(ctx, types.QueryPath(query.Id))
		require.True(t, found)      

		// Run SubmitCrossChainQueryResult
		_, err := IBCQueryKeeper.SubmitCrossChainQueryResult(ctx, &tc.queryResult)
		if tc.expPass {
			require.NoError(t, err)
			// After a relayer sumits query result, query should be deleted from store
			_, found := IBCQueryKeeper.GetCrossChainQuery(ctx, types.QueryPath(query.Id))
			require.False(t, found)

			queryResult, found := IBCQueryKeeper.GetCrossChainQueryResult(ctx, types.ResultQueryPath(query.Id))
			require.True(t, found)
			require.Equal(t, queryResult.Id, tc.queryResult.Id)
			require.Equal(t, queryResult.Data, tc.queryResult.Data)
			require.Equal(t, queryResult.Result, tc.queryResult.Result)
			require.Equal(t, senderAddr.String(), queryResult.Sender)
		} else {
			require.Error(t, err)
		}
    }
}


// TestSubmitPruneCrossChainQueryResult validates PruneCrossChainQueryResult against the spec.
// If the caller has the right to clean the query result, the function delete the query result from the the local, private store.
func TestSubmitPruneCrossChainQueryResult(t *testing.T) {
	_, _, invalidSenderAddr := testdata.KeyTestPubAddr()
	_, _, senderAddr := testdata.KeyTestPubAddr()

    var (
		queryResult    types.MsgSubmitCrossChainQueryResult
		IBCQueryKeeper ibcquerykeeper.Keeper
		ctx            sdk.Context
		mocks          testkeeper.MockedKeepers
    )

	testCases := []struct {
        name          string
		expPass      bool
        setUp        func() (types.MsgSubmitCrossChainQueryResult)
		pruneMsg     types.MsgSubmitPruneCrossChainQueryResult
    }{
        {
            "success: valid query id", true,
			func () (types.MsgSubmitCrossChainQueryResult) {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				queryResult = testkeeper.SetupForSubmitCrossChainQueryResultState(IBCQueryKeeper, ctx, mocks, senderAddr.String())
				dummyCap := &capabilitytypes.Capability{}
				gomock.InOrder(		
					mocks.MockScopedKeeper.EXPECT().GetCapability(
						ctx, types.FormatQueryCapabilityIdentifier(queryResult.Id, queryResult.Sender),
					).Return(dummyCap, true).Times(1),
					mocks.MockScopedKeeper.EXPECT().AuthenticateCapability(
						ctx, dummyCap, types.FormatQueryCapabilityIdentifier(queryResult.Id, queryResult.Sender),
					).Return(true),
				)
				return queryResult
			},
			types.MsgSubmitPruneCrossChainQueryResult {
				Id: "query-0",       
				Sender: senderAddr.String(),
			},
        },
		{
            "fail: invalid query id", false,
			func () (types.MsgSubmitCrossChainQueryResult) {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				queryResult = testkeeper.SetupForSubmitCrossChainQueryResultState(IBCQueryKeeper, ctx, mocks, senderAddr.String())
				return queryResult
			},
			types.MsgSubmitPruneCrossChainQueryResult {
				Id: "query-1",
				Sender: senderAddr.String(),
			},
        },
		{
            "success: valid query id", true,
			func () (types.MsgSubmitCrossChainQueryResult) {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				testkeeper.SetupForSubmitCrossChainQueryResultState(IBCQueryKeeper, ctx, mocks, senderAddr.String())
				queryResult = testkeeper.SetupForSubmitCrossChainQueryResultState(IBCQueryKeeper, ctx, mocks, senderAddr.String())
				dummyCap := &capabilitytypes.Capability{}
				gomock.InOrder(		
					mocks.MockScopedKeeper.EXPECT().GetCapability(
						ctx, types.FormatQueryCapabilityIdentifier(queryResult.Id, queryResult.Sender),
					).Return(dummyCap, true).Times(1),
					mocks.MockScopedKeeper.EXPECT().AuthenticateCapability(
						ctx, dummyCap, types.FormatQueryCapabilityIdentifier(queryResult.Id, queryResult.Sender),
					).Return(true),
				)
				return queryResult
			},
			types.MsgSubmitPruneCrossChainQueryResult {
				Id: "query-1",       
				Sender: senderAddr.String(),
			},
        },
		{
            "fail: invalid capability", false,
			func () (types.MsgSubmitCrossChainQueryResult) {
				IBCQueryKeeper, ctx, _, mocks = testkeeper.GetIBCQueryKeeperAndCtx(t, testkeeper.NewInMemKeeperParams(t))
				queryResult = testkeeper.SetupForSubmitCrossChainQueryResultState(IBCQueryKeeper, ctx, mocks, senderAddr.String())
				dummyCap := &capabilitytypes.Capability{}
				gomock.InOrder(		
					mocks.MockScopedKeeper.EXPECT().GetCapability(
						ctx, types.FormatQueryCapabilityIdentifier(queryResult.Id, invalidSenderAddr.String()),
					).Return(dummyCap, true).Times(1),
					mocks.MockScopedKeeper.EXPECT().AuthenticateCapability(
						ctx, dummyCap, types.FormatQueryCapabilityIdentifier(queryResult.Id, senderAddr.String()),
					).Return(false),
				)
				return queryResult
			},
			types.MsgSubmitPruneCrossChainQueryResult {
				Id: "query-0",
				Sender: invalidSenderAddr.String(),
			},
        },
	}

	for _, tc := range testCases {
		queryResult := tc.setUp()
		_, found := IBCQueryKeeper.GetCrossChainQueryResult(ctx, types.ResultQueryPath(queryResult.Id))
		require.True(t, found)   
		
		// Run SubmitPruneCrossChainQueryResult
		fmt.Println(tc.name, queryResult.Id, senderAddr.String(), tc.pruneMsg.Sender)
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