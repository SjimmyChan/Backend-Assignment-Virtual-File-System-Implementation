/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"sort"
	"github.com/spf13/cobra"
)

// listFilesCmd represents the listFiles command
var listFilesCmd = &cobra.Command{
	Use:   "list-files [username] [foldername] [--sorted-name|--sorted-created] [asc|desc]",
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

		sorted_name, _ := cmd.Flags().GetString("sorted-name")
		sorted_created, _ := cmd.Flags().GetString("sorted-created")

		if sorted_name != "" && sorted_created != "" {
			fmt.Println("Warning: Please only choose one of sorting factor.")
			return
		}

		listFiles(username, foldername, sorted_name, sorted_created)
	},
}

func init() {


	listFilesCmd.Flags().StringP("username", "u", "", "username")
	if err := listFilesCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	listFilesCmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := listFilesCmd.MarkFlagRequired("foldername"); err != nil {
		fmt.Println(err)
	}

	listFilesCmd.Flags().String("sorted-name", "", "sorted by username")

	listFilesCmd.Flags().String("sorted-created", "", "sorted by create time")

	rootCmd.AddCommand(listFilesCmd)
}

var files []File

func listFiles(username string, foldername string, sorted_name string, sorted_created string) {
	users := getUsersInformation()
	user_exist, user_index := checkUserExist(users, username)
	if !user_exist {
		fmt.Println("Error: The username:" + username + " doesn't exist.")
		return
	}

	folders := users[user_index].Folders
	folder_exist, folder_index := checkFolderExist(&folders, foldername)
	if !folder_exist {
		fmt.Println("Error: The foldername:" + foldername + " doesn't exist.")
		return
	}

	files := folders[folder_index].Files
	if len(files) == 0 {
		fmt.Println("Warning: This " + foldername + " doesn't have any files.")
		return
	}

	if sorted_name != "" {
		sort.Sort(fileNameList(files))
		if sorted_name == "asc" {
			fmt.Println("[filename] | [description] | [created at] | [foldername] | [username]")
			for index := 0; index < len(files); index++ {
				fmt.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at +  " | " + foldername + " | " + username)
			}
		} else if sorted_name == "desc" {
			fmt.Println("[filename] | [description] | [created at] | [foldername] | [username]")
			for index := len(files) - 1; index >= 0; index-- {
				fmt.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at +  " | " + foldername + " | " + username)
			}
		} else {
			fmt.Println("Error: Please use asc/desc as sorting method")
		}
	} else if sorted_created != "" {
		sort.Sort(fileCreatedList(files))
		if sorted_created == "asc" {
			fmt.Println("[filename] | [description] | [created at] | [foldername] | [username]")
			for index := 0; index < len(files); index++ {
				fmt.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at + " | " + foldername + " | " + username)
			}
		} else if sorted_created == "desc" {
			fmt.Println("[filename] | [description] | [created at] | [foldername] | [username]")
			for index := len(files) - 1; index >= 0; index-- {
				fmt.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at + " | " + foldername + " | " + username)
			}
		} else {
			fmt.Println("Error: Please use asc/desc as sorting method")
		}
	} else {
		sort.Sort(fileNameList(files))
		fmt.Println("[filename] | [description] | [created at] | [foldername] | [username]")
		for index := 0; index < len(files); index++ {
			fmt.Println(files[index].Filename + " | " + files[index].Description + " | " + files[index].Created_at + " | " + foldername + " | " + username)
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
	return f[i].Created_at < f[j].Created_at
}

func (f fileCreatedList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
