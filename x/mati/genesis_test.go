package mati_test

import (
	"testing"

	keepertest "github.com/jemrickrioux/mati/testutil/keeper"
	"github.com/jemrickrioux/mati/testutil/nullify"
	"github.com/jemrickrioux/mati/x/mati"
	"github.com/jemrickrioux/mati/x/mati/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MatiKeeper(t)
	mati.InitGenesis(ctx, *k, genesisState)
	got := mati.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
