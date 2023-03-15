package keeper

import (
	"fmt"
	"strconv"

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
		//sdkaddress := sdk.AccAddress(req.Address).String()
		store.Set(types.StudentStoreKey(req.Address), bz)
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
	}

	// sdkaddress := sdk.AccAddress(req.Address).String()
	key := types.AdminstoreKey(req.Address)
	store.Set(key, bz)
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
	// if _, err := sdk.AccAddressFromBech32(req.Admin); err != nil {
	// 	panic(fmt.Errorf("invalid  authority address: %w", err))
	// }
	store := ctx.KVStore(k.storeKey)
	req.Status = types.LeaveStatus_STATUS_ACCEPTED
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.AcceptedLeavesStoreKey(req.Admin, fmt.Sprint(req.LeaveId)), bz)
		r := store.Get(types.AcceptedLeavesStoreKey(req.Admin, fmt.Sprint(req.LeaveId)))
		var res types.AcceptLeaveRequest
		k.cdc.Unmarshal(r, &res)
		fmt.Println(res)
	}
	return nil
}

// ----------------------->> APPLY LEAVE   <<-----------------------------
func (k Keeper) ApplyLeave(ctx sdk.Context, req *types.ApplyLeaveRequest) error {
	// if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
	// 	panic(fmt.Errorf("invalid  authority address: %w", err))
	// } else {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	}
	addr := types.LeavesCounterKey(req.Address)
	counter := store.Get(addr)
	if counter == nil {
		store.Set(addr, []byte("1"))
	} else {
		c, err := strconv.Atoi(string(counter))
		if err != nil {
			panic(err)
		}
		c = c + 1
		store.Set(addr, []byte(fmt.Sprint(c)))
	}
	counter = store.Get(addr)
	store.Set(types.AppliedLeavesStoreKey(req.LeaveId, string(counter)), bz)
	// r := store.Get(types.AppliedLeavesStoreKey(req.LeaveId, string(counter)))
	// var res types.ApplyLeaveRequest
	// k.cdc.Unmarshal(r, &res)
	// panic(res)

	// panic(leaves)
	// }
	return nil
}

//--------------------->> GET STUDENTS  <<---------------------------------------------
func (k Keeper) Getstudents(ctx sdk.Context, getstudents *types.GetstudentsRequest) []*types.AddStudentRequest {
	store := ctx.KVStore(k.storeKey)
	var students []*types.AddStudentRequest
	itr := sdk.KVStorePrefixIterator(store, types.StudentKey)
	defer itr.Close()
	for ; itr.Valid(); itr.Next() {
		var student types.AddStudentRequest
		k.cdc.Unmarshal(itr.Value(), &student)
		students = append(students, &student)
		fmt.Println("Students are:", student)
	}
	return students
}

// -------------------->> GET STUDENT <<-----------------------------
func (k Keeper) Getstdent(ctx sdk.Context, Address string) (req types.AddStudentRequest, err error) {
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

func (k Keeper) GetLeaveRqst(ctx sdk.Context, getLeaves *types.GetLeaveRequest) []*types.ApplyLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var leaves []*types.ApplyLeaveRequest
	itr := sdk.KVStorePrefixIterator(store, types.AppliedLeavesKey)
	defer itr.Close()
	// itr := store.Iterator(types.AppliedLeavesKey, nil)
	for ; itr.Valid(); itr.Next() {
		var leave types.ApplyLeaveRequest
		k.cdc.Unmarshal(itr.Value(), &leave)
		leaves = append(leaves, &leave)
		fmt.Println(leaves)
	}
	// panic(leaves)
	return leaves
}

// -------------------------->> GET APPROVE LEAVES <<-------------------------------
func (k Keeper) GetAcceptLeaves(ctx sdk.Context, getleaves *types.GetLeaveApproveRequest) []*types.AcceptLeaveRequest {
	store := ctx.KVStore(k.storeKey)
	var approve []*types.AcceptLeaveRequest
	itr := sdk.KVStoreReversePrefixIterator(store, types.AcceptedLeavesKey)
	//itr := store.Iterator(types.AcceptedLeavesKey, nil)
	for ; itr.Valid(); itr.Next() {
		var approveleave types.AcceptLeaveRequest
		k.cdc.Unmarshal(itr.Value(), &approveleave)
		approve = append(approve, &approveleave)
		fmt.Println(approve)
	}
	return approve
}

// -------------------------------->> GET ADMIN <<---------------------------------------
func (k Keeper) GetAdminn(ctx sdk.Context, Address string) (req types.RegisterAdminRequest, err error) {
	if _, err := sdk.AccAddressFromBech32(Address); err != nil {
		panic(err)
	}
	store := ctx.KVStore(k.storeKey)
	admin := store.Get(types.AdminstoreKey(Address))
	if admin == nil {
		panic("no admin")
	}

	if err := k.cdc.Unmarshal(admin, &req); err != nil {
		return req, err
	}

	fmt.Println(req)
	return req, err

}
func (k Keeper) GetLeavestatus(ctx sdk.Context, Admin string, LeaveId string) (req types.AcceptLeaveRequest, err error) {
	store := ctx.KVStore(k.storeKey)
	leavestatus := store.Get(types.AcceptedLeavesStoreKey(Admin, LeaveId))
	if leavestatus == nil {
		panic("No Leaves")
	}
	if err := k.cdc.Unmarshal(leavestatus, &req); err != nil {
		return req, err
	}
	fmt.Println(req)
	return req, err
}

// func (k Keeper) GetAdmins(ctx sdk.Context) []*types.GetRegisterAdminRequest {
// 	fmt.Println("step5.1")

// 	store := ctx.KVStore(k.storeKey)
// 	val := store.Get(types.AdminKey)
// 	fmt.Println("val", val)

// 	fmt.Println("step5.2", k.storeKey, store)

// 	itr := sdk.KVStorePrefixIterator(store, types.AdminKey)
// 	defer itr.Close()

// 	fmt.Println("step5.3")

// 	if itr == nil {
// 		panic("prefix iterator is empty")
// 	}
// 	fmt.Println("step5.4")

// 	for ; itr.Valid(); itr.Next() {
// 		panic(fmt.Sprintf("asdfasdf", itr.Key()))
// 		// panic(fmt.Sprintf("asdfasdf.value", string(prefixItr.Value())))
// 		// prefixItr.Key()
// 		// prefixItr.Next()
// 	}

// 	return []*types.GetRegisterAdminRequest{}
// 	// key := types.AdminstoreKey(req.Address)
// }

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
