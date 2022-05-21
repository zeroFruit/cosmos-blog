package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeDeleteCommentRequest = "delete_comment"

var _ sdk.Msg = &DeleteCommentRequest{}

func NewDeleteCommentRequest(creator string, commentID uint64, postID uint64) *DeleteCommentRequest {
	return &DeleteCommentRequest{
		Creator:   creator,
		CommentId: commentID,
		PostId:    postID,
	}
}

func (msg *DeleteCommentRequest) Route() string {
	return RouterKey
}

func (msg *DeleteCommentRequest) Type() string {
	return TypeDeleteCommentRequest
}

func (msg *DeleteCommentRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *DeleteCommentRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *DeleteCommentRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
