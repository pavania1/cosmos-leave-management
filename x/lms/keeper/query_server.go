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

func (k Keeper) GetStudent(goCtx context.Context, req *types.GetstudentRequest) (*types.GetstudentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	fmt.Println("getStudent")
	k.Getstudent(ctx, req.Address)
	k.GetStudent(ctx, req)
	return &types.GetstudentResponse{}, nil
}

func (k Keeper) GetAdmin(goCtx context.Context, req *types.GetRegisterAdminRequest) (*types.GetRegisterAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetAdminn(ctx, req.Name)
	return &types.GetRegisterAdminResponse{}, nil
}

func (k Keeper) GetLeave(goCtx context.Context, req *types.GetLeaveRequest) (*types.GetLeaveRequestResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("Empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetLeaveRqst(ctx, req)
	return &types.GetLeaveRequestResponse{}, nil
}

func (k Keeper) GetAcceptLeave(goCtx context.Context, req *types.GetLeaveApproveRequest) (*types.GetLeaveApproveResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	k.GetAcceptLeaves(ctx, req)
	return &types.GetLeaveApproveResponse{}, nil
}

func (k Keeper) GetStudents(goCtx context.Context, req *types.GetstudentsRequest) (*types.GetstudentsResponse, error) {
	if req == nil {
		return nil, sdkerrors.ErrInvalidRequest.Wrap("empty request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	student1 := []*types.Student{}

	students := k.Getstudents(ctx, req)
	for _, stud := range students {
		student1 = append(student1, &stud)
	}
	res := types.GetstudentsResponse{
		Student: student1,
	}
	k.Getstudents(ctx, req)

	return &res, nil
}
