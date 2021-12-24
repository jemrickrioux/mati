package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jemrickrioux/mati/testutil/keeper"
	"github.com/jemrickrioux/mati/x/mati/keeper"
	"github.com/jemrickrioux/mati/x/mati/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MatiKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
