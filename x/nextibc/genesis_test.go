package nextibc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "xnet/testutil/keeper"
	"xnet/testutil/nullify"
	"xnet/x/nextibc"
	"xnet/x/nextibc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.NextibcKeeper(t)
	nextibc.InitGenesis(ctx, *k, genesisState)
	got := nextibc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
