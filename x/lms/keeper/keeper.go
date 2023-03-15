package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
}

// func NewKeeper(key storetypes.Storekey, cdc codec.BinaryCodec) keeper {
// 	if addr := sdk.GetModuleAddress(); addr == nil {
// 		panic("the module address has not been set")
// 	}
// 	return Keeper{
// 		cdc:      cdc,
// 		storekey: storekey,
// 	}
// }
