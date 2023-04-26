/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listFilesCmd represents the listFiles command
var listFilesCmd = &cobra.Command{
	Use:   "list-files [username] [foldername] [--sorted-name|--sorted-created] [asc|desc]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listFiles called")
	},
}

func init() {

	var username string
	var foldername string
	var filename string

	listFilesCmd.Flags().StringVarP(&username, "username", "u", "", "username")
	listFilesCmd.Flags().StringVarP(&foldername, "foldername", "f", "", "foldername")
	listFilesCmd.Flags().StringVarP(&filename, "filename", "i", "", "filename")
	
	listFilesCmd.MarkFlagsRequiredTogether("username", "foldername", "filename")

	rootCmd.AddCommand(listFilesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listFilesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listFilesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
