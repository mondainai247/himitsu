package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendCidHash = "send_cid_hash"

var _ sdk.Msg = &MsgSendCidHash{}

func NewMsgSendCidHash(creator string, tomodachi string, cidHash string) *MsgSendCidHash {
	return &MsgSendCidHash{
		Creator:   creator,
		Tomodachi: tomodachi,
		CidHash:   cidHash,
	}
}

func (msg *MsgSendCidHash) Route() string {
	return RouterKey
}

func (msg *MsgSendCidHash) Type() string {
	return TypeMsgSendCidHash
}

func (msg *MsgSendCidHash) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendCidHash) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendCidHash) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
