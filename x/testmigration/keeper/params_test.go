package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "test-migration/testutil/keeper"
	"test-migration/x/testmigration/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TestmigrationKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
