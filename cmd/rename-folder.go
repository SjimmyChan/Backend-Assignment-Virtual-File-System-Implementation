/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// renameFolderCmd represents the renameFolder command
var renameFolderCmd = &cobra.Command{
	Use:   "rename-folder [username] [foldername] [new-folder-name]",
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

		new_foldername, _ := cmd.Flags().GetString("new-folder-name")
		if err := checkValidation(1, new_foldername, 30); err != nil {
			fmt.Println(err)
			return
		}

		renameFolder(username, foldername, new_foldername)

		fmt.Println("renameFolder called")
	},
}

func init() {

	renameFolderCmd.Flags().StringP("username", "u", "", "username")
	if err := renameFolderCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	renameFolderCmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := renameFolderCmd.MarkFlagRequired("foldername"); err != nil {
		fmt.Println(err)
	}

	renameFolderCmd.Flags().StringP("new-folder-name", "n", "", "new-folder-name")
	if err := renameFolderCmd.MarkFlagRequired("new-folder-name"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(renameFolderCmd)
}

func renameFolder(username string, foldername string, new_foldername string) {
	fmt.Println(username, foldername, new_foldername)
}