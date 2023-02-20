package keeper

import (
	"context"

	"github.com/pavania1/cosmos-LMS/x/lms/types"
)

var _ types.QueryServer = &queryServer{}

type queryServer struct {
	Keeper
	types.UnimplementedQueryServer
}

func (k queryServer) GetStudent(context.Context, *types.GetstudentsRequest) (*types.GetstudentsResponse, error) {
	return &types.GetstudentsResponse{}, nil
}

//func (k queryServer) RegisterAdmin(context.Context, *types.GetRegisterAdminRequest) (*types.GetRegisterAdminResponse, error) {
//	return &types.GetRegisterAdminResponse{}, nil
//}

func (k queryServer) ApplyLeave(context.Context, *types.GetLeaveRequest) (*types.GetLeaveRequestResponse, error) {
	return &types.GetLeaveRequestResponse{}, nil
}

func (k queryServer) AcceptLeave(context.Context, *types.GetLeaveApproveRequest) (*types.GetLeaveApproveResponse, error) {
	return &types.GetLeaveApproveResponse{}, nil
}
