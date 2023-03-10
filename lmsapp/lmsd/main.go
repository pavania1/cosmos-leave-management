package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	// "github.com/cosmos/cosmos-sdk/simapp"
	// "lms-cosmos/lmsapp"
	// "github.com/Leave-Management-System/lms-cosmos/lmsapp"
	"github.com/pavania1/cosmos-leave-management/lmsapp"
	"github.com/pavania1/cosmos-leave-management/lmsapp/lmsd/cmd"
	// "github.com/pavania1/cosmos-leave-management/lmsapp/simd/cmd"
	// "github.com/cosmos/cosmos-sdk/simapp/simd/cmd"
	// "lms-cosmos/lmsapp/simd/cmd"
	// "github.com/Leave-Management-System/lms-cosmos/lmsapp/simd/cmd"
	// "github.com/pavania1/cosmos-leave-management/lmsapp/simd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", lmsapp.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
