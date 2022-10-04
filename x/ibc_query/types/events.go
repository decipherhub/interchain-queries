package types

import (
	fmt "fmt"

	host "github.com/cosmos/ibc-go/v5/modules/core/24-host"
)

const (
	EventTypeTimeout = "timeout"
	EventSendQuery   = "send_query"
	EventTypePacket  = "ibc_query_packet"

	AttributeKeyQueryID     = "query_id"
	AttributeKeyQueryHeight = "query_height"
	AttributeKeyQueryPath   = "query_path"
	AttributeKeyChainId     = "chain_id"
	AttributeKeyAckSuccess  = "success"
	AttributeKeyAckError    = "error"
)

var (
	AttributeValueCategory = fmt.Sprintf("%s_%s", host.ModuleName, ModuleName)
)
