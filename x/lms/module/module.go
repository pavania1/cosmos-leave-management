package module

import (
	"context"
	"encoding/json"
	"github.com/pavania1/cosmos-leave-management/codec"
)
var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)
type AppModuleBasic struct {
	cdc codec.Codec
}
