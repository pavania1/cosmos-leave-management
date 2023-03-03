package types

const (
	ModuleName = "leavemanagementsystem"
	StoreKey   = ModuleName
	// RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	AdminKey          = []byte{0x01}
	StudentKey        = []byte{0x02}
	LeaveKey          = []byte{0x03}
	AcceptedLeavesKey = []byte{0x04}
	AppliedLeavesKey  = []byte{0x05}
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
func leavestoreKey(leave string) []byte {
	key := make([]byte, len(LeaveKey)+len(leave))
	copy(key, LeaveKey)
	copy(key[len(LeaveKey):], []byte(leave))
	return key
}

// --------------------------------------------->> Accept leave key <<-----------------------------------
func AcceptLeaveStoreKey(leave string) []byte {
	key := make([]byte, len(AcceptedLeavesKey)+len(leave))
	copy(key, AcceptedLeavesKey)
	copy(key[len(AcceptedLeavesKey):], []byte(leave))
	return key
}

// ----------------------------------->> Apply leave key <<---------------------------------------
func AppliedLeavesStoreKey(leave string) []byte {
	key := make([]byte, len(AppliedLeavesKey)+len(leave))
	copy(key, AppliedLeavesKey)
	copy(key[len(AppliedLeavesKey):], []byte(leave))
	return key
}
