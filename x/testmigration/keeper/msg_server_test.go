package keeper_test

import (
	"context"
	"testing"

	keepertest "test-migration/testutil/keeper"
	"test-migration/x/testmigration/keeper"
	"test-migration/x/testmigration/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TestmigrationKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
