package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "interchain-queries/testutil/keeper"
	"interchain-queries/x/interchainqueries/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.InterchainqueriesKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
