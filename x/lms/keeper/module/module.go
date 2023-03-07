package module

import (
	"context"
	"encoding/json"

	// "cosmossdk.io/core/appmodule"

	abci "github.com/tendermint/tendermint/abci/types"

	// "cosmossdk.io/core/appmodule"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/nft"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	// "github.com/cosmos/cosmos-sdk/types/module"
	// "github.com/cosmos/cosmos-sdk/x/mint/types"

	// "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/pavania1/cosmos-leave-management/x/lms/types"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	// "github.com/pavania1/cosmos-leave-management/codec"
	//"github.com/pavania1/cosmos-leave-management/x/lms"
	"github.com/pavania1/cosmos-leave-management/x/lms/client/cli/cmd"
	"github.com/pavania1/cosmos-leave-management/x/lms/keeper"

	// "github.com/pavania1/cosmos-leave-managemnet/x/lms/client/cli/cmd"
	"github.com/spf13/cobra"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
	//_ module.AppModuleSimulation = AppModule{}
)

type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the lms module's name
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterServices registers a gRPC query service to respond to the
// module-specific gRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	// types.RegisterMsgServer()
	types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterLegacyAminoCodec registers the lms module's types for the given codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterInterfaces registers the lms module's interface types
func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// DefaultGenesis returns default genesis state as raw bytes for the nft
// module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}

// ValidateGenesis performs genesis state validation for the nft module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return sdkerrors.Wrapf(err, "failed to unmarshal %s genesis state", types.ModuleName)
	}

	return types.ValidateGenesis(&data)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the nft module.
func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *gwruntime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// GetQueryCmd returns the cli query commands for the nft module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cmd.GetQueryCmd()
}

// GetTxCmd returns the transaction commands for the nft module
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cmd.GetTxCmd()
}

// AppModule implements the sdk.AppModule interface
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
	//registry cdctypes.InterfaceRegistry
}

func NewAppModule(cdc codec.Codec, k keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         k,
		//registry:       registry,
	}
}

var _ module.AppModule = AppModule{}

func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// Name returns the lms module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

//RegisterInvariants does nothing, there are no invariants to enforce
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) NewHandler() sdk.Handler {
	return nil
}
func (AppModule) LegacyQuerierHandler(LegacyQuerier *codec.LegacyAmino) sdk.Querier {
	return nil
}
func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(nft.RouterKey, nil)
}

// InitGenesis performs genesis initialization for the nft module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	// var genesisState nft.GenesisState
	// cdc.MustUnmarshalJSON(data, &genesisState)
	// am.keeper.InitGenesis(ctx, &genesisState)
	// return []abci.ValidatorUpdate{}
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the nft
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	//gs := am.keeper.ExportGenesis(ctx)
	//return cdc.MustMarshalJSON(gs)
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 1 }
