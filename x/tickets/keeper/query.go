package keeper

import (
	"tickets/x/tickets/types"
)

var _ types.QueryServer = Keeper{}
