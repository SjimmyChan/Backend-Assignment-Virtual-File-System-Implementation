/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

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

		listFolders(username)

		fmt.Println("listFolders called")
	},
}

func init() {

	listFoldersCmd.Flags().StringP("username", "u", "", "username")
	if err := listFoldersCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(listFoldersCmd)
}

func listFolders(username string) {
	fmt.Println(username)
}
