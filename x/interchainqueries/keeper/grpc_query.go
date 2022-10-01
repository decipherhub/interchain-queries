package keeper

import (
	"interchain-queries/x/interchainqueries/types"
)

var _ types.QueryServer = Keeper{}
