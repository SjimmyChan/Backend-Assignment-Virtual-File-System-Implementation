/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listFilesCmd represents the listFiles command
var listFilesCmd = &cobra.Command{
	Use:   "list-files [username] [foldername] [--sorted-name|--sorted-created] [asc|desc]",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		if err := checkValidation(0, username, 30); err != nil {
			fmt.Println(err)
			return
		}

		foldername, _ := cmd.Flags().GetString("foldername")
		if err := checkValidation(1, foldername, 30); err != nil {
			fmt.Println(err)
			return
		}

		listFiles(username, foldername)

		fmt.Println("listFiles called")
	},
}

func init() {


	listFilesCmd.Flags().StringP("username", "u", "", "username")
	if err := listFilesCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	listFilesCmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := listFilesCmd.MarkFlagRequired("foldername"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(listFilesCmd)
}

func listFiles(username string, foldername string) {
	fmt.Println(username, foldername)
}
