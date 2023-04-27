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

		new_foldername, _ := cmd.Flags().GetString("new_foldername")
		if err := checkValidation(1, new_foldername, 30); err != nil {
			fmt.Println(err)
			return
		}

		if succeed := renameFolder(username, foldername, new_foldername); succeed {
			fmt.Println("Rename foldername from " + foldername + " to " + new_foldername + " successfully.")
		}
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

	renameFolderCmd.Flags().StringP("new_foldername", "n", "", "new_foldername")
	if err := renameFolderCmd.MarkFlagRequired("new_foldername"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(renameFolderCmd)
}

func renameFolder(username string, foldername string, new_foldername string)(succeed bool) {
	
	users := getUsersInformation()
	user_exist, user_index := checkUserExist(users, username)
	if !user_exist {
		fmt.Println("Error: The username:" + username + " doesn't exist.")
		return false
	}

	folders := &users[user_index].Folders
	folder_exist, folder_index := checkFolderExist(folders, foldername)
	if !folder_exist {
		fmt.Println("Error: The foldername:" + foldername + " doesn't exist.")
		return false
	}
	
	if foldername == new_foldername {
		fmt.Println("Warning: The foldername is already called " + foldername + ".")
		return false
	}

	new_folder_exist, _ := checkFolderExist(folders, new_foldername)
	if new_folder_exist {
		fmt.Println("Error: Cannot change foldername to " + new_foldername + ", since it has already existed.")
		return false
	}

	(*folders)[folder_index].Foldername = new_foldername

	if err := saveUsersInformation(users); err != nil {
		fmt.Println(err)
		return false
	}
	return true

}