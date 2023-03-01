package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pavania1/cosmos-leave-management/x/lms/types"
)

// func NewMsgServerImpl(k Keeper) types.MsgServer {
// 	return &msgServer{
// 		Keeper: k,
// 	}
// }

var _ types.MsgServer = Keeper{}

// type msgServer struct {
// 	Keeper
// 	types.UnimplementedMsgServer
// }

func (k Keeper) MsgAddStudent(goCtx context.Context, req *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	// if k.students!=req.Students{
	// 	return nil, errors.Wrapf(k.AddStudent,)
	// }
	// if err != nil {
	// 	return nil, err
	// }
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.AddStudent(ctx, req)
	return &types.AddStudentResponse{}, nil

}
func (k Keeper) MsgRegisterAdmin(goCtx context.Context, req *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.AdminRegister(ctx, req)
	return &types.RegisterAdminResponse{}, nil
}
func (k Keeper) MsgApplyLeave(goCtx context.Context, req *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.ApplyLeave(ctx, req)
	return &types.ApplyLeaveResponse{}, nil
}
func (k Keeper) MsgAcceptLeave(context.Context, *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {

	return &types.AcceptLeaveResponse{}, nil
}
