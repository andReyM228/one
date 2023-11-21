package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "one/testutil/keeper"
	"one/x/allowed/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AllowedKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
