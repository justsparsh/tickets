package keeper

import (
	"context"
	"fmt"

	"tickets/x/tickets/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteTicket(goCtx context.Context, msg *types.MsgDeleteTicket) (*types.MsgDeleteTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetTicket(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	// if msg.Creator != val.Creator {
	//     return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	// }
	k.RemoveTicket(ctx, msg.Id)

	return &types.MsgDeleteTicketResponse{}, nil
}
