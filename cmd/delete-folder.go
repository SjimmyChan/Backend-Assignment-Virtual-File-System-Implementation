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
	Short: "",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		if err := checkValidation(username, 30); err != nil {
			fmt.Println(err)
			return
		}

		foldername, _ := cmd.Flags().GetString("foldername")
		if err := checkValidation(foldername, 30); err != nil {
			fmt.Println(err)
			return
		}

		deleteFolder(username, foldername)

		fmt.Println("deleteFolder called")
	},
}

func init() {

	deleteFolderCmd.Flags().StringP("username", "u", "", "username")
	if err := deleteFolderCmd.MarkFlagRequired("username"); err != nil {
		fmt.Println(err)
	}

	deleteFolderCmd.Flags().StringP("foldername", "f", "", "foldername")
	if err := deleteFolderCmd.MarkFlagRequired("foldername"); err != nil {
		fmt.Println(err)
	}

	rootCmd.AddCommand(deleteFolderCmd)
}

func deleteFolder(username string, foldername string) {
	fmt.Println(username, foldername)
}