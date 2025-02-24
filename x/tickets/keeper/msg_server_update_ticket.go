package keeper

import (
	"context"
	"fmt"

	"tickets/x/tickets/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateTicket(goCtx context.Context, msg *types.MsgUpdateTicket) (*types.MsgUpdateTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var ticket = types.Ticket{
		Id:    msg.Id,
		Name:  msg.Name,
		Event: msg.Event,
	}
	_, found := k.GetTicket(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	// if msg.Creator != val.Creator {
	//     return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }
	k.SetTicket(ctx, ticket)
	return &types.MsgUpdateTicketResponse{}, nil
}
