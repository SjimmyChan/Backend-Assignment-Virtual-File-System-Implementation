package cmd_test

import (
	"bytes"
	"testing"

	"github.com/SjimmyChan/IsCoollab-Backend-Assignment-Virtual-File-System-Implementation/cmd"
	"github.com/spf13/cobra"
)

func TestRegisterCmd(t *testing.T) {
	
	// store exist user inforamtion and initial json file
	exist_users_inforamtion := cmd.GetUsersInformation()
	cmd.InitialUsersInformation()

	testRegisterCmd := &cobra.Command{
		Use: "test-register",
		Run: RegisterCmdRunE,
	}
	RegisterCmdFlags(testRegisterCmd)

	registerTests := []struct {
		input 	[]string
		output 	string
	}{
		{
			input: []string{"-u", "user1"},
			output: "Add user1 successfully.\n",
		},
		{
			input: []string{"-u", "user1"},
			output: "Error: The user1 has already existed.\n",
		},
		{
			input: []string{"-u", "user?"},
			output: "Error: username contains invalid chars.\n",
		},
		{
			input: []string{"-u", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			output: "Error: The username must be less than 30 chars and greater than 1 char.\n",
		},
	}

	for _, test := range registerTests {
		actual_output := execute(t, testRegisterCmd, test.input ... )

		expected_output := test.output
		if expected_output != actual_output {
			t.Errorf("Expected output '%s', but got '%s'", expected_output, actual_output)
		}
	}

	// store back originial user inforamtion
	cmd.SaveUsersInformation(exist_users_inforamtion)
}

func execute(t *testing.T, c *cobra.Command, args ...string) (string) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	c.Execute()
	
	return buf.String()
}

func RegisterCmdRunE(c *cobra.Command, args []string) {
	username, _ := c.Flags().GetString("test-username")
	if err := cmd.CheckValidation(0, username, 30); err != nil {
		c.Print(err.Error())
		return
	}

	if succeed := cmd.RegisterUser(c, username); succeed {
		c.Println("Add " + username + " successfully.")
	}
}

func RegisterCmdFlags(c *cobra.Command) {
	c.Flags().StringP("test-username", "u", "", "test-username")
	c.MarkFlagRequired("test-username")
}