/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

// createFolderCmd represents the createFolder command
var createFolderCmd = &cobra.Command{
	Use:   "create-folder [username] [foldername] [description]?",
	Short: "",
	Long: ``,
	Run: CreateFolderCmdRunE,
}

func CreateFolderCmdRunE(cmd *cobra.Command, args []string) {
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

	description, _ := cmd.Flags().GetString("description")

	if succeed := CreateFolder(cmd, username, foldername, description); succeed {
		cmd.Println("Create " + foldername + " in " + username + " successfully.")
	}
}

func CreateFolderCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("username", "u", "", "username")
	if err := cmd.MarkFlagRequired("username"); err != nil {
		cmd.Print(err.Error())
	}

	cmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := cmd.MarkFlagRequired("foldername"); err != nil {
		cmd.Print(err.Error())
	}

	cmd.Flags().StringP("description", "d", "", "description")
}

func init() {

	CreateFolderCmdFlags(createFolderCmd)

	rootCmd.AddCommand(createFolderCmd)
}

func CreateFolder(cmd *cobra.Command, username string, foldername string, description string) (succeed bool) {

	users := GetUsersInformation()
	user_exist, user_index := checkUserExist(users, username)
	if !user_exist {
		cmd.Println("Error: The username:" + username + " doesn't exist.")
		return false
	}

	folders := &users[user_index].Folders
	folder_exist, _ := checkFolderExist(folders, foldername)
	if folder_exist {
		cmd.Println("Error: The foldername:" + foldername + " has already existed.")
		return false
	}
	
	folder := Folder{
		Foldername: foldername, 
		Description: description, 
		Created_at: time.Now(), 
		Files: []File{},
	}
	*folders = append(*folders, folder)
	
	if err := SaveUsersInformation(users); err != nil {
		cmd.Print(err.Error())
		return false
	}
	return true
}
