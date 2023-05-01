/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (

	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register [username]",
	Short: "resgister a new user",
	Long: ``,
	Run: RegisterCmdRunE,
}

func RegisterCmdRunE(cmd *cobra.Command, args []string) {
	username, _ := cmd.Flags().GetString("username")
	if err := CheckValidation(0, username, 30); err != nil {
		cmd.Print(err.Error())
		return
	}

	if succeed := RegisterUser(cmd, username); succeed {
		cmd.Println("Add " + username + "successfully.")
	}
}

func RegisterCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("username", "u", "", "username")
	cmd.MarkFlagRequired("username")
}

func init() {

	RegisterCmdFlags(registerCmd)

	rootCmd.AddCommand(registerCmd)
}

func RegisterUser(cmd *cobra.Command, username string)(succeed bool) {
	
	users := GetUsersInformation()
	exist, _ := checkUserExist(users, username)
	if exist {
		cmd.Println("Error: The " + username + " has already existed.")
		return false
	}

	user := User{Username: username, Folders: []Folder{}, }
	users = append(users, user)
	
	if err := SaveUsersInformation(users); err != nil {
		cmd.Println(err.Error())
		return false
	}
	return true
}