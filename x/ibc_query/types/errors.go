package types

import (
	errorsmod "cosmossdk.io/errors"
)

// ibc_query sentinel errors
var (
	ErrInvalidVersion          = errorsmod.Register(ModuleName, 2, "invalid 31-IBC-query version")
	ErrInvalidQuery            = errorsmod.Register(ModuleName, 3, "invalid cross chain query")
	ErrTimeout                 = errorsmod.Register(ModuleName, 4, "cross chain query timeout")
	ErrCrossChainQueryNotFound = errorsmod.Register(ModuleName, 5, "no query found for given query id")
	ErrMaxTransferChannels     = errorsmod.Register(ModuleName, 6, "max transfer channels")
	ErrUnknownDataType         = errorsmod.Register(ModuleName, 7, "unknown data type")
	ErrInvalidCapability       = errorsmod.Register(ModuleName, 8, "invalid capability")
)
