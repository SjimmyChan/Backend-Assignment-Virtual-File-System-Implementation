/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
)

// createFileCmd represents the createFile command
var createFileCmd = &cobra.Command{
	Use:   "create-file [username] [foldername] [filename] [description]?",
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

		description, _ := cmd.Flags().GetString("description")

		if succeed := createFile(username, foldername, filename, description); succeed {
			fmt.Println("Create file:" + filename + " successfully.")
		}
	},
}


func init() {

	// create flags
	createFileCmd.Flags().StringP("username", "u", "", "username")
	if err := createFileCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	createFileCmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := createFileCmd.MarkFlagRequired("foldername"); err != nil {
		fmt.Println(err)
	}

	createFileCmd.Flags().StringP("filename", "i", "", "filename")
	if err := createFileCmd.MarkFlagRequired("filename"); err != nil {
		fmt.Println(err)
	}

	createFileCmd.Flags().StringP("description", "d", "", "description")

	rootCmd.AddCommand(createFileCmd)
}

func createFile(username string, foldername string, filename string, description string)(succeed bool) {
	
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

	files := &(*folders)[folder_index].Files
	file_exist, _ := checkFileExist(files, filename)
	if file_exist {
		fmt.Println("Error: The filename:" + filename + " has already existed.")
		return false
	}

	current_time := time.Now()
	file := File{
		Filename: filename, 
		Description: description, 
		Created_at: current_time.Format("01-02-2006 15:04:05"), 
	}

	*files = append(*files, file)
	
	if err := saveUsersInformation(users); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
