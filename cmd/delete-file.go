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

		deleteFile(username, foldername, filename)

		fmt.Println("deleteFile called")
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

func deleteFile(username string, foldername string, filename string) {
	fmt.Println(username, foldername, filename)
}
