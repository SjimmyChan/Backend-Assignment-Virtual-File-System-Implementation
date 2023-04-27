/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register [username]",
	Short: "resgister a new user",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		if err := checkValidation(username, 30); err != nil {
			fmt.Println(err)
			return
		}

		registerUser(username)
		fmt.Println("register called")
	},
}

func init() {
	
	registerCmd.Flags().StringP("username", "u", "", "username")
	registerCmd.MarkFlagRequired("username")

	rootCmd.AddCommand(registerCmd)
}

func registerUser(username string) {
	fmt.Println(username)
	users := getUsersInformation()
	fmt.Println(users)
}