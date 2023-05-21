package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"himitsu/x/himitsu/types"
)

func (k msgServer) SendCidHash(goCtx context.Context, msg *types.MsgSendCidHash) (*types.MsgSendCidHashResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgSendCidHashResponse{}, nil
}
