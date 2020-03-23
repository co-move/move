package move

import (
	"encoding/json"
	"fmt"

	"github.com/co-move/move/x/move/internal/types"

	"github.com/co-move/move/ship"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns a handler for "nameservice" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgRunMoveScript:
			return handleMsgRunMoveScript(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized Move Msg type: %v", msg.Type()))
		}
	}
}

// Handle a message to run Move script
func handleMsgRunMoveScript(ctx sdk.Context, keeper Keeper, msg types.MsgRunMoveScript) (*sdk.Result, error) {

	jsontxt := ship.RunScript(msg.Script)
	var res ship.VMResult
	err := json.Unmarshal([]byte(jsontxt), &res)
	if err != nil {
		return nil, err
	}
	if res.Code > 0 {
		return nil, sdkerrors.Wrap(types.ErrExecutionFailed, res.Msg)
	}

	return &sdk.Result{}, nil // return
}
