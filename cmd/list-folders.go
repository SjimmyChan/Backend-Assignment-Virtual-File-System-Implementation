/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"sort"
	"github.com/spf13/cobra"
)

// listFoldersCmd represents the listFolders command
var listFoldersCmd = &cobra.Command{
	Use:   "list-folders [username] [--sorted-name|--sorted-created] [asc|desc]",
	Short: "",
	Long: ``,
	Run: ListFoldersCmdRunE,
}

func ListFoldersCmdRunE(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	if err := CheckValidation(0, username, 30); err != nil {
		cmd.Println(err.Error())
		return
	}

	sorted_name, _ := cmd.Flags().GetString("sort-name")
	sorted_created, _ := cmd.Flags().GetString("sort-created")

	if sorted_name != "" && sorted_created != "" {
		cmd.Println("Warning: Please only choose one of sorting factor.")
		return
	}

	ListFolders(cmd, username, sorted_name, sorted_created)
}

func ListFoldersCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("username", "u", "", "username")
	if err := cmd.MarkFlagRequired("username"); err != nil {
		cmd.Println(err.Error())
	}

	cmd.Flags().StringP("sort-name", "", "", "sorted by username")

	cmd.Flags().StringP("sort-created", "", "", "sorted by create time")
}

func init() {

	ListFoldersCmdFlags(listFoldersCmd)

	rootCmd.AddCommand(listFoldersCmd)
}

var folders []Folder

func ListFolders(cmd *cobra.Command, username string, sorted_name string, sorted_created string) {

	users := GetUsersInformation()
	user_exist, user_index := checkUserExist(users, username)
	if !user_exist {
		cmd.Println("Error: The username:" + username + " doesn't exist.")
		return
	}

	folders := users[user_index].Folders
	if len(folders) == 0 {
		cmd.Println("Warning: The " + username + " doesn't have any folders.")
		return
	}

	if sorted_name != "" {
		sort.Sort(folderNameList(folders))
		if sorted_name == "asc" {
			cmd.Println("[foldername] | [description] | [created at] | [username]")
			for index := 0; index < len(folders); index++ {
				cmd.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at.Format("01-02-2006 15:04:05") + " | " + username)
			}
		} else if sorted_name == "desc" {
			cmd.Println("[foldername] | [description] | [created at] | [username]")
			for index := len(folders) - 1; index >= 0; index-- {
				cmd.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at.Format("01-02-2006 15:04:05") + " | " + username)
			}
		} else {
			cmd.Println("Error: Please use asc/desc as sorting method.")
		}
	} else if sorted_created != "" {
		sort.Sort(folderCreatedList(folders))
		if sorted_created == "asc" {
			cmd.Println("[foldername] | [description] | [created at] | [username]")
			for index := 0; index < len(folders); index++ {
				cmd.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at.Format("01-02-2006 15:04:05") + " | " + username)
			}
		} else if sorted_created == "desc" {
			cmd.Println("[foldername] | [description] | [created at] | [username]")
			for index := len(folders) - 1; index >= 0; index-- {
				cmd.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at.Format("01-02-2006 15:04:05") + " | " + username)
			}
		} else {
			cmd.Println("Error: Please use asc/desc as sorting method.")
		}
	} else {
		sort.Sort(folderNameList(folders))
		cmd.Println("[foldername] | [description] | [created at] | [username]")
		for index := 0; index < len(folders); index++ {
			cmd.Println(folders[index].Foldername + " | " + folders[index].Description + " | " + folders[index].Created_at.Format("01-02-2006 15:04:05") + " | " + username)
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
	return f[i].Created_at.Before(f[j].Created_at)
}

func (f folderCreatedList) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}