package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrExecutionFailed = sdkerrors.Register(ModuleName, 1, "Script failed to execute in Move VM")
)
