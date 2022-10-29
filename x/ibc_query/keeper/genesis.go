package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-queries/x/ibc_query/types"
)

// InitGenesis initializes the application state from a provided genesis state
func (k Keeper) InitGenesis(ctx sdk.Context, state types.GenesisState) {
	k.SetPort(ctx, state.PortId)
	for _, query := range state.Queries {
		k.SetCrossChainQuery(ctx, types.QueryPath(query.Id), *query)
	}
	for _, result := range state.Results {
		k.SetCrossChainQueryResult(ctx, types.ResultQueryPath(result.Id), *result)
	}

	if !k.IsBound(ctx, state.PortId) {
		err := k.BindPort(ctx, state.PortId)
		if err != nil {
			panic(fmt.Sprintf("could not claim port capability: %v", err))
		}
	}
}

// ExportGenesis returns the application exported genesis
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Queries: k.GetAllCrossChainQueries(ctx),
		Results: k.GetAllCrossChainQueryResults(ctx),
	}
}
