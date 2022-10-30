package keeper

import (
	"testing"

	"github.com/cosmos/interchain-queries/x/ibc_query/types"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	ibcquerykeeper "github.com/cosmos/interchain-queries/x/ibc_query/keeper"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

// Parameters needed to instantiate an in-memory keeper
type InMemKeeperParams struct {
	Cdc            *codec.ProtoCodec
	StoreKey       *storetypes.KVStoreKey
	ParamsSubspace *paramstypes.Subspace
	Ctx            sdk.Context
}

func NewInMemKeeperParams(t testing.TB) InMemKeeperParams {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	paramsSubspace := typesparams.NewSubspace(cdc,
		codec.NewLegacyAmino(),
		storeKey,
		memStoreKey,
		paramstypes.ModuleName,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	
	return InMemKeeperParams{
		Cdc:            cdc,
		StoreKey:       storeKey,
		ParamsSubspace: &paramsSubspace,
		Ctx:            ctx,
	}
}


// A struct holding pointers to any mocked external keeper needed for provider/consumer keeper setup.
type MockedKeepers struct {
	*MockScopedKeeper
	*MockChannelKeeper
	*MockPortKeeper
	*MockICS4Wrapper
}

// NewMockedKeepers instantiates a struct with pointers to properly instantiated mocked keepers.
func NewMockedKeepers(ctrl *gomock.Controller) MockedKeepers {
	return MockedKeepers{
		MockScopedKeeper:      NewMockScopedKeeper(ctrl),
		MockChannelKeeper:     NewMockChannelKeeper(ctrl),
		MockPortKeeper:        NewMockPortKeeper(ctrl),
		MockICS4Wrapper:       NewMockICS4Wrapper(ctrl),
	}
}


// NewInMemIBCQueryKeeper instantiates an in-mem provider keeper from params and mocked keepers
func NewInMemIBCQueryKeeper(params InMemKeeperParams, mocks MockedKeepers) ibcquerykeeper.Keeper {
	return ibcquerykeeper.NewKeeper(
		params.Cdc,
		params.StoreKey,
		mocks.MockScopedKeeper,
		mocks.MockICS4Wrapper,
		mocks.MockChannelKeeper,
		mocks.MockPortKeeper,
	)
}


// Returns an in-memory provider keeper, context, controller, and mocks, given a test instance and parameters.
//
// Note: Calling ctrl.Finish() at the end of a test function ensures that
// no unexpected calls to external keepers are made.
func GetIBCQueryKeeperAndCtx(t *testing.T, params InMemKeeperParams) (
	ibcquerykeeper.Keeper, sdk.Context, *gomock.Controller, MockedKeepers) {

	ctrl := gomock.NewController(t)
	mocks := NewMockedKeepers(ctrl)
	return NewInMemIBCQueryKeeper(params, mocks), params.Ctx, ctrl, mocks
}

// Returns an in-memory provider keeper, context, controller, and mocks after query data is stored in the store
func GetSubmitCrossChainQueryKeeperAndCtx(IBCQueryKeeper ibcquerykeeper.Keeper, 
	ctx sdk.Context, mocks MockedKeepers) (
	ibcquerykeeper.Keeper, sdk.Context, MockedKeepers, *types.MsgSubmitCrossChainQueryResponse) {

	var( errFromNewCap  error )
	senderAddr := "cosmos37dtl0mjt3t77dpoyz2edqzjpszulwhgnga7ljs"
	dummyCap := &capabilitytypes.Capability{}

	msg := types.MsgSubmitCrossChainQuery{
		Path:               "test/query_path",
		LocalTimeoutHeight: clienttypes.NewHeight(0, uint64(ctx.BlockHeight() + 50)),
		LocalTimeoutStamp:  uint64(ctx.BlockTime().UnixNano() + 5000000),
		QueryHeight:        uint64(ctx.BlockHeight()),
		ChainId:            "ibc-query",
		Sender:             senderAddr,
	}

	nextQuerySeq := IBCQueryKeeper.GetNextQuerySequence(ctx)
	queryID := types.FormatQueryIdentifier(nextQuerySeq)
	gomock.InOrder(			
		mocks.MockScopedKeeper.EXPECT().NewCapability(
			ctx, types.FormatQueryCapabilityIdentifier(queryID, msg.Sender),
		).Return(dummyCap, errFromNewCap),
	)
	query, _ := IBCQueryKeeper.SubmitCrossChainQuery(ctx, &msg)
	return IBCQueryKeeper, ctx, mocks, query
}


// Returns an in-memory provider keeper, context, controller, and mocks after query result data is stored in the store
func GetSubmitCrossChainQueryResultKeeperAndCtx(IBCQueryKeeper ibcquerykeeper.Keeper, 
	ctx sdk.Context, mocks MockedKeepers) (
	ibcquerykeeper.Keeper, sdk.Context, MockedKeepers, types.MsgSubmitCrossChainQueryResult) {

	senderAddr := "cosmos37dtl0mjt3t77dpoyz2edqzjpszulwhgnga7ljs"
	IBCQueryKeeper, ctx, mocks, query := GetSubmitCrossChainQueryKeeperAndCtx(IBCQueryKeeper, ctx, mocks)

	queryResult := types.MsgSubmitCrossChainQueryResult{
		Id:           query.Id,
		QueryHeight:  uint64(ctx.BlockHeight()),
		Result:       types.QueryResult_QUERY_RESULT_SUCCESS,
		Data:         []byte("test result data"),
		Sender:       senderAddr,
	}
	IBCQueryKeeper.SubmitCrossChainQueryResult(ctx, &queryResult)
	return IBCQueryKeeper, ctx, mocks, queryResult
}