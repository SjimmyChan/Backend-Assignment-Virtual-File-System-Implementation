package cmd_test

import (
	"testing"

	"github.com/SjimmyChan/IsCoollab-Backend-Assignment-Virtual-File-System-Implementation/cmd"
	"github.com/spf13/cobra"
)

func TestRenameFolderCmd(t *testing.T) {
	
	// store exist user inforamtion and initial json file
	exist_users_inforamtion := cmd.GetUsersInformation()
	cmd.InitialUsersInformation()

	testRenameFolderCmd := &cobra.Command{
		Use: "rename-folder",
		Run: RenameFolderCmdRunE,
	}
	RenameFolderCmdFlags(testRenameFolderCmd)

	cmd.SaveUsersInformation(cmd.CreateFakeData("user1", "folder1", "", "", ""))

	renameFolderTests := []struct {
		input 	[]string
		output 	string
	}{
		{
			input: []string{"-u", "abc", "-f", "folder1", "-n", "folder2"},
			output: "Error: The username:abc doesn't exist.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "abc", "-n", "folder2"},
			output: "Error: The foldername:abc doesn't exist.\n",
		},
		{
			input: []string{"-u", "user?", "-f", "folder1", "-n", "folder2"},
			output: "Error: username contains invalid chars.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder?", "-n", "folder2"},
			output: "Error: foldername contains invalid chars.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "-n", "folder?"},
			output: "Error: new_foldername contains invalid chars.\n",
		},
		{
			input: []string{"-u", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-f", "folder1", "-n", "folder2"},
			output: "Error: The username must be less than 30 chars and greater than 1 char.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-n", "folder2"},
			output: "Error: The foldername must be less than 30 chars and greater than 1 char.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "-n", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			output: "Error: The new_foldername must be less than 30 chars and greater than 1 char.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "-n", "folder2"},
			output: "Rename foldername from folder1 to folder2 successfully.\n",
		},
	}

	for _, test := range renameFolderTests {
		actual_output := execute(t, testRenameFolderCmd, test.input ... )

		expected_output := test.output
		if expected_output != actual_output {
			t.Errorf("Expected output '%s', but got '%s'", expected_output, actual_output)
		}
	}

	// store back originial user inforamtion
	cmd.SaveUsersInformation(exist_users_inforamtion)
}

func RenameFolderCmdRunE(c *cobra.Command, args []string) {
	username, _ := c.Flags().GetString("username")
	if err := cmd.CheckValidation(0, username, 30); err != nil {
		c.Print(err.Error())
		return
	}

	foldername, _ := c.Flags().GetString("foldername")
	if err := cmd.CheckValidation(1, foldername, 30); err != nil {
		c.Print(err.Error())
		return
	}

	new_foldername, _ := c.Flags().GetString("new_foldername")
	if err := cmd.CheckValidation(3, new_foldername, 30); err != nil {
		c.Print(err.Error())
		return
	}

	if succeed := cmd.RenameFolder(c, username, foldername, new_foldername); succeed {
		c.Println("Rename foldername from " + foldername + " to " + new_foldername + " successfully.")
	}
}

func RenameFolderCmdFlags(c *cobra.Command) {
	c.Flags().StringP("username", "u", "", "username")
	if err := c.MarkFlagRequired("username"); err != nil {
		c.Print(err.Error())
	}

	c.Flags().StringP("foldername", "f", "", "foldername")
	if err := c.MarkFlagRequired("foldername"); err != nil {
		c.Print(err.Error())
	}

	c.Flags().StringP("new_foldername", "n", "", "new_foldername")
	if err := c.MarkFlagRequired("new_foldername"); err != nil {
		c.Print(err.Error())
	}
}