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
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/pavania1/cosmos-leave-management/x/lms/types"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
// var queryCmd = &cobra.Command{
// 	Use:   "query",
// 	Short: "A brief description of your command",
// 	Long: `A longer description that spans multiple lines and likely contains examples
// and usage of using your command. For example:

// Cobra is a CLI library for Go that empowers applications.
// This application is a tool to generate the needed files
// to quickly create a Cobra application.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("query called")
// 	},
// }

func GetQueryCmd() *cobra.Command {
	studentqueryCmd := &cobra.Command{
		Use:   types.ModuleName,
		Short: "This is about the Get the queries",
		Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command.`,
		RunE: client.ValidateCmd,
	}
	studentqueryCmd.AddCommand(
		GetStudentCmd(),
		GetRegisterAdminCmd(),
		GetStudentscmd(),
		GetLeaveCmd(),
		GetAcceptLeaveCmd(),
	)
	return studentqueryCmd
}

// ------------------------->> GET STUDENT <<---------------------------------

func GetStudentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "GetStudent",
		Short: "This command is used for GetStudent",
		Long:  `Example:./simd query leavemanagementsystem GetStudent cosmos1jwjx39qetc5t865lhk43fxps25778ruzxffzy2 123.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			student := types.GetstudentRequest{
				//Id:      args[0],
				Address: args[0],
			}

			//queryClient := types.NewGetStudentRequest(Id, Address)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.GetStudent(cmd.Context(), &student)
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)

		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// ----------------------------->> GET STUDENTS <<-----------------------------------
func GetStudentscmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Getstudents",
		Short: "This is about the GetAllStudents ",
		Long:  `Example:./simd query leavemanagementsystem Getstudents.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetStudents(cmd.Context(), &types.GetstudentsRequest{})
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// ------------------------------>> GET REGISTER ADMIN <<-------------------------------------
func GetRegisterAdminCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "GetRegisterAdmin",
		Short: "This is about the Get admin by Address",
		Long:  `Example:./simd query leavemanagementsystem GetRegisterAdmin cosmos1wlqey0t8d95u5v9xegyhvnnup2vcdq82fmx2gy`,

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)

			if err != nil {
				panic(err)
			}
			Admin := types.GetRegisterAdminRequest{
				Address: args[0],
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.GetAdmin(cmd.Context(), &Admin)
			if err != nil {

				panic(err)
			}

			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// ------------------------------->> GET ACCEPT lEAVE <<------------------------------
func GetAcceptLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "GetAcceptLeave",
		Short: "This is about the Get AcceptLeaves",
		Long:  `Example:./simd query leavemanagementsystem GetAcceptLeave`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetAcceptLeave(cmd.Context(), &types.GetLeaveApproveRequest{})
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}

// ---------------------------->> GET  APPLY LEAVE <<-------------------------------
func GetLeaveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "GetLeave",
		Short: "This command is used for Getleaves leave",
		Long:  `Example:./simd query leavemanagementsystem GetLeave.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				panic(err)
			}
			panic("called 1")
			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.GetLeave(cmd.Context(), &types.GetLeaveRequest{})
			if err != nil {
				panic(err)
			}
			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}

func init() {
	//rootCmd.AddCommand(queryCmd)
	rootCmd.AddCommand(GetStudentCmd())
	rootCmd.AddCommand(GetRegisterAdminCmd())
	rootCmd.AddCommand(GetStudentscmd())
	rootCmd.AddCommand(GetLeaveCmd())
	rootCmd.AddCommand(GetAcceptLeaveCmd())

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
