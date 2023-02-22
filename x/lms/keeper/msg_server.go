package keeper

import (
	"context"

	"github.com/pavania1/cosmos-leave-management/x/lms/types"
)

func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &msgServer{
		Keeper: k,
	}
}

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
	types.UnimplementedMsgServer
}

func (k msgServer) AddStudent(goCtx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	// if k.students!=req.Students{
	// 	return nil, errors.Wrapf(k.AddStudent,)
	// }
	// if err != nil {
	// 	return nil, err
	// }

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
