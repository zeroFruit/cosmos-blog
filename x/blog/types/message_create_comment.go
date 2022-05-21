package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateComment = "create_comment"

var _ sdk.Msg = &CreateCommentRequest{}

func NewMsgCreateComment(creator string, postID uint64, title string, body string) *CreateCommentRequest {
	return &CreateCommentRequest{
		Creator: creator,
		PostId:  postID,
		Title:   title,
		Body:    body,
	}
}

func (msg *CreateCommentRequest) Route() string {
	return RouterKey
}

func (msg *CreateCommentRequest) Type() string {
	return TypeMsgCreateComment
}

func (msg *CreateCommentRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *CreateCommentRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *CreateCommentRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
