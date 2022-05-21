package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeCreatePostRequest = "create_post"

var _ sdk.Msg = &CreatePostRequest{}

func NewCreatePostRequest(creator string, title string, body string) *CreatePostRequest {
	return &CreatePostRequest{
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *CreatePostRequest) Route() string {
	return RouterKey
}

func (msg *CreatePostRequest) Type() string {
	return TypeCreatePostRequest
}

func (msg *CreatePostRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *CreatePostRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *CreatePostRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
