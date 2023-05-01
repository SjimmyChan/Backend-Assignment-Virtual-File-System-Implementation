/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (

	"github.com/spf13/cobra"
)

// deleteFolderCmd represents the deleteFolder command
var deleteFolderCmd = &cobra.Command{
	Use:   "delete-folder [username] [foldername]",
	Short: "",
	Long: ``,
	Run: DeleteFolderCmdRunE,
}

func DeleteFolderCmdRunE(cmd *cobra.Command, args []string) {
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

	if succeed := DeleteFolder(cmd, username, foldername); succeed {
		cmd.Println("Delete folder:" + foldername + " successfully.")	
	}

}

func DeleteFolderCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("username", "u", "", "username")
	if err := cmd.MarkFlagRequired("username"); err != nil {
		cmd.Print(err.Error())
	}

	cmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := cmd.MarkFlagRequired("foldername"); err != nil {
		cmd.Print(err.Error())
	}
}

func init() {

	DeleteFolderCmdFlags(deleteFolderCmd)

	rootCmd.AddCommand(deleteFolderCmd)
}

func DeleteFolder(cmd *cobra.Command, username string, foldername string)(succeed bool) {
	
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

	copy((*folders)[folder_index:], (*folders)[folder_index+1:])
	(*folders)[len(*folders)-1] = Folder{}
	*folders = (*folders)[: len(*folders)-1]

	if err := SaveUsersInformation(users); err != nil {
		cmd.Print(err.Error())
		return false
	}
	return true
}