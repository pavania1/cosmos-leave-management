package types

import (
	"errors"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = &AddStudentRequest{}
	_ sdk.Msg = &AcceptLeaveRequest{}
	_ sdk.Msg = &ApplyLeaveRequest{}
	_ sdk.Msg = &RegisterAdminRequest{}
)

func NewAddStudentRequest(Admin string, name string, address string, id string) *AddStudentRequest {
	return &AddStudentRequest{
		Admin:   Admin,
		Address: address,
		Name:    name,
		Id:      id,
	}
}
func NewRegisterAdminRequest(Address string, name string) *RegisterAdminRequest {
	return &RegisterAdminRequest{
		Address: Address,
		Name:    name,
	}
}
func NewApplyLeaveRequest(Address string, Reason string, From *time.Time, To *time.Time) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Address: Address,
		Reason:  Reason,
		From:    From,
		To:      To,
	}
}
func NewAcceptLeaveRequest(Admin string, LeaveId uint64, Status LeaveStatus) *AcceptLeaveRequest {
	return &AcceptLeaveRequest{
		Admin:   Admin,
		LeaveId: LeaveId,
		Status:  Status,
	}
}
func NewGetStudentRequest(Id string, Address string) *GetstudentRequest {
	return &GetstudentRequest{
		Id:      Id,
		Address: Address,
	}
}

func NewGetAdminRequest(name string) *GetRegisterAdminRequest {
	return &GetRegisterAdminRequest{
		Name: name,
	}

}
func NewGetStudentsRequest() *GetstudentsRequest {
	return &GetstudentsRequest{}
}

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
	} else if msg.Name == "" {
		return errors.New("name cant be null")
	} else if msg.Admin == "" {
		return errors.New("Admin cant be null")
	} else if msg.Address == "" {
		return errors.New("address cant be null")

	} else if msg.Id == "" {
		return errors.New("Id cant be null")
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
	} else if msg.Admin == "" {
		return errors.New("name cant be null")
	} else if msg.LeaveId == 0 {
		return errors.New("Leave id cant be null")
	} else if msg.Status == 0 {
		return errors.New("status must be specified")
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
	} else if msg.Address == "" {
		return errors.New("address cant be null")
	} else if msg.Reason == "" {
		return errors.New("Reason cant be null")
	} else if msg.From == nil {
		return errors.New("form date is must not be nill")
	} else if msg.To == nil {
		return errors.New("to date is must be mentioned")
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
	} else if msg.Name == "" {
		return errors.New("name cant be null")
	} else if msg.Address == "" {
		return errors.New("address cant be null")
	}
	return nil
}
