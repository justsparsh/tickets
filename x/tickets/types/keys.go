package types

const (
	// ModuleName defines the module name
	ModuleName = "tickets"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_tickets"
)

var (
	ParamsKey = []byte("p_tickets")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const TicketKey = "Ticket/value"
