package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "tickets/testutil/keeper"
	"tickets/x/tickets/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.TicketsKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
