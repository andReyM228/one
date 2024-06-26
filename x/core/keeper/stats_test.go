package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/andReyM228/one/testutil/keeper"
	"github.com/andReyM228/one/testutil/nullify"
	"github.com/andReyM228/one/x/core/keeper"
	"github.com/andReyM228/one/x/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStats(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Stats {
	items := make([]types.Stats, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStats(ctx, items[i])
	}
	return items
}

func TestStatsGet(t *testing.T) {
	keeper, ctx := keepertest.CoreKeeper(t)
	items := createNStats(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStats(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStatsRemove(t *testing.T) {
	keeper, ctx := keepertest.CoreKeeper(t)
	items := createNStats(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStats(ctx,
			item.Index,
		)
		_, found := keeper.GetStats(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStatsGetAll(t *testing.T) {
	keeper, ctx := keepertest.CoreKeeper(t)
	items := createNStats(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStats(ctx)),
	)
}
