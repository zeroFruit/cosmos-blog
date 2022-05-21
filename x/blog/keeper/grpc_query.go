package keeper

import (
	"github.com/cosmonaut/blog/x/blog/types"
)

var _ types.QueryServiceServer = Keeper{}
