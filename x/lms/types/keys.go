package types

const (
	ModuleName = "leavemanagementsystem"
	StoreKey   = ModuleName
	// RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	AdminKey   = []byte{0x01}
	StudentKey = []byte{0x02}
	//LeaveKey          = []byte{0x03}
	AcceptedLeavesKey = []byte{0x04}
	AppliedLeavesKey  = []byte{0x05}
	CounterKey        = []byte{0x03}
	SequenceKey       = []byte{0x06}
)

//------------------------->> Studentkey <<-----------------------------------------------
func StudentStoreKey(studentAddress string) []byte {
	key := make([]byte, len(StudentKey)+len(studentAddress))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], []byte(studentAddress))
	return key
}

// --------------------------------->> Admin key <<-----------------------------------------
func AdminstoreKey(admin string) []byte {
	key := make([]byte, len(AdminKey)+len(admin))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], []byte(admin))
	return key
}

// --------------------------------->> Leave key <<-----------------------------------------
// func leavestoreKey(leave string) []byte {
// 	key := make([]byte, len(LeaveKey)+len(leave))
// 	copy(key, LeaveKey)
// 	copy(key[len(LeaveKey):], []byte(leave))
// 	return key
// }

// --------------------------------------------->> Accept leave key <<-----------------------------------

func AcceptedLeavesStoreKey(admin string, leaveId string) []byte {
	key := make([]byte, len(AcceptedLeavesKey)+len(admin)+len(leaveId))
	copy(key, AcceptedLeavesKey)
	copy(key[len(AcceptedLeavesKey):], []byte(admin))
	copy(key[len(AcceptedLeavesKey)+len(admin):], []byte(leaveId))
	return key
}

// ----------------------------------->> Apply leave key <<---------------------------------------
func AppliedLeavesStoreKey(leaveId string, leavesCount string) []byte {
	key := make([]byte, len(AppliedLeavesKey)+len(leaveId)+len(leavesCount))
	copy(key, AppliedLeavesKey)
	copy(key[len(AppliedLeavesKey):], []byte(leaveId))
	copy(key[len(AppliedLeavesKey)+len(leaveId):], []byte(leavesCount))
	return key
}
func LeavesCounterKey(id string) []byte {
	key := make([]byte, len(CounterKey)+len(id))
	copy(key, CounterKey)
	copy(key[len(CounterKey):], []byte(id))
	return key
}
