package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/interchain-queries/x/ibc-query/keeper"
)

func GetQueryIdentifier() (string, error) {
	return keeper.Keeper.GenerateQueryIdentifier(keeper.Keeper{}, sdk.Context{})
}
