package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is the module name router key
const RouterKey = ModuleName // this was defined in your key.go file

// MsgSetName defines a SetName message
type MsgRunMoveScript struct {
	Script string         `json:"script"`
	Args   []string         `json:"args"`
	Address sdk.AccAddress `json:"address"`
}

// NewMsgRunMoveScript is a constructor function for MsgRunMoveScript
func NewMsgRunMoveScript(script string, args []string, address sdk.AccAddress) MsgRunMoveScript {
	return MsgRunMoveScript{
		Script:  script,
		Args: args,
		Address: address,
	}
}

// Route should return the name of the module
func (msg MsgRunMoveScript) Route() string { return RouterKey }

// Type should return the action
func (msg MsgRunMoveScript) Type() string { return "run_move_script" }

// ValidateBasic runs stateless checks on the message
func (msg MsgRunMoveScript) ValidateBasic() error {
	if msg.Address.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Address.String())
	}
	if len(msg.Script) == 0 || len(msg.Script) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Script cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgRunMoveScript) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgRunMoveScript) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}
