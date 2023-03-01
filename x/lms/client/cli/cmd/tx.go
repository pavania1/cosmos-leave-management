/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pavania1/cosmos-leave-management/x/lms/types"
	"github.com/spf13/cobra"
)

func GetTxCmd() *cobra.Command {
	studentTxCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "|lms|",
		Long:  "lms module commands",
		RunE:  client.ValidateCmd,
	}
	studentTxCmd.AddCommand(
		RegisterAdminCmd(),
		AddStudentCmd(),
		AcceptLeaveCmd(),
		ApplyLeaveCmd(),
	)
	return studentTxCmd
}
func AddStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Addstudent",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examplesand usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			admin := args[0]
			//address := args[1]
			address, _ := sdk.AccAddressFromBech32(args[1])
			name := args[2]
			id := args[3]

			msgClient := types.NewAddStudentRequest(admin, address.String(), name, id)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	return cmd
}
func RegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "RegisterAdmin",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examplesand usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			fmt.Println("args[0]", args[0])
			fmt.Println("args[1]", args[1])
			Address, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				panic(err)
			}
			fmt.Println("address", Address)
			name := args[1]
			msgClient := types.NewRegisterAdminRequest(Address, name)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func AcceptLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "AcceptLeave",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examplesand usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			Admin := args[0]
			LeaveId := args[1]
			leaveId, err := strconv.ParseUint(LeaveId, 10, 64)
			if err != nil {
				return err
			}
			msgClient := types.NewAcceptLeaveRequest(Admin, leaveId, types.LeaveStatus_STATUS_ACCEPTED)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
func ApplyLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ApplyLeave",
		Short: "A brief description of your command",
		Long:  `A longer description that spans multiple lines and likely contains examplesand usage of using your command.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				panic(err)
			}
			dateString := "2023-02-22"
			fromdate, _ := time.Parse("2023-02-22", dateString)
			todate, _ := time.Parse("2023-02-22", "2023-02-24")
			//Address := args[0]
			Address, _ := sdk.AccAddressFromBech32(args[0])
			Reason := args[1]
			From := &fromdate
			To := &todate
			msgClient := types.NewApplyLeaveRequest(Address, Reason, From, To)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msgClient)
		},
	}
	return cmd
}
func init() {
	//rootCmd.AddCommand(txCmd)
	rootCmd.AddCommand(AddStudentCmd())
	rootCmd.AddCommand(RegisterAdminCmd())
	rootCmd.AddCommand(AcceptLeaveCmd())
	rootCmd.AddCommand(ApplyLeaveCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// txCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// txCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
