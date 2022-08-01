package keeper

import (
	"test-migration/x/testmigration/types"
)

var _ types.QueryServer = Keeper{}
