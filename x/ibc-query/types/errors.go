package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ibc-query sentinel errors
var (
	ErrInvalidVersion          = sdkerrors.Register(ModuleName, 2, "invalid 31-IBC-query version")
	ErrInvalidQuery            = sdkerrors.Register(ModuleName, 3, "invalid cross chain query")
	ErrTimeout                 = sdkerrors.Register(ModuleName, 4, "cross chain query timeout")
	ErrCrossChainQueryNotFound = sdkerrors.Register(ModuleName, 5, "no query found for given query id")
	ErrMaxTransferChannels     = sdkerrors.Register(ModuleName, 6, "max transfer channels")
	ErrUnknownDataType         = sdkerrors.Register(ModuleName, 7, "unknown data type")
)
