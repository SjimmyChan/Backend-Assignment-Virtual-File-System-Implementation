/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
)

// createFolderCmd represents the createFolder command
var createFolderCmd = &cobra.Command{
	Use:   "create-folder [username] [foldername] [description]?",
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

		description, _ := cmd.Flags().GetString("description")

		if succeed := createFolder(username, foldername, description); succeed {
			fmt.Println("Create folder:" + foldername + " successfully.")
		}
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

func createFolder(username string, foldername string, description string) (succeed bool) {

	users := getUsersInformation()
	user_exist, user_index := checkUserExist(users, username)
	if !user_exist {
		fmt.Println("Error: The username:" + username + " doesn't exist.")
		return false
	}

	folders := &users[user_index].Folders
	folder_exist, _ := checkFolderExist(folders, foldername)
	if folder_exist {
		fmt.Println("Error: The foldername:" + foldername + " has already existed.")
		return false
	}
	
	current_time := time.Now()
	folder := Folder{
		Foldername: foldername, 
		Description: description, 
		Created_at: current_time.Format("01-02-2006 15:04:05"), 
		Files: []File{},
	}
	*folders = append(*folders, folder)
	
	if err := saveUsersInformation(users); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
