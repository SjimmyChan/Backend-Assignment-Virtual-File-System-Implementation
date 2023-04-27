/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteFolderCmd represents the deleteFolder command
var deleteFolderCmd = &cobra.Command{
	Use:   "delete-folder [username] [foldername]",
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

		if succeed := deleteFolder(username, foldername); succeed {
			fmt.Println("Delete folder:" + foldername + " successfully.")	
		}

	},
}

func init() {

	deleteFolderCmd.Flags().StringP("username", "u", "", "username")
	if err := deleteFolderCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	deleteFolderCmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := deleteFolderCmd.MarkFlagRequired("foldername"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(deleteFolderCmd)
}

func deleteFolder(username string, foldername string)(succeed bool) {
	
	users := getUsersInformation()
	user_exist, user_index := checkUserExist(users, username)
	if !user_exist {
		fmt.Println("Error: The username:" + username + " doesn't exist.")
		return false
	}

	folders := &users[user_index].Folders
	folder_exist, folder_index := checkFolderExist(folders, foldername)
	if !folder_exist {
		fmt.Println("Error: The foldername:" + foldername + " doesn't existed.")
		return false
	}

	copy((*folders)[folder_index:], (*folders)[folder_index+1:])
	(*folders)[len(*folders)-1] = Folder{}
	*folders = (*folders)[: len(*folders)-1]

	if err := saveUsersInformation(users); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}