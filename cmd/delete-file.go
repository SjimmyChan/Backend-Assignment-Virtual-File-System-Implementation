/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// deleteFileCmd represents the deleteFile command
var deleteFileCmd = &cobra.Command{
	Use:   "delete-file [username] [foldername] [filename]",
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

		filename, _ := cmd.Flags().GetString("filename")
		if err := checkValidation(2, filename, 30); err != nil {
			fmt.Println(err)
			return
		}

		if succeed := deleteFile(username, foldername, filename); succeed {
			fmt.Println("Delete file:" + filename + " successfully.")
		}
	},
}

func init() {

	deleteFileCmd.Flags().StringP("username", "u", "", "username")
	if err := deleteFileCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	deleteFileCmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := deleteFileCmd.MarkFlagRequired("foldername"); err != nil {
		fmt.Println(err)
	}

	deleteFileCmd.Flags().StringP("filename", "i", "", "filename")
	if err := deleteFileCmd.MarkFlagRequired("filename"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(deleteFileCmd)
}

func deleteFile(username string, foldername string, filename string)(succeed bool) {
	
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

	files := &(*folders)[folder_index].Files
	file_exist, file_index := checkFileExist(files, filename)
	if !file_exist {
		fmt.Println("Error: The filename:" + filename + " doesn't existed.")
		return false
	}

	copy((*files)[file_index:], (*files)[file_index+1:])
	(*files)[len(*files)-1] = File{}
	*files = (*files)[: len(*files)-1]

	if err := saveUsersInformation(users); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
