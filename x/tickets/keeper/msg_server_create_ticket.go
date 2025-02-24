package keeper

import (
	"context"

	"tickets/x/tickets/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTicket(goCtx context.Context, msg *types.MsgCreateTicket) (*types.MsgCreateTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var ticket = types.Ticket{
		Name:  msg.Name,
		Event: msg.Event,
	}
	id := k.AppendTicket(
		ctx,
		ticket,
	)
	return &types.MsgCreateTicketResponse{
		Id: id,
	}, nil
}
