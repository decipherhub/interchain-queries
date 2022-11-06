package e2e

import (
	"errors"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
	"github.com/cosmos/interchain-queries/app"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
	"strings"
	"time"
)

func (i *IBCQueryTestSuite) sendIBCQuery(query testIBCQuery) (queryId string, capKey string, err error) {
	msg := types.NewMsgSubmitCrossChainQuery(query.path, clienttypes.NewHeight(0, 0), uint64(time.Now().Add(time.Hour).UnixNano()), 0, i.queriedChain.ChainID, query.sender)
	resp, err := i.queryingChain.App.(*app.App).IBCQueryKeeper.SubmitCrossChainQuery(i.queryingChain.GetContext(), msg)
	if err != nil {
		return "", "", err
	}
	return resp.Id, resp.CapKey, nil
}

func (i *IBCQueryTestSuite) submitIBCQueryResult(result testIBCQueryResult) error {
	msg := types.NewMsgSubmitCrossChainQueryResult(result.id, 0, result.result, result.data, result.sender, nil)
	_, err := i.queryingChain.App.(*app.App).IBCQueryKeeper.SubmitCrossChainQueryResult(i.queryingChain.GetContext(), msg)
	if err != nil {
		return err
	}
	return nil
}

func (i *IBCQueryTestSuite) pruneIBCQueryResult(queryId string, sender string) (result types.QueryResult, data []byte, err error) {
	msg := types.NewMsgSubmitPruneCrossChainQueryResult(queryId, sender)
	resp, err := i.queryingChain.App.(*app.App).IBCQueryKeeper.SubmitPruneCrossChainQueryResult(i.queryingChain.GetContext(), msg)
	if err != nil {
		return 0, []byte{}, err
	}

	return resp.Result, resp.Data, nil
}

func (i *IBCQueryTestSuite) getIBCQuery(queryId string) (path string, isExist bool) {
	query, isExist := i.queryingChain.App.(*app.App).IBCQueryKeeper.GetCrossChainQuery(i.queryingChain.GetContext(), types.QueryPath(queryId))
	return query.Path, isExist
}

func (i *IBCQueryTestSuite) relayToQueriedChainBankModule(queryId string, path string, denom string, sender string) (testIBCQueryResult, error) {
	moduleName, target, addr, err := pathParser(path)
	if err != nil {
		return testIBCQueryResult{}, err
	}
	if moduleName != "bank" {
		return testIBCQueryResult{}, errors.New("module name in path is not bank module")
	}
	if target != "balances" {
		return testIBCQueryResult{}, errors.New("query target is not balance")
	}

	balance, err := i.grpcQueryToBankModule(addr, denom)
	if err != nil {
		return testIBCQueryResult{}, err
	}
	data, err := balance.Marshal()
	if err != nil {
		return testIBCQueryResult{}, err
	}
	return testIBCQueryResult{id: queryId, result: types.QueryResult_QUERY_RESULT_SUCCESS, data: data, sender: sender}, nil
}

func pathParser(path string) (moduleName string, target string, addr string, err error) {
	// path: http://{nodeAddr}/cosmos/{moduleName}/{version}/{target}/{targetAddr}
	splitArr := strings.Split(path, "/")
	if len(splitArr) != 8 {
		return "", "", "", errors.New("invalid path")
	}

	return splitArr[4], splitArr[6], splitArr[7], nil
}

func (i *IBCQueryTestSuite) grpcQueryToBankModule(addr string, denom string) (balance *cosmostypes.Coin, err error) {
	request := &banktypes.QueryBalanceRequest{Address: addr, Denom: denom}
	resp, err := i.queriedChain.App.(*app.App).BankKeeper.Balance(i.queriedChain.GetContext(), request)
	if err != nil {
		return nil, err
	}
	return resp.Balance, err
}

func (i *IBCQueryTestSuite) setAddrBalanceInQueriedChain(addr cosmostypes.AccAddress, balance cosmostypes.Coin) error {
	ctx := i.queriedChain.GetContext()
	convAddr := i.queriedChain.App.(*app.App).AccountKeeper.NewAccountWithAddress(ctx, addr)
	i.queriedChain.App.(*app.App).AccountKeeper.SetAccount(ctx, convAddr)
	bankKeeper := i.queriedChain.App.(*app.App).BankKeeper
	amounts := []cosmostypes.Coin{balance}
	if err := bankKeeper.MintCoins(ctx, minttypes.ModuleName, amounts); err != nil {
		return err
	}

	return bankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, addr, amounts)
}

func trimQueriedData(str string) string {
	str = strings.Trim(str, "\n\t")
	str = strings.Replace(str, "\x12\x04", " ", -1)
	return str
}
