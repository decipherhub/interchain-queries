package types

import (
	"encoding/hex"

	ics23 "github.com/confio/ics23/go"
	sdk "github.com/cosmos/cosmos-sdk/types"
	captypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v5/modules/core/02-client/types"
)

const (
	TypeMsgSubmitCrossChainQuery            = "submitCrossChainQuery"
	TypeMsgSubmitCrossChainQueryResult      = "submitCrossChainQueryResult"
	TypeMsgSubmitPruneCrossChainQueryResult = "submitPruneCrossChainQueryResult"
)

// NewMsgSubmitCrossChainQuery creates a new instance of NewMsgSubmitCrossChainQuery
func NewMsgSubmitCrossChainQuery(path string, localTimeoutHeight clienttypes.Height, localTimeoutStamp uint64, queryHeight uint64, chainId string, sender string) *MsgSubmitCrossChainQuery {
	return &MsgSubmitCrossChainQuery{
		Path:               path,
		LocalTimeoutHeight: localTimeoutHeight,
		LocalTimeoutStamp:  localTimeoutStamp,
		QueryHeight:        queryHeight,
		ChainId:            chainId,
		Sender:             sender,
	}
}

func (msg MsgSubmitCrossChainQuery) GetQueryPath() []byte {
	src := []byte(msg.Path)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return dst
}

func (msg MsgSubmitCrossChainQuery) GetTimeoutHeight() clienttypes.Height {
	return msg.LocalTimeoutHeight
}

func (msg MsgSubmitCrossChainQuery) GetTimeoutTimestamp() uint64 { return msg.LocalTimeoutStamp }

func (msg MsgSubmitCrossChainQuery) GetQueryHeight() uint64 { return msg.QueryHeight }

// ValidateBasic implements sdk.Msg and performs basic stateless validation
func (msg MsgSubmitCrossChainQuery) ValidateBasic() error {

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgSubmitCrossChainQuery) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

// Route implements sdk.Msg
func (msg MsgSubmitCrossChainQuery) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgSubmitCrossChainQuery) Type() string {
	return TypeMsgSubmitCrossChainQuery
}

// GetSignBytes implements sdk.Msg.
func (msg MsgSubmitCrossChainQuery) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// NewMsgSubmitCrossChainQueryResult creates a new instance of NewMsgSubmitCrossChainQueryResult
func NewMsgSubmitCrossChainQueryResult(id string, queryHeight uint64, result QueryResult, data []byte, proofs []*ics23.ProofSpec) *MsgSubmitCrossChainQueryResult {
	return &MsgSubmitCrossChainQueryResult{
		Id:          id,
		QueryHeight: queryHeight,
		Result:      result,
		Data:        data,
		ProofSpecs:  proofs,
	}
}

func (msg MsgSubmitCrossChainQueryResult) GetId() string { return msg.Id }

func (msg MsgSubmitCrossChainQueryResult) GetQueryHeight() uint64 {
	return msg.QueryHeight
}

func (msg MsgSubmitCrossChainQueryResult) GetResult() QueryResult { return msg.Result }

func (msg MsgSubmitCrossChainQueryResult) GetData() []byte { return msg.Data }

// ValidateBasic implements sdk.Msg and performs basic stateless validation
func (msg MsgSubmitCrossChainQueryResult) ValidateBasic() error {
	// TODO: Validate with proof
	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgSubmitCrossChainQueryResult) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

// Route implements sdk.Msg
func (msg MsgSubmitCrossChainQueryResult) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgSubmitCrossChainQueryResult) Type() string {
	return TypeMsgSubmitCrossChainQueryResult
}

// GetSignBytes implements sdk.Msg.
func (msg MsgSubmitCrossChainQueryResult) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}

// NewMsgSubmitPruneCrossChainQueryResult creates a new instance of MsgSubmitPruneCrossChainQueryResult
func NewMsgSubmitPruneCrossChainQueryResult(id string, capKey *captypes.Capability, sender string) *MsgSubmitPruneCrossChainQueryResult {
	return &MsgSubmitPruneCrossChainQueryResult{
		Id:     id,
		CapKey: capKey,
		Sender: sender,
	}
}

// ValidateBasic implements sdk.Msg and performs basic stateless validation
func (msg MsgSubmitPruneCrossChainQueryResult) ValidateBasic() error {
	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgSubmitPruneCrossChainQueryResult) GetSigners() []sdk.AccAddress {
	signer, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signer}
}

// Route implements sdk.Msg
func (msg MsgSubmitPruneCrossChainQueryResult) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg MsgSubmitPruneCrossChainQueryResult) Type() string {
	return TypeMsgSubmitPruneCrossChainQueryResult
}

// GetSignBytes implements sdk.Msg.
func (msg MsgSubmitPruneCrossChainQueryResult) GetSignBytes() []byte {
	return sdk.MustSortJSON(AminoCdc.MustMarshalJSON(&msg))
}
