/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"sort"
	"github.com/spf13/cobra"
)

// listFoldersCmd represents the listFolders command
var listFoldersCmd = &cobra.Command{
	Use:   "list-folders [username] [--sorted-name|--sorted-created] [asc|desc]",
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		if err := checkValidation(0, username, 30); err != nil {
			fmt.Println(err)
			return
		}

		sorted_name, _ := cmd.Flags().GetString("sorted-name")
		sorted_created, _ := cmd.Flags().GetString("sorted-created")

		if sorted_name != "" && sorted_created != "" {
			fmt.Println("Warning: Please only choose one of sorting factor.")
			return
		}

		listFolders(username, sorted_name, sorted_created)
	},
}

func init() {

	listFoldersCmd.Flags().StringP("username", "u", "", "username")
	if err := listFoldersCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	listFoldersCmd.Flags().StringP("sorted-name", "", "", "sorted by username")

	listFoldersCmd.Flags().StringP("sorted-created", "", "", "sorted by create time")

	rootCmd.AddCommand(listFoldersCmd)
}

var folders []Folder

func listFolders(username string, sorted_name string, sorted_created string) {

	users := getUsersInformation()
	user_exist, user_index := checkUserExist(users, username)
	if !user_exist {
		fmt.Println("Error: The username:" + username + " doesn't exist.")
		return
	}

	folders := users[user_index].Folders
	if len(folders) == 0 {
		fmt.Println("Warning: This " + username + "doesn't have any folders.")
		return
	}

	if sorted_name != "" {
		sort.Sort(folderNameList(folders))
		if sorted_name == "asc" {
			fmt.Println("[foldername] | [description] | [created at] | [username]")
			for index := 0; index < len(folders); index++ {
				fmt.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at + " | " + username)
			}
		} else if sorted_name == "desc" {
			fmt.Println("[foldername] | [description] | [created at] | [username]")
			for index := len(folders) - 1; index >= 0; index-- {
				fmt.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at + " | " + username)
			}
		} else {
			fmt.Println("Error: Please use asc/desc as sorting method")
		}
	} else if sorted_created != "" {
		sort.Sort(folderCreatedList(folders))
		if sorted_created == "asc" {
			fmt.Println("[foldername] | [description] | [created at] | [username]")
			for index := 0; index < len(folders); index++ {
				fmt.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at + " | " + username)
			}
		} else if sorted_created == "desc" {
			fmt.Println("[foldername] | [description] | [created at] | [username]")
			for index := len(folders) - 1; index >= 0; index-- {
				fmt.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at + " | " + username)
			}
		} else {
			fmt.Println("Error: Please use asc/desc as sorting method")
		}
	} else {
		sort.Sort(folderNameList(folders))
		fmt.Println("[foldername] | [description] | [created at] | [username]")
		for index := 0; index < len(folders); index++ {
			fmt.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at + " | " + username)
		}
	}
}

type folderNameList []Folder

func (f folderNameList) Len() int {
	return len(f)
}

func (f folderNameList) Less(i, j int) bool {
	return f[i].Foldername < f[j].Foldername
}

func (f folderNameList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

type folderCreatedList []Folder

func (f folderCreatedList) Len() int {
	return len(f)
}

func (f folderCreatedList) Less(i, j int) bool {
	return f[i].Created_at < f[j].Created_at
}

func (f folderCreatedList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}