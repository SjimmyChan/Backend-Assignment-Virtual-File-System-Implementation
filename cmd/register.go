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
		if err := checkValidation(0, username, 30); err != nil {
			fmt.Println(err)
			return
		}

		if succeed := registerUser(username); succeed {
			fmt.Println("Add user:" + username + "successfully.")
		}
	},
}

func init() {
	
	registerCmd.Flags().StringP("username", "u", "", "username")
	registerCmd.MarkFlagRequired("username")

	rootCmd.AddCommand(registerCmd)
}

func registerUser(username string)(succeed bool) {
	
	users := getUsersInformation()
	exist, _ := checkUserExist(users, username)
	if exist {
		fmt.Println("Error: The username:" + username + " has already existed.")
		return false
	}

	user := User{Username: username, Folders: []Folder{}, }
	users = append(users, user)
	
	if err := saveUsersInformation(users); err != nil {
		fmt.Println(err)
		return false
	}
	return true
}