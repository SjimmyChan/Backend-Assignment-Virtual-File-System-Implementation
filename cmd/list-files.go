/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"sort"
	
	"github.com/spf13/cobra"
)

// listFilesCmd represents the listFiles command
var listFilesCmd = &cobra.Command{
	Use:   "list-files [username] [foldername] [--sorted-name|--sorted-created] [asc|desc]",
	Short: "",
	Long: ``,
	Run: ListFileCmdRunE,
}

func ListFileCmdRunE(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	if err := CheckValidation(0, username, 30); err != nil {
		cmd.Println(err.Error())
		return
	}

	foldername, _ := cmd.Flags().GetString("foldername")
	if err := CheckValidation(1, foldername, 30); err != nil {
		cmd.Println(err.Error())
		return
	}

	sorted_name, _ := cmd.Flags().GetString("sorted-name")
	sorted_created, _ := cmd.Flags().GetString("sorted-created")

	if sorted_name != "" && sorted_created != "" {
		cmd.Println("Warning: Please only choose one of sorting factor.")
		return
	}

	ListFiles(cmd, username, foldername, sorted_name, sorted_created)
}

func ListFileCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("username", "u", "", "username")
	if err := cmd.MarkFlagRequired("username"); err != nil {
		cmd.Println(err.Error())
	}

	cmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := cmd.MarkFlagRequired("foldername"); err != nil {
		cmd.Println(err.Error())
	}

	cmd.Flags().String("sorted-name", "", "sorted by username")

	cmd.Flags().String("sorted-created", "", "sorted by create time")
}

func init() {

	ListFileCmdFlags(listFilesCmd)

	rootCmd.AddCommand(listFilesCmd)
}

func ListFiles(cmd *cobra.Command, username string, foldername string, sorted_name string, sorted_created string) {
	users := GetUsersInformation()
	user_exist, user_index := checkUserExist(users, username)
	if !user_exist {
		cmd.Println("Error: The username:" + username + " doesn't exist.")
		return
	}

	folders := users[user_index].Folders
	folder_exist, folder_index := checkFolderExist(&folders, foldername)
	if !folder_exist {
		cmd.Println("Error: The foldername:" + foldername + " doesn't exist.")
		return
	}

	files := folders[folder_index].Files
	if len(files) == 0 {
		cmd.Println("Warning: This folder is empty.")
		return
	}

	if sorted_name != "" {
		sort.Sort(fileNameList(files))
		if sorted_name == "asc" {
			cmd.Println("[filename] | [description] | [created at] | [foldername] | [username]")
			for index := 0; index < len(files); index++ {
				cmd.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at.Format("01-02-2006 15:04:05") +  " | " + foldername + " | " + username)
			}
		} else if sorted_name == "desc" {
			cmd.Println("[filename] | [description] | [created at] | [foldername] | [username]")
			for index := len(files) - 1; index >= 0; index-- {
				cmd.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at.Format("01-02-2006 15:04:05") +  " | " + foldername + " | " + username)
			}
		} else {
			cmd.Println("Error: Please use asc/desc as sorting method.")
		}
	} else if sorted_created != "" {
		sort.Sort(fileCreatedList(files))
		if sorted_created == "asc" {
			cmd.Println("[filename] | [description] | [created at] | [foldername] | [username]")
			for index := 0; index < len(files); index++ {
				cmd.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at.Format("01-02-2006 15:04:05") + " | " + foldername + " | " + username)
			}
		} else if sorted_created == "desc" {
			cmd.Println("[filename] | [description] | [created at] | [foldername] | [username]")
			for index := len(files) - 1; index >= 0; index-- {
				cmd.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at.Format("01-02-2006 15:04:05") + " | " + foldername + " | " + username)
			}
		} else {
			cmd.Println("Error: Please use asc/desc as sorting method.")
		}
	} else {
		sort.Sort(fileNameList(files))
		cmd.Println("[filename] | [description] | [created at] | [foldername] | [username]")
		for index := 0; index < len(files); index++ {
			cmd.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at.Format("01-02-2006 15:04:05") + " | " + foldername + " | " + username)
		}
	}
}

type fileNameList []File

func (f fileNameList) Len() int {
	return len(f)
}

func (f fileNameList) Less(i, j int) bool {
	return f[i].Filename < f[j].Filename
}

func (f fileNameList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type fileCreatedList []File

func (f fileCreatedList) Len() int {
	return len(f)
}

func (f fileCreatedList) Less(i, j int) bool {
	return f[i].Created_at.Before(f[j].Created_at)
}

func (f fileCreatedList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
