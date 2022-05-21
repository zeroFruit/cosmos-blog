package keeper_test

import (
	"testing"

	testkeeper "github.com/cosmonaut/blog/testutil/keeper"
	"github.com/cosmonaut/blog/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.BlogKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.ParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.ParamsResponse{Params: params}, response)
}
