package e2e

import (
	ibctesting "github.com/cosmos/ibc-go/v5/testing"
	"github.com/cosmos/interchain-queries/testutil/simapp"
	"github.com/stretchr/testify/suite"
)

// IBCQueryTestSuite QueriedChain doesn't receive packet in IBCQuery
// So we don't need to set up IBC connection
type IBCQueryTestSuite struct {
	suite.Suite
	coordinator   *ibctesting.Coordinator
	queryingChain *ibctesting.TestChain
	queriedChain  *ibctesting.TestChain
}

func (suite *IBCQueryTestSuite) SetupTest() {
	suite.coordinator, suite.queryingChain, suite.queriedChain = simapp.NewCrossChainQueryCoordinator(suite.T())
}
