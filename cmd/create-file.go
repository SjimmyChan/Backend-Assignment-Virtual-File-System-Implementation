/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"time"
	"github.com/spf13/cobra"
)

// createFileCmd represents the createFile command
var createFileCmd = &cobra.Command{
	Use:   "create-file [username] [foldername] [filename] [description]?",
	Short: "",
	Long: ``,
	Run: CreateFileCmdRunE,
}

func CreateFileCmdRunE(cmd *cobra.Command, args []string) {
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

	filename, _ := cmd.Flags().GetString("filename")
	if err := CheckValidation(2, filename, 30); err != nil {
		cmd.Print(err.Error())
		return
	}

	description, _ := cmd.Flags().GetString("description")

	if succeed := CreateFile(cmd, username, foldername, filename, description); succeed {
		cmd.Println("Create file:" + filename + " in " + username + "/" + foldername + " successfully.")
	}
}

func CreateFileCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("username", "u", "", "username")
	if err := cmd.MarkFlagRequired("username"); err != nil {
		cmd.Print(err.Error())
	}

	cmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := cmd.MarkFlagRequired("foldername"); err != nil {
		cmd.Print(err.Error())
	}

	cmd.Flags().StringP("filename", "i", "", "filename")
	if err := cmd.MarkFlagRequired("filename"); err != nil {
		cmd.Print(err.Error())
	}

	cmd.Flags().StringP("description", "d", "", "description")
}


func init() {

	CreateFileCmdFlags(createFileCmd)

	rootCmd.AddCommand(createFileCmd)
}

func CreateFile(cmd *cobra.Command, username string, foldername string, filename string, description string)(succeed bool) {
	
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

	files := &(*folders)[folder_index].Files
	file_exist, _ := checkFileExist(files, filename)
	if file_exist {
		cmd.Println("Error: The filename:" + filename + " has already existed.")
		return false
	}

	file := File{
		Filename: filename, 
		Description: description, 
		Created_at: time.Now(), 
	}

	*files = append(*files, file)
	
	if err := SaveUsersInformation(users); err != nil {
		cmd.Println(err)
		return false
	}
	return true
}
