
# Backend Assignment Virtual File System Implementation - IsCoolLab

This is the project dedicated to IsCoolLab's first round coding challenge.

# Installation

In order to make this project work properly, we need to install Go(version 1.20.3), for detail inforamtion on installation please visit [go](https://go.dev/doc/install).

Also, this project was implemented using cobra, please visite its official [documentation](github.com/spf13/cobra) to install.

# Usage

Before using the features, please make sure you're under `IsCoollab-Backend-Assignment-Virtual-File-System-Implementation` folder.

## 1. User Registration

- `register [-u/--username] [username]`

	**Response:**
	
	- Add `[username]` successfully.
	- Error: The username:`[username]` has already existed.
	- Error: username contains invalid chars.
		
## 2. Folder Management

- `create-folder [-u/--username] [username] [-f/--foldername] [foldername] [-d/--description]? [description]?`

	**Response:**	
	- Create `[foldername]` in `[username]` successfully.
	- Error: The username:`[username]` doesn't exist.
	- Error: The foldername:`[foldername]` has already existed.
	- Error: username contains invalid chars.
	- Error: foldername contains invalid chars.
		
- `delete-folder [-u/--username] [username] [-f/--foldername] [foldername]`

	**Response:**
	- Delete `[foldername]` in `[username]` successfully.
	- Error: The username:`[username]` doesn't exist.
	- Error: The foldername:`[foldername]` doesn't exist.
	- Error: username contains invalid chars.
	- Error: foldername contains invalid chars.

- `list-folders [-u/--username] [username] [--sort-name] [asc/desc] [--sort-created] [asc/desc]`

	**Response:**
	- List all the folders within the `[username]` scope in following formats:
			`[foldername] | [description] | [created at] | [username]`. 
			
		Each field should be separated by whitespace or tab characters. The `[created at]` is a human-readable date/time format.
		
		The order of printed folder information is determined by the `--sort-name` or `--sort-created` combined with `asc` or `desc` flgs.
		
		The `--sort-name` flag means sorting by `[foldername] `.
		
		If neither `--sort-name` nor `--sort-created` is provided, sort the list by `[foldername]` in ascending order.
	- Warning: The `[username]` doesn't have any folders.
	- Error: The username:`[username]` doesn't exist.
	- Error: username contains invalid chars.
	- If the input has enable both `--sort-name` and `--sort-created`, will have: `Warning: Please only choose one of sorting factor.`
	- If any of `--sort-name`'s and `--sort-created`'s value is not `asc` or `desc`, will have `Error: Please use asc/desc as sorting method.`

- `rename-folder [-u/--username] [username] [-f/--foldername] [foldername] [-n/--new-folder-name] [new-folder-name]`

	**Response:**
	- Rename `[foldername]` to `[new-foldername]` successfully.
	- Error: The username:`[username]` doesn't exist.
	- Error: The foldername:`[foldername]` doesn't exist.
	- Error: username contains invalid chars.
	- Error: foldername contains invalid chars.
	- Error: new-folder-name contains invalid chars.
	- If `[new-folder-name]` is the same as `[foldername]`, will have `Warning: The foldername is already called [new-folder-name].`
	- If `[new-folder-name]` has already used by others (exclude `[foldername]`), will have `"Error: Cannot change foldername to [new-folder-name], since it has already existed.`
## File Management
	
- `create-file [-u/--username] [username] [-f/--foldername] [foldername] [-i/--filename] [filename] [-d/--description]? [description]?`

	**Response:**
	- Create `[filename]` in `[username]`/`[foldername]` successfully.
	- Error: The username:`[username]` doesn't exist.
	- Error: The foldername:`[foldername]` doesn't exist.
	- Error: The filename:`[filename]` has already existed.
	- Error: username contains invalid chars.
	- Error: foldername contains invalid chars.
	- Error: filename contains invalid chars.

- `delete-file [-u/--username] [username] [-f/--foldername] [foldername] [-i/--filename] [filename]`

	**Response:**
	- Delete `[filename]` in `[username]`/`[foldername]` successfully.
	- Error: The username:`[username]` doesn't exist.
	- Error: The foldername:`[foldername]` doesn't exist.
	- Error: The filename:`[filename]` doesn't exist.
	- Error: username contains invalid chars.
	- Error: foldername contains invalid chars.
	- Error: filename contains invalid chars.

- `list-folders [-u/--username] [username] [-f/--foldername] [foldername] [--sort-name] [asc/desc] [--sort-created] [asc/desc]`

	**Response:**
	- List all the folders within the `[username]` scope in following formats:
			`[filename] | [description] | [created at] | [foldername] | [username]`. 
			
		Each field should be separated by whitespace or tab characters. The `[created at]` is a human-readable date/time format.
		
		The order of printed folder information is determined by the `--sort-name` or `--sort-created` combined with `asc` or `desc` flgs.
		
		The `--sort-name` flag means sorting by `[filename] `.
		
		If neither `--sort-name` nor `--sort-created` is provided, sort the list by `[filename]` in ascending order.
	- Warning: This folder is empty.
	- Error: The username:`[username]` doesn't exist.
	- Error: The username:`[foldername]` doesn't exist.
	- Error: username contains invalid chars.
	- Error: foldername contains invalid chars.
	- If the input has enable both `--sort-name` and `--sort-created`, will have: `Warning: Please only choose one of sorting factor.`
	- If any of `--sort-name`'s and `--sort-created`'s value is not `asc` or `desc`, will have `Error: Please use asc/desc as sorting method.`

# Example
		
The `#` below is a prompt to inform the user that they can type commands.The following examples
demonstrate the usage of various commands in the virtual file system:
**Warm Reminder: Please make sure you're under `IsCoollab-Backend-Assignment-Virtual-File-System-Implementation` folder.**

Register two users, user1 and user2
```plaintext
# go run main.go register -u user1
Add user1 successfully.
```

Register user2
```plaintext
# go run main.go register -u user2
Add user2 successfully.
```

Create a folder for user1 and user2 with the same folder name
```plaintext
# go run main.go create-folder -u user1 -f folder1
Create folder1 in user1 successfully.
# go run main.go create-folder -u user2 -f folder1
Create folder1 in user2 successfully.
```
Attempt to create a folder with an existing name for user1
```plaintext
# go run main.go create-folder -u user1 -f folder1
Error: foldername:folder1 has already existed.
```
Create a folder with a description for user1
```plaintext
# go run main.go create-folder -u user1 -f folder2 -d this-is-folder-2
Create folder2 in user1 successfully.
```

List folders for user1 sorted by name in ascending order
```plaintext
# go run main.go list-folders -u user1 --sort-name asc
folder1 | | 01-01-2023 15:00:00 | user1
folder2 | this-is-folder-2 | 01-01-2023 15:00:10 | user1
```

List folders for user2
```plaintext
# go run main.go list-folders -u user2
folder1 |  | 01-01-2023 15:05:00 | user2
```

Create a file with a description for user1 in folder1
```plaintext
# go run main.go create-file -u user1 -f folder1 -i file1 -d this-is-file1
Create file1 in user1/folder1 successfully.
```

Create a file named config with a description for user1 in folder1
```plaintext
# go run main.go create-file -u user1 -f folder1 -i config -d a-config-file
Create config in user1/folder1 successfully.
```

Attempt to create an existing file.
```plaintext
# go run main.go create-file -u user1 -f folder1 -i config -d a-config-file
Error: the filename:config has already existed.
```

Attempt to create an file for an unregistered user.
```plaintext
# go run main.go create-file user-abc folder-abc config a-config-file
Error: The username:user-abc doesn't exist.
```

Attempt to type a unsupported command
```plaintext
# go run main.go list data
Error: unknown command "list" for "main.go"

Did you mean this?
        list-files
        list-folders
```

Attempt to list files with incorrect flags
```plaintext
# go run main.go list-files -u user1 -f folder1 --sort a
Error: unknown flag: --sort
Usage:
  main.go list-files [username] [foldername] [--sorted-name|--sorted-created] [asc|desc]
```
```plaintext
# go run main.go list-files -u user1 -f folder1 --sort-name desc
file1 | this-is-file1 | 01-01-2023 15:00:20 | folder1 | user1
config | a-config-file | 01-01-2023 15:00:30 | folder1 | user1
```

# Testing
To run unit test for this project, we need to change our directory:
```bash
cd cmd
```
Then we can start runing unit test:
```bash
go test -v
```
This will listed the test results for each functionalities, which will print out the result as following:
```bash
=== RUN TestCreateFileCmd
--- PASS: TestCreateFileCmd (0.00s)

=== RUN TestCreateFolderCmd
--- PASS: TestCreateFolderCmd (0.00s)

=== RUN TestDeleteFileCmd
--- PASS: TestDeleteFileCmd (0.00s)

=== RUN TestDeleteFolderCmd
--- PASS: TestDeleteFolderCmd (0.00s)

=== RUN TestListFilesCmd
--- PASS: TestListFilesCmd (0.00s)

=== RUN TestListFoldersCmd
--- PASS: TestListFoldersCmd (0.00s)

=== RUN TestRegisterCmd
--- PASS: TestRegisterCmd (0.00s)

=== RUN TestRenameFolderCmd
--- PASS: TestRenameFolderCmd (0.00s)

PASS
ok github.com/SjimmyChan/IsCoollab-Backend-Assignment-Virtual-File-System-Implementation/cmd 0.017s
```