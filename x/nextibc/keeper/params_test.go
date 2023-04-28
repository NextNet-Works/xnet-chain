package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "xnet/testutil/keeper"
	"xnet/x/nextibc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NextibcKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
