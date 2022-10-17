package simapp

import (
	"encoding/json"
	"testing"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/interchain-queries/app"

	ibctesting "github.com/cosmos/ibc-go/v5/testing"

	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"
)

func SetupTestingApp() (ibctesting.TestingApp, map[string]json.RawMessage) {
	db := tmdb.NewMemDB()
	encoding := cosmoscmd.MakeEncodingConfig(app.ModuleBasics)
	testApp := app.New(log.NewNopLogger(), db, nil, true, map[int64]bool{}, simapp.DefaultNodeHome, 5, encoding, simapp.EmptyAppOptions{}).(ibctesting.TestingApp)
	return testApp, app.NewDefaultGenesisState(encoding.Marshaler)
}

// NewCoordinator initializes Coordinator with 2 TestChains
func NewCrossChainQueryCoordinator(t *testing.T) (*ibctesting.Coordinator, *ibctesting.TestChain, *ibctesting.TestChain) {
	ibctesting.DefaultTestingAppInit = SetupTestingApp
	coordinator := ibctesting.NewCoordinator(t, 2)
	queryingChain := coordinator.Chains[ibctesting.GetChainID(1)]
	queriedChain := coordinator.Chains[ibctesting.GetChainID(2)]
	return coordinator, queryingChain, queriedChain
}
