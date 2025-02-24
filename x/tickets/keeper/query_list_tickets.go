package keeper

import (
	"context"

	"tickets/x/tickets/types"

	// sdk "github.com/cosmos/cosmos-sdk/types"
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	// "github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListTickets(goCtx context.Context, req *types.QueryListTicketsRequest) (*types.QueryListTicketsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketKey))

	var tickets []types.Ticket
	iterator := store.Iterator(nil, nil)

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ticket types.Ticket
		if err := k.cdc.Unmarshal(iterator.Value(), &ticket); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		tickets = append(tickets, ticket)
	}

	return &types.QueryListTicketsResponse{Ticket: tickets}, nil
}
