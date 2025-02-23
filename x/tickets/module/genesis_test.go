package tickets_test

import (
	"testing"

	keepertest "tickets/testutil/keeper"
	"tickets/testutil/nullify"
	tickets "tickets/x/tickets/module"
	"tickets/x/tickets/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TicketsKeeper(t)
	tickets.InitGenesis(ctx, k, genesisState)
	got := tickets.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
