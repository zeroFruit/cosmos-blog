package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/cosmonaut/blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Comments(goCtx context.Context, req *types.CommentsRequest) (*types.CommentsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var comments []*types.Comment

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the key-value module store using the store key (in this case store key is "chain")
	store := ctx.KVStore(k.storeKey)

	// Get the part of the store that keeps posts (using post key, which is "Post-value-")
	commentStore := prefix.NewStore(store, []byte(types.CommentKey))

	// Get the post by ID
	post := k.GetPost(ctx, req.Id)

	// Get the post ID
	postID := post.Id

	// Paginate the posts store based on PageRequest
	pageRes, err := query.Paginate(commentStore, req.Pagination, func(key []byte, value []byte) error {
		var comment types.Comment
		if err := k.cdc.Unmarshal(value, &comment); err != nil {
			return err
		}

		if comment.PostId == postID {
			comments = append(comments, &comment)
		}

		return nil
	})

	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Return a struct containing a list of posts and pagination info
	return &types.CommentsResponse{Post: &post, Comment: comments, Pagination: pageRes}, nil
}
