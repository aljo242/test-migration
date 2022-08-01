package testmigration_test

import (
	"testing"

	keepertest "test-migration/testutil/keeper"
	"test-migration/testutil/nullify"
	"test-migration/x/testmigration"
	"test-migration/x/testmigration/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestmigrationKeeper(t)
	testmigration.InitGenesis(ctx, *k, genesisState)
	got := testmigration.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
