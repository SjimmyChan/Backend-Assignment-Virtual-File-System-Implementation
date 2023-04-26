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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listFolders called")
	},
}

func init() {

	var username string

	listFoldersCmd.Flags().StringVarP(&username, "username", "u", "", "username")
	listFoldersCmd.MarkFlagRequired("username")

	rootCmd.AddCommand(listFoldersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFoldersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFoldersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
