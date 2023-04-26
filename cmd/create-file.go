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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createFile called")
	},
}


func init() {

	var username string
	var foldername string
	var filename string
	var description string

	createFileCmd.Flags().StringVarP(&username, "username", "u", "", "username")
	createFileCmd.Flags().StringVarP(&foldername, "foldername", "f", "", "foldername")
	createFileCmd.Flags().StringVarP(&filename, "filename", "i", "", "filename")
	createFileCmd.Flags().StringVarP(&description, "description", "d", "", "description")

	createFileCmd.MarkFlagsRequiredTogether("username", "foldername", "filename")

	rootCmd.AddCommand(createFileCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createFileCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createFileCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
