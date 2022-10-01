package interchainqueries_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "interchain-queries/testutil/keeper"
	"interchain-queries/testutil/nullify"
	"interchain-queries/x/interchainqueries"
	"interchain-queries/x/interchainqueries/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.InterchainqueriesKeeper(t)
	interchainqueries.InitGenesis(ctx, *k, genesisState)
	got := interchainqueries.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
