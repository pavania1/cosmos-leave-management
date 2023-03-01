package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pavania1/cosmos-leave-management/x/lms/types"
	//"google.golang.org/genproto/googleapis/actions/sdk/v2"
)

//NewKeeper

func NewKeeper(cdc codec.BinaryCodec, storekey storetypes.StoreKey) Keeper {
	if _, err := sdk.AccAddressFromBech32("h"); err != nil {
		panic(fmt.Errorf("invalid  authority address: %w", err))
	}
	return Keeper{
		cdc:      cdc,
		storeKey: storekey,
	}
}

func (k Keeper) AddStudent(ctx sdk.Context, req *types.AddStudentRequest) error {
	// if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
	// 	panic(fmt.Errorf("Invalid Authority Address:%w", err))
	// }

	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		// store.Set(types.StudentKey, bz)
		sdkaddress := sdk.AccAddress(req.Address).String()
		store.Set(types.StudentStoreKey(sdkaddress), bz)
	}
	return nil
}
func (k Keeper) CheckAddstudent(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	sdkaddress := sdk.AccAddress(address).String()

	val := store.Get(types.StudentStoreKey(sdkaddress))
	if val == nil {
		fmt.Println("student did not Added")
		return false
	} else {
		fmt.Println("student added")
		return true
	}
}

func (k Keeper) AdminRegister(ctx sdk.Context, req *types.RegisterAdminRequest) error {
	// if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
	// 	panic(fmt.Errorf("invalid authority address: %w", err))
	// }
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		sdkaddress := sdk.AccAddress(req.Address).String()
		store.Set(types.AdminstoreKey(sdkaddress), bz)

		// val := store.Get(types.AdminstoreKey(sdkaddress))
		// ans := types.RegisterAdminRequest{}
		// k.cdc.Unmarshal(val, &ans)
		// fmt.Println("the marshalled get data  is", ans)
	}
	return nil
}
func (k Keeper) CheckAdminregister(ctx sdk.Context, address string) bool {
	store := ctx.KVStore(k.storeKey)
	sdkaddress := sdk.AccAddress(address).String()

	val := store.Get(types.AdminstoreKey(sdkaddress))
	if val == nil {
		fmt.Println("admin did not register")
		return false
	} else {
		fmt.Println("admin registered")
		return true
	}
}
func (k Keeper) AcceptLeave(ctx sdk.Context, req *types.AcceptLeaveRequest) error {
	// if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
	// 	panic(fmt.Errorf("invalid  authority address: %w", err))
	// }
	store := ctx.KVStore(k.storeKey)
	req.Status = types.LeaveStatus_STATUS_ACCEPTED
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.LeaveKey, bz)
	}
	return nil
}
func (k Keeper) ApplyLeave(ctx sdk.Context, req *types.ApplyLeaveRequest) error {
	// if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
	// 	panic(fmt.Errorf("invalid  authority address: %w", err))
	// }
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.LeaveKey, bz)
	}
	return nil
}

// func (k Keeper) CheckLeaveStatus(ctx sdk.Context, studentAddress string) (types.Leave, error) {
// 	// if _, err := sdk.AccAddressFromBech32(studentAddress); err != nil {
// 	// 	return types.Leave{}, err
// 	}

// 	store := ctx.KVStore(k.leavestoreKey)
// 	if store.Get(types.StudentStoreKey(studentAddress)) == nil {
// 		return types.Leave{}, types.ErrStudentDoesNotExist
// 	}

// 	leaveId, _ := strconv.Atoi(string(store.Get(types.LeaveCounterStoreKey(studentAddress))))

// 	// if a student never applied for a leave
// 	if leaveId == 0 {
// 		return types.Leave{}, types.ErrLeaveNeverApplied
// 	}

// 	val := store.Get(types.LeaveStoreKey(studentAddress, leaveId))
// 	var leave *types.MsgAcceptLeaveRequest
// 	k.cdc.Unmarshal(val, leave)
// 	return types.Leave{
// 		Address: leave.Admin,
// 		Status:  (leave.Status == types.LeaveStatus_STATUS_ACCEPTED),
// 	}, nil
// }
