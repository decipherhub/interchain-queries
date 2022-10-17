package e2e

import (
	captypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	testing "github.com/cosmos/ibc-go/v5/testing"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/interchain-queries/app"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
	"time"
)

type testIBCQuery struct {
	path        string
	queryHeight uint64
	sender      string
}

type testIBCQueryResult struct {
	id          string
	queryHeight uint64
	result      types.QueryResult
	data        []byte
}

type testCase struct {
	path string
	queryHeight uint64
	expData []byte
}

func sendIBCQuery(i *IBCQueryTestSuite, query testIBCQuery) (queryId string, capKey *captypes.Capability, err error) {
	msg := types.NewMsgSubmitCrossChainQuery(query.path, clienttypes.NewHeight(0, 0), uint64(time.Now().Add(time.Hour).UnixNano()), query.queryHeight, i.queriedChain.ChainID, query.sender)
	resp, err := i.queryingChain.App.(*app.App).IBCQueryKeeper.SubmitCrossChainQuery(i.queryingChain.GetContext(), msg)
	if err != nil {
		return "", captypes.NewCapability(0), err
	}
	return resp.Id, resp.CapKey, nil
}

func submitIBCQueryResult(i *IBCQueryTestSuite, result testIBCQueryResult) error {
	msg := types.NewMsgSubmitCrossChainQueryResult(result.id, result.queryHeight, result.result, result.data, nil)
	_, err := i.queryingChain.App.(*app.App).IBCQueryKeeper.SubmitCrossChainQueryResult(i.queryingChain.GetContext(), msg)
	if err != nil {
		return err
	}
	return nil
}

func pruneIBCQueryResult(i *IBCQueryTestSuite, queryId string, capKey *captypes.Capability, sender string) (result types.QueryResult, data []byte, err error) {
	msg := types.NewMsgSubmitPruneCrossChainQueryResult(queryId, capKey, sender)
	resp, err := i.queryingChain.App.(*app.App).IBCQueryKeeper.SubmitPruneCrossChainQueryResult(i.queryingChain.GetContext(), msg)
	if err != nil {
		return 0, []byte{}, err
	}

	return resp.Result, resp.Data, nil
}

func grpcQueryToBankModule(i *IBCQueryTestSuite, addr string) (data []byte) {
	request := &banktypes.QueryBalanceRequest{Address: addr}
	i.queriedChain.App.(*app.App).BankKeeper.Balance(i.queryingChain.GetContext(), )
}

func (i *IBCQueryTestSuite) TestIBCQuery() {
	i.SetupTest()
	query := testIBCQuery{path: , queryHeight: , }
	id, capKey, err := sendIBCQuery(i, )
}

func checkInterface(app testing.TestingApp) {

}
