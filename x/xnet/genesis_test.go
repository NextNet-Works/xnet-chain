package xnet_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "xnet/testutil/keeper"
	"xnet/testutil/nullify"
	"xnet/x/xnet"
	"xnet/x/xnet/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.XnetKeeper(t)
	xnet.InitGenesis(ctx, *k, genesisState)
	got := xnet.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
