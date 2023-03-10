package keeper

import (
	"context"

	"github.com/pavania1/cosmos-LMS/x/lms/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
	types.UnimplementedMsgServer
}

func (k msgServer) AddStudent(context.Context, *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	return &types.AddStudentResponse{}, nil
}
func (k msgServer) RegisterAdmin(context.Context, *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	return &types.RegisterAdminResponse{}, nil
}
func (k msgServer) ApplyLeave(context.Context, *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	return &types.ApplyLeaveResponse{}, nil
}
func (k msgServer) AcceptLeave(context.Context, *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	return &types.AcceptLeaveResponse{}, nil
}
