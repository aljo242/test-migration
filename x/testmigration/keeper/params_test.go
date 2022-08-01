package keeper_test

import (
	"testing"

	testkeeper "test-migration/testutil/keeper"
	"test-migration/x/testmigration/types"

	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TestmigrationKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
