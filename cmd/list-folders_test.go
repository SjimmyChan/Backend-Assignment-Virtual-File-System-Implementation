package cmd_test

import (
	"testing"
	"time"

	"github.com/SjimmyChan/IsCoollab-Backend-Assignment-Virtual-File-System-Implementation/cmd"
	"github.com/spf13/cobra"
)

func TestListFoldersCmd(t *testing.T) {
	
	// store exist user inforamtion and initial json file
	exist_users_inforamtion := cmd.GetUsersInformation()
	cmd.InitialUsersInformation()

	testListFoldersCmd := &cobra.Command{
		Use: "list-folders",
		Run: ListFoldersCmdRunE,
	}
	ListFoldersCmdFlags(testListFoldersCmd)
	current_time := time.Now()
	cmd.SaveUsersInformation(cmd.CreateFakeListData("user1", []string{"folder1", "folder2"}, []string{}, 0, current_time))

	listFoldersTests := []struct {
		input 	[]string
		output 	string
	}{
		{
			input: []string{"-u", "user1", "--sorted-name", "", "--sorted-created", ""},
			output: "[foldername] | [description] | [created at] | [username]\n" +
					"folder1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | user1\n" +
					"folder2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | user1\n",
		},
		{
			input: []string{"-u", "abc"},
			output: "Error: The username:abc doesn't exist.\n",
		},
		{
			input: []string{"-u", "user?"},
			output: "Error: username contains invalid chars.\n",
		},
		{
			input: []string{"-u", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
			output: "Error: The username must be less than 30 chars and greater than 1 char.\n",
		},
		{
			input: []string{"-u", "user1", "--sorted-name", "asc", "--sorted-created", ""},
			output: "[foldername] | [description] | [created at] | [username]\n" +
					"folder1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | user1\n" +
					"folder2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | user1\n",
		},
		{
			input: []string{"-u", "user1", "--sorted-name", "desc", "--sorted-created", ""},
			output: "[foldername] | [description] | [created at] | [username]\n" +
					"folder2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | user1\n" +
					"folder1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | user1\n",
		},
		{
			input: []string{"-u", "user1", "--sorted-name", "", "--sorted-created", "asc"},
			output: "[foldername] | [description] | [created at] | [username]\n" +
					"folder1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | user1\n" +
					"folder2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | user1\n",
		},
		{
			input: []string{"-u", "user1", "--sorted-name", "", "--sorted-created", "desc"},
			output: "[foldername] | [description] | [created at] | [username]\n" +
					"folder2 |  | " + current_time.Add(time.Hour*1).Format("01-02-2006 15:04:05") + " | user1\n" +
					"folder1 |  | " + current_time.Format("01-02-2006 15:04:05") + " | user1\n",
		},
		{
			input: []string{"-u", "user1", "--sorted-name", "asc", "--sorted-created", "asc"},
			output: "Warning: Please only choose one of sorting factor.\n",
		},
		{
			input: []string{"-u", "user1", "--sorted-name", "aaa", "--sorted-created", ""},
			output: "Error: Please use asc/desc as sorting method.\n",
		},
		{
			input: []string{"-u", "user1", "--sorted-name", "", "--sorted-created", "aaa"},
			output: "Error: Please use asc/desc as sorting method.\n",
		},
	}

	for _, test := range listFoldersTests {
		actual_output := execute(t, testListFoldersCmd, test.input ... )

		expected_output := test.output
		if expected_output != actual_output {
			t.Errorf("Expected output '%s', but got '%s'", expected_output, actual_output)
		}
	}

	// test null file
	cmd.SaveUsersInformation(cmd.CreateFakeListData("user1", []string{}, []string{}, 0, current_time))
	listFilesTest := struct {
		input 	[]string
		output 	string
	}{
		input: []string{"-u", "user1", "--sorted-name", "", "--sorted-created", ""},
		output: "Warning: The user1 doesn't have any folders.\n",
	}
	actual_output := execute(t, testListFoldersCmd, listFilesTest.input ... )
	expected_output := listFilesTest.output
	if expected_output != actual_output {
		t.Errorf("Expected output '%s', but got '%s'", expected_output, actual_output)
	}

	// store back originial user inforamtion
	cmd.SaveUsersInformation(exist_users_inforamtion)
}

func ListFoldersCmdRunE(c *cobra.Command, args []string) {
	username, _ := c.Flags().GetString("username")
	if err := cmd.CheckValidation(0, username, 30); err != nil {
		c.Print(err.Error())
		return
	}

	sorted_name, _ := c.Flags().GetString("sorted-name")
	sorted_created, _ := c.Flags().GetString("sorted-created")

	if sorted_name != "" && sorted_created != "" {
		c.Println("Warning: Please only choose one of sorting factor.")
		return
	}

	cmd.ListFolders(c, username, sorted_name, sorted_created)
}

func ListFoldersCmdFlags(c *cobra.Command) {
	c.Flags().StringP("username", "u", "", "username")
	if err := c.MarkFlagRequired("username"); err != nil {
		c.Print(err.Error())
	}

	c.Flags().String("sorted-name", "", "sorted by username")

	c.Flags().String("sorted-created", "", "sorted by create time")
}