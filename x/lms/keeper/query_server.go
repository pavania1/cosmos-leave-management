package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pavania1/cosmos-leave-management/x/lms/types"
)

var _ types.QueryServer = Keeper{}

// type queryServer struct {
// 	Keeper
// 	types.UnimplementedQueryServer
// }

//    Get Student
func (k Keeper) GetStudent(goCtx context.Context, req *types.GetstudentRequest) (*types.GetstudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	fmt.Println("getStudent")
	student, _ := k.Getstdent(ctx, req.Address)
	res := types.GetstudentResponse{
		Student: &student,
	}
	return &res, nil
}

//        Get Admin
func (k Keeper) GetAdmin(goCtx context.Context, req *types.GetRegisterAdminRequest) (*types.GetRegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	admin, _ := k.GetAdminn(ctx, req.Address)
	// if err != nil {
	// 	//panic("------------->")
	// 	panic(err)
	// }
	res := types.GetRegisterAdminResponse{
		Admin: &admin,
	}
	return &res, nil
}

// Get Leave(Apply)
func (k Keeper) GetLeave(goCtx context.Context, req *types.GetLeaveRequest) (*types.GetLeaveRequestResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	res := k.GetLeaveRqst(ctx, req)
	// panic(res)
	return &types.GetLeaveRequestResponse{
		Leaverequest: res,
	}, nil
}

// Get Accept Leave
func (k Keeper) GetAcceptLeave(goCtx context.Context, req *types.GetLeaveApproveRequest) (*types.GetLeaveApproveResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	//k.GetAcceptLeaves(ctx, req)
	res := k.GetAcceptLeaves(ctx, req)
	return &types.GetLeaveApproveResponse{
		Getleaverequest: res,
	}, nil
}

// Get All Students
func (k Keeper) GetStudents(goCtx context.Context, req *types.GetstudentsRequest) (*types.GetstudentsResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	students := k.Getstudents(ctx, req)
	res := types.GetstudentsResponse{
		Student: students,
	}
	return &res, nil

}
