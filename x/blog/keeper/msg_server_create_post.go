package keeper

import (
	"context"

	"github.com/cosmonaut/blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.CreatePostRequest) (*types.CreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	post := types.Post{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Title:     msg.Title,
		Body:      msg.Body,
		CreatedAt: ctx.BlockHeight(),
	}

	id := k.AppendPost(ctx, post)

	return &types.CreatePostResponse{
		Id: id,
	}, nil
}
