package e2e

import (
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
)

type testIBCQuery struct {
	path   string
	sender string
}

type testIBCQueryResult struct {
	id     string
	result types.QueryResult
	data   []byte
	sender string
}

func TestIBCQueryTestSuite(t *testing.T) {
	suite.Run(t, new(IBCQueryTestSuite))
}

func (i *IBCQueryTestSuite) TestIBCQueryToBank() {
	_, _, senderAddr := testdata.KeyTestPubAddr()
	_, _, targetAddr := testdata.KeyTestPubAddr()
	const denom = "testDenom"
	const denomAmount = 1000
	balance := cosmostypes.NewInt64Coin(denom, denomAmount)
	err := i.setAddrBalanceInQueriedChain(targetAddr, balance)
	if err != nil {
		i.Require().Fail("fail to mint coin for test addr", err.Error())
		return
	}

	queryPath := "http://blockchains:27011/cosmos/bank/v1beta1/balances/" + targetAddr.String()
	query := testIBCQuery{path: queryPath, sender: senderAddr.String()}
	queryId, err := i.sendIBCQuery(query)
	if err != nil {
		i.Require().Fail("fail to send IBC query", err.Error())
		return
	}

	path, isExist := i.getIBCQuery(queryId)
	if !isExist {
		i.Require().Fail("IBC query doesn't exist in querying chain")
		return
	}

	result, err := i.relayToQueriedChainBankModule(queryId, path, denom, senderAddr.String())
	if err != nil {
		i.Require().Fail("fail to relay IBC query to queried chian", err.Error())
		return
	}

	err = i.submitIBCQueryResult(result)
	if err != nil {
		i.Require().Fail("fail to submit query result to querying chain", err.Error())
	}

	queryResult, data, err := i.pruneIBCQueryResult(queryId, senderAddr.String())
	if err != nil {
		i.Require().Fail("fail to prune query result", err.Error())
	}
	if queryResult != types.QueryResult_QUERY_RESULT_SUCCESS {
		i.Require().Fail("result is not result_success")
	}
	amountStr := strconv.Itoa(denomAmount)
	queriedData := string(data)
	queriedData = trimQueriedData(queriedData)
	i.Require().Equal(denom+" "+amountStr, queriedData)
}
