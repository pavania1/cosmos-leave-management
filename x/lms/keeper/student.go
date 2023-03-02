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
	return Keeper{
		cdc:      cdc,
		storeKey: storekey,
	}
}

// -------------------------->> ADD STUDENT <<-----------------------------------
func (k Keeper) AddStudent(ctx sdk.Context, req *types.AddStudentRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		panic(fmt.Errorf("Invalid Authority Address:%w", err))
	}

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

// --------------------------> ADD The ADMIN <<----------------------------

func (k Keeper) AdminRegister(ctx sdk.Context, req *types.RegisterAdminRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		panic(fmt.Errorf("invalid authority address: %w", err))
	}
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		sdkaddress := sdk.AccAddress(req.Address).String()
		store.Set(types.AdminstoreKey(sdkaddress), bz)
	}
	return nil
}

// -------------------->> THIS function is used in test file (checking the admin) <<---------------------
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

// ----------------------->> ACCEPT LEAVE <<------------------------------
func (k Keeper) AcceptLeave(ctx sdk.Context, req *types.AcceptLeaveRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
		panic(fmt.Errorf("invalid  authority address: %w", err))
	}
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

// ----------------------->> APPLY LEAVE   <<-----------------------------
func (k Keeper) ApplyLeave(ctx sdk.Context, req *types.ApplyLeaveRequest) error {
	if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
		panic(fmt.Errorf("invalid  authority address: %w", err))
	}
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.LeaveKey, bz)
	}
	return nil
}

//--------------------->> GET STUDENTS  <<---------------------------------------------
func (k Keeper) Getstudents(ctx sdk.Context, getstudents *types.GetstudentsRequest) []types.Student {
	store := ctx.KVStore(k.storeKey)
	var students []types.Student
	itr := store.Iterator(types.StudentKey, nil)
	for ; itr.Valid(); itr.Next() {
		var student types.Student
		k.cdc.Unmarshal(itr.Value(), &student)
		fmt.Println("Students are:", student)
	}
	return students
}

// -------------------->> GET STUDENT <<-----------------------------
func (k Keeper) Getstudent(ctx sdk.Context, Address string) (req types.Student, err error) {
	if _, err := sdk.AccAddressFromBech32(Address); err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storeKey)
	student := store.Get(types.StudentStoreKey(Address))
	fmt.Println(student)
	if student == nil {
		panic("student not have")
	}
	fmt.Println(student)
	k.cdc.MustUnmarshal(student, &req)
	fmt.Println(req)
	return req, err
}

//<------------------GET LEAVE REQUESTS------------------------------->//

func (k Keeper) GetLeaveRqst(ctx sdk.Context, getLeaves *types.GetLeaveRequest) {
	store := ctx.KVStore(k.storeKey)
	var leaves types.ApplyLeaveRequest
	itr := store.Iterator(types.AppliedLeavesKey, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &leaves)
		fmt.Println(leaves)
	}
}

// -------------------------->> GET APPROVE LEAVES <<-------------------------------
func (k Keeper) GetAcceptLeaves(ctx sdk.Context, getleaves *types.GetLeaveApproveRequest) {
	store := ctx.KVStore(k.storeKey)
	var approve types.AcceptLeaveRequest
	itr := store.Iterator(types.AcceptedLeavesKey, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &approve)
		fmt.Println(approve)
	}
}

// -------------------------------->> GET ADMIN <<---------------------------------------
func (k Keeper) GetAdminn(ctx sdk.Context, Address string) []byte {
	if _, err := sdk.AccAddressFromBech32(Address); err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storeKey)

	// ans := store.Get(types.AdminKey)
	// var admin types.Admin
	// k.cdc.Unmarshal(ans, &admin)
	return store.Get(types.AdminstoreKey(Address))
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
