package cmd_test

import (
	"testing"
	"time"

	"github.com/SjimmyChan/IsCoollab-Backend-Assignment-Virtual-File-System-Implementation/cmd"
	"github.com/spf13/cobra"
)

func TestListFilesCmd(t *testing.T) {
	
	// store exist user inforamtion and initial json file
	exist_users_inforamtion := cmd.GetUsersInformation()
	cmd.InitialUsersInformation()

	testListFilesCmd := &cobra.Command{
		Use: "list-files",
		Run: ListFilesCmdRunE,
	}
	ListFilesCmdFlags(testListFilesCmd)
	current_time := time.Now()
	cmd.SaveUsersInformation(cmd.CreateFakeListData("user1", []string{"folder1"}, []string{"file1", "file2"}, 1, current_time))

	listFilesTests := []struct {
		input 	[]string
		output 	string
	}{
		{
			input: []string{"-u", "user1", "-f", "folder1", "--sorted-name", "", "--sorted-created", ""},
			output: "[filename] | [description] | [created at] | [foldername] | [username]\n" +
					"file1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | folder1 | user1\n" +
					"file2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | folder1 | user1\n",
		},
		{
			input: []string{"-u", "abc", "-f", "folder1"},
			output: "Error: The username:abc doesn't exist.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "abc"},
			output: "Error: The foldername:abc doesn't exist.\n",
		},
		{
			input: []string{"-u", "user?", "-f", "folder1"},
			output: "Error: username contains invalid chars.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder?"},
			output: "Error: foldername contains invalid chars.\n",
		},
		{
			input: []string{"-u", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "-f", "folder1"},
			output: "Error: The username must be less than 30 chars and greater than 1 char.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			output: "Error: The foldername must be less than 30 chars and greater than 1 char.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "--sorted-name", "asc", "--sorted-created", ""},
			output: "[filename] | [description] | [created at] | [foldername] | [username]\n" +
					"file1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | folder1 | user1\n" +
					"file2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | folder1 | user1\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "--sorted-name", "desc", "--sorted-created", ""},
			output: "[filename] | [description] | [created at] | [foldername] | [username]\n" +
					"file2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | folder1 | user1\n" +
					"file1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | folder1 | user1\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "--sorted-name", "", "--sorted-created", "asc"},
			output: "[filename] | [description] | [created at] | [foldername] | [username]\n" +
					"file1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | folder1 | user1\n" +
					"file2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | folder1 | user1\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "--sorted-name", "", "--sorted-created", "desc"},
			output: "[filename] | [description] | [created at] | [foldername] | [username]\n" +
					"file2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | folder1 | user1\n" +
					"file1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | folder1 | user1\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "--sorted-name", "asc", "--sorted-created", "asc"},
			output: "Warning: Please only choose one of sorting factor.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "--sorted-name", "aaa", "--sorted-created", ""},
			output: "Error: Please use asc/desc as sorting method.\n",
		},
		{
			input: []string{"-u", "user1", "-f", "folder1", "--sorted-name", "", "--sorted-created", "aaa"},
			output: "Error: Please use asc/desc as sorting method.\n",
		},
	}

	for _, test := range listFilesTests {
		actual_output := execute(t, testListFilesCmd, test.input ... )

		expected_output := test.output
		if expected_output != actual_output {
			t.Errorf("Expected output '%s', but got '%s'", expected_output, actual_output)
		}
	}

	// test null file
	cmd.SaveUsersInformation(cmd.CreateFakeListData("user1", []string{"folder1"}, []string{}, 1, current_time))
	listFilesTest := struct {
		input 	[]string
		output 	string
	}{
		input: []string{"-u", "user1", "-f", "folder1", "--sorted-name", "", "--sorted-created", ""},
		output: "Warning: This folder is empty.\n",
	}
	actual_output := execute(t, testListFilesCmd, listFilesTest.input ... )
	expected_output := listFilesTest.output
	if expected_output != actual_output {
		t.Errorf("Expected output '%s', but got '%s'", expected_output, actual_output)
	}

	// store back originial user inforamtion
	cmd.SaveUsersInformation(exist_users_inforamtion)
}

func ListFilesCmdRunE(c *cobra.Command, args []string) {
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

	sorted_name, _ := c.Flags().GetString("sorted-name")
	sorted_created, _ := c.Flags().GetString("sorted-created")

	if sorted_name != "" && sorted_created != "" {
		c.Println("Warning: Please only choose one of sorting factor.")
		return
	}

	cmd.ListFiles(c, username, foldername, sorted_name, sorted_created)
}

func ListFilesCmdFlags(c *cobra.Command) {
	c.Flags().StringP("username", "u", "", "username")
	if err := c.MarkFlagRequired("username"); err != nil {
		c.Println(err.Error())
	}

	c.Flags().StringP("foldername", "f", "", "foldername")
	if err := c.MarkFlagRequired("foldername"); err != nil {
		c.Println(err.Error())
	}

	c.Flags().String("sorted-name", "", "sorted by username")

	c.Flags().String("sorted-created", "", "sorted by create time")
}