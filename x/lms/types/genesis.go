package types

//var _ types.UnpackInterfacesMessage = GenesisState{}

func NewGenesisState(responses []*AcceptLeaveRequest) *GenesisState {
	return &GenesisState{
		Admin: responses,
	}
}

// // ValidateGenesis check the given genesis state has no integrity issues
func ValidateGenesis(data *GenesisState) error {
	return nil
}

// // DefaultGenesisState - Return a default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}
