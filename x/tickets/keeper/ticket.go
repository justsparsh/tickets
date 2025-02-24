package keeper

import (
	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"

	"tickets/x/tickets/types"
)

func (k Keeper) AppendTicket(ctx sdk.Context, ticket types.Ticket) string {
	id := uuid.NewString()
	ticket.Id = id
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketKey))
	appendedValue := k.cdc.MustMarshal(&ticket)
	store.Set([]byte(id), appendedValue)
	return id
}

func (k Keeper) GetTicket(ctx sdk.Context, id string) (val types.Ticket, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketKey))
	b := store.Get([]byte(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetTicket(ctx sdk.Context, ticket types.Ticket) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketKey))
	t := k.cdc.MustMarshal(&ticket)
	store.Set([]byte(ticket.Id), t)
}

func (k Keeper) RemoveTicket(ctx sdk.Context, id string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketKey))
	store.Delete([]byte(id))
}
