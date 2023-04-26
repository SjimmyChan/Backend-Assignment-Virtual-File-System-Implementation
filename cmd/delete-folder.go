/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteFolderCmd represents the deleteFolder command
var deleteFolderCmd = &cobra.Command{
	Use:   "delete-folder [username] [foldername]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteFolder called")
	},
}

func init() {

	var username string
	var foldername string

	deleteFolderCmd.Flags().StringVarP(&username, "username", "u", "", "username")
	deleteFolderCmd.Flags().StringVarP(&foldername, "foldername", "f", "", "foldername")
	
	deleteFolderCmd.MarkFlagsRequiredTogether("username", "foldername")

	rootCmd.AddCommand(deleteFolderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteFolderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteFolderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
