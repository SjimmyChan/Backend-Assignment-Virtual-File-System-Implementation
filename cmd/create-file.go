/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// createFileCmd represents the createFile command
var createFileCmd = &cobra.Command{
	Use:   "create-file [username] [foldername] [filename] [description]?",
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

		filename, _ := cmd.Flags().GetString("filename")
		if err := checkValidation(filename, 30); err != nil {
			fmt.Println(err)
			return
		}

		description, _ := cmd.Flags().GetString("description")

		createFile(username, foldername, filename, description)

		fmt.Println("createFile called")
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createFile(username string, foldername string, filename string, description string) {
	fmt.Println(username, foldername, filename, description)
}
