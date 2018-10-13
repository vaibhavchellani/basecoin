package checkpoint

import (
	"encoding/json"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		// NOTE msg already has validate basic run
		switch msg := msg.(type) {
		case MsgCheckpoint:
			return handleMsgCheckpoint(ctx, msg, k)
		default:
			return sdk.ErrTxDecode("Invalid message in checkpoint module ").Result()
		}
	}
}

func handleMsgCheckpoint(ctx sdk.Context, msg MsgCheckpoint, k Keeper) sdk.Result {
	fmt.Printf("entered handler with message %v", msg)
	//TODO check last block in last checkpoint (startBlock of new checkpoint == last block of prev endpoint)
	// TODO insert checkpoint in state
	logger := ctx.Logger().With("module", "x/baseapp")
	//valid := validateCheckpoint(msg.StartBlock,msg.EndBlock,msg.RootHash)

	var res int64
	valid := true
	if valid {
		logger.Error("root hash matched !! ")
		res = k.addCheckpoint(ctx, msg.StartBlock, msg.EndBlock, msg.RootHash)
	} else {
		logger.Error("Root hash no match ;(")
		return ErrBadBlockDetails(k.codespace).Result()

	}
	var out CheckpointBlockHeader
	json.Unmarshal(k.getCheckpoint(ctx, res), &out)
	fmt.Printf("******* block end added is *******", out.EndBlock)

	//TODO add validation
	// send tags
	return sdk.Result{}
}
