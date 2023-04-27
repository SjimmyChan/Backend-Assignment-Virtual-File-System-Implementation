/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createFolderCmd represents the createFolder command
var createFolderCmd = &cobra.Command{
	Use:   "create-folder [username] [foldername] [description]?",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		if err := checkValidation(username, 30); err != nil {
			fmt.Println(err)
			return
		}

		foldername, _ := cmd.Flags().GetString("foldername")
		if err := checkValidation(foldername, 30); err != nil {
			fmt.Println(err)
			return
		}

		description, _ := cmd.Flags().GetString("description")

		createFolder(username, foldername, description)

		fmt.Println("createFolder called")
	},
}

func init() {

	createFolderCmd.Flags().StringP("username", "u", "", "username")
	if err := createFolderCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	createFolderCmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := createFolderCmd.MarkFlagRequired("foldername"); err != nil {
		fmt.Println(err)
	}

	createFolderCmd.Flags().StringP("description", "d", "", "description")

	rootCmd.AddCommand(createFolderCmd)
}

func createFolder(username string, foldername string, description string) {
	fmt.Println(username, foldername, description)
}
