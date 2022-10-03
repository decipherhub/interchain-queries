package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewIBCQueryPacketData contructs a new IBCQueryPacketData instance
func NewIBCQueryPacketData(
	id string, path string, queryHeight uint64,
) IBCQueryPacketData {
	return IBCQueryPacketData{
		Id:          id,
		Path:        path,
		QueryHeight: queryHeight,
	}
}

// ValidateBasic is used for validating the query packet data
func (iqpd IBCQueryPacketData) ValidateBasic() error {
	if strings.TrimSpace(iqpd.Path) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "query path cannot be blank")
	}
	return nil
}

// GetBytes is a helper for serialising
func (iqpd IBCQueryPacketData) GetBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&iqpd))
}
