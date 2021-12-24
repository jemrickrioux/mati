package keeper_test

import (
	"testing"

	testkeeper "github.com/jemrickrioux/mati/testutil/keeper"
	"github.com/jemrickrioux/mati/x/mati/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MatiKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
