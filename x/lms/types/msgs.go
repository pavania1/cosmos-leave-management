package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &AddStudentRequest{}
	_ sdk.Msg = &AcceptLeaveRequest{}
	_ sdk.Msg = &ApplyLeaveRequest{}
	_ sdk.Msg = &RegisterAdminRequest{}
)

// GetSignBytes implements the LegacyMsg interface.
func (msg AddStudentRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg *AddStudentRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32("")
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg *AddStudentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(""); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}
func (msg AcceptLeaveRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// GetSigners returns the expected signers for a MsgUpdateParams message.
func (msg AcceptLeaveRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32("")
	return []sdk.AccAddress{addr}
}

// ValidateBasic does a sanity check on the provided data.
func (msg AcceptLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(""); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}

// similarly for ApplyLeaveRequest

func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
func (msg *ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32("")
	return []sdk.AccAddress{addr}
}
func (msg *ApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(""); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}

// for Admin

func (msg RegisterAdminRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}
func (msg *RegisterAdminRequest) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32("")
	return []sdk.AccAddress{addr}
}
func (msg *RegisterAdminRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(""); err != nil {
		return sdkerrors.Wrap(err, "invalid authority address")
	}
	return nil
}
