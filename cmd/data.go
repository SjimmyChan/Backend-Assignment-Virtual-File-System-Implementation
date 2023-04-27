package cmd


import (
	"regexp"
	"errors"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

type file struct {
	filename string `json:"filename"`
	description string `json:"description"`
	created_at string `json:"created_at"`
}

type folder struct {
	foldername string `json:"foldername"`
	description string `json:"description"`
	created_at string `json:"created_at"`
	files []file `json:"files"`
}

type user struct {
	username string `json:"username"`
	folders []folder `json:"folders`
}

func getUsersInformation()(users []user) {
	fileBytes, err := ioutil.ReadFile(getFilePath())

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(fileBytes, &users)

	if err != nil {
		panic(err)
	}

	return users
}

func saveUsersInformation(users []user) {

	userBytes, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(getFilePath(), userBytes, 0644)
	if err != nil {
		panic(err)
	}
}

func checkValidation(input string, length int) (err error) {
	if len(input) > length || len(input) == 0 {
		return errors.New("Error: username/foldername/filename must be less than " + strconv.Itoa(length) + " and greater than 1 charactors!")
	}

	var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")

	if !alphanumeric.MatchString(input) {
		return errors.New("Error: username/foldername/filename must contains only alphabet, number and underscores!")
	}

	return nil
}

func getFilePath() (path string) {
	absPath, err := filepath.Abs("../IsCoollab-Backend-Assignment-Virtual-File-System-Implementation/cmd/data/users_information.json")
	if err != nil {
		panic(err)
		return
	}
	return absPath
}