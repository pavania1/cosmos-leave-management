package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&AddStudentRequest{}, "cosmos-sdk/Addstudentrequest", nil)
	cdc.RegisterConcrete(&RegisterAdminRequest{}, "cosmos-sdk/RegisterAdmin", nil)
	cdc.RegisterConcrete(&ApplyLeaveRequest{}, "cosmos-sdk/ApplyLeave", nil)
	cdc.RegisterConcrete(&AcceptLeaveRequest{}, "cosmos-sdk/AcceptLeave", nil)
}
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		&AddStudentRequest{},
		&RegisterAdminRequest{},
		&ApplyLeaveRequest{},
		&AcceptLeaveRequest{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &Msg_ServiceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	//cryptocodec.RegisterCrypto(amino)
	//RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
	RegisterLegacyAminoCodec(authzcodec.Amino)

}
