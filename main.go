package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

type Users struct {
	Users []User `json: "users"`
}

type User struct {
	Alias    string `json: "alias"`
	Email    string `json: "email"`
	Password string `json: "password"`
}

func main() {
	//  Subcommands
	cloneAwsCommand := flag.NewFlagSet("clone", flag.ExitOnError)

	cloneUrl := cloneAwsCommand.String("url", "", "Repositorie HTTPS/SSH URL (Required)")
	cloneProfile := cloneAwsCommand.String("profile", "", "Profile name to use cloning repository (Required)")
	cloneProjectName := cloneAwsCommand.String("projectName", "", "Folder name to copy repo (optional)")

	if len(os.Args) < 2 {
		cloneAwsCommand.PrintDefaults()
		log.Fatalln("Subcommand is required")
	}

	switch os.Args[1] {
	case "clone":
		cloneAwsCommand.Parse(os.Args[2:])

		if *cloneProfile == "" {
			cloneAwsCommand.PrintDefaults()
			log.Fatalln("No profile value to find")
		}
		credentials := findCredentials(*cloneProfile)

		if *cloneUrl == "" {
			cloneAwsCommand.PrintDefaults()
			log.Fatalln("")
		}

		urlRepo := getRepoPath(*cloneUrl)

		userEncode := url.QueryEscape(credentials.Email)
		passwordEncode := url.QueryEscape(credentials.Password)

		fullUrlEncode := "https://" + userEncode + ":" + passwordEncode + "@" + urlRepo

		executeCloneCommand(fullUrlEncode, *cloneProjectName)
	}
}

func getRepoPath(url string) string {
	urlSplit := strings.Split(url, "//")

	return urlSplit[1]
}

func executeCloneCommand(url string, projectName string) {
	if projectName != "" {
		if err := exec.Command("git", "clone", url, projectName).Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		if err := exec.Command("git", "clone", url).Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func findCredentials(alias string) User {
	user, err := user.Current()
	credentialsPath := user.HomeDir + "/.codecommit/" + "credentials.json"
	jsonFile, err := os.Open(credentialsPath)
	if err != nil {
		log.Fatalln(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		if alias == users.Users[i].Alias {
			return users.Users[i]
		}
	}

	fmt.Println("No profile with the value:", alias)
	log.Fatalln("Please add the user/password/alias in the credentials.json file inside" + user.HomeDir + "/.aws folder")
	return User{}
}
