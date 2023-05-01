/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (

	"github.com/spf13/cobra"
)

// renameFolderCmd represents the renameFolder command
var renameFolderCmd = &cobra.Command{
	Use:   "rename-folder [username] [foldername] [new-folder-name]",
	Short: "",
	Long: ``,
	Run: RenameFolderCmdRunE,
}

func RenameFolderCmdRunE(cmd *cobra.Command, args []string) {

	username, _ := cmd.Flags().GetString("username")
	if err := CheckValidation(0, username, 30); err != nil {
		cmd.Print(err.Error())
		return
	}

	foldername, _ := cmd.Flags().GetString("foldername")
	if err := CheckValidation(1, foldername, 30); err != nil {
		cmd.Print(err.Error())
		return
	}

	new_foldername, _ := cmd.Flags().GetString("new_foldername")
	if err := CheckValidation(3, new_foldername, 30); err != nil {
		cmd.Print(err.Error())
		return
	}

	if succeed := RenameFolder(cmd, username, foldername, new_foldername); succeed {
		cmd.Println("Rename foldername from " + foldername + " to " + new_foldername + " successfully.")
	}
}

func RenameFolderCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("username", "u", "", "username")
	if err := cmd.MarkFlagRequired("username"); err != nil {
		cmd.Print(err.Error())
	}

	cmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := cmd.MarkFlagRequired("foldername"); err != nil {
		cmd.Print(err.Error())
	}

	cmd.Flags().StringP("new_foldername", "n", "", "new_foldername")
	if err := cmd.MarkFlagRequired("new_foldername"); err != nil {
		cmd.Print(err.Error())
	}
}

func init() {

	RenameFolderCmdFlags(renameFolderCmd)

	rootCmd.AddCommand(renameFolderCmd)
}

func RenameFolder(cmd *cobra.Command, username string, foldername string, new_foldername string)(succeed bool) {
	
	users := GetUsersInformation()
	user_exist, user_index := checkUserExist(users, username)
	if !user_exist {
		cmd.Println("Error: The username:" + username + " doesn't exist.")
		return false
	}

	folders := &users[user_index].Folders
	folder_exist, folder_index := checkFolderExist(folders, foldername)
	if !folder_exist {
		cmd.Println("Error: The foldername:" + foldername + " doesn't exist.")
		return false
	}
	
	if foldername == new_foldername {
		cmd.Println("Warning: The foldername is already called " + foldername + ".")
		return false
	}

	new_folder_exist, _ := checkFolderExist(folders, new_foldername)
	if new_folder_exist {
		cmd.Println("Error: Cannot change foldername to " + new_foldername + ", since it has already existed.")
		return false
	}

	(*folders)[folder_index].Foldername = new_foldername

	if err := SaveUsersInformation(users); err != nil {
		cmd.Print(err.Error())
		return false
	}
	return true

}