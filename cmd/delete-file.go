/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// deleteFileCmd represents the deleteFile command
var deleteFileCmd = &cobra.Command{
	Use:   "delete-file [username] [foldername] [filename]",
	Short: "",
	Long: ``,
	Run: DeleteFileCmdRunE,
}

func DeleteFileCmdRunE(cmd *cobra.Command, args []string) {
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

	if succeed := DeleteFile(cmd, username, foldername, filename); succeed {
		cmd.Println("Delete file:" + filename + " successfully.")
	}
}

func DeleteFileCmdFlags(cmd *cobra.Command) {
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
}

func init() {

	DeleteFileCmdFlags(deleteFileCmd)

	rootCmd.AddCommand(deleteFileCmd)
}

func DeleteFile(cmd *cobra.Command, username string, foldername string, filename string)(succeed bool) {
	
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
	file_exist, file_index := checkFileExist(files, filename)
	if !file_exist {
		cmd.Println("Error: The filename:" + filename + " doesn't exist.")
		return false
	}

	copy((*files)[file_index:], (*files)[file_index+1:])
	(*files)[len(*files)-1] = File{}
	*files = (*files)[: len(*files)-1]

	if err := SaveUsersInformation(users); err != nil {
		cmd.Print(err.Error())
		return false
	}
	return true
}
