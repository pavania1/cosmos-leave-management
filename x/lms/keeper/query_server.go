package keeper

import (
	"context"

	"github.com/pavania1/cosmos-leave-management/x/lms/types"
)

var _ types.QueryServer = Keeper{}

// type queryServer struct {
// 	Keeper
// 	types.UnimplementedQueryServer
// }

func (k Keeper) GetStudent(context.Context, *types.GetstudentRequest) (*types.GetstudentResponse, error) {
	return &types.GetstudentResponse{}, nil
}

func (k Keeper) GetAdmin(context.Context, *types.GetRegisterAdminRequest) (*types.GetRegisterAdminResponse, error) {
	return &types.GetRegisterAdminResponse{}, nil
}

func (k Keeper) GetLeave(context.Context, *types.GetLeaveRequest) (*types.GetLeaveRequestResponse, error) {
	return &types.GetLeaveRequestResponse{}, nil
}

func (k Keeper) GetAcceptLeave(context.Context, *types.GetLeaveApproveRequest) (*types.GetLeaveApproveResponse, error) {
	return &types.GetLeaveApproveResponse{}, nil
}

func (k Keeper) GetStudents(context.Context, *types.GetstudentsRequest) (*types.GetstudentsResponse, error) {
	return &types.GetstudentsResponse{}, nil
}
