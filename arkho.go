package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

func main() {
	//  Subcommands
	cloneAwsCommand := flag.NewFlagSet("clone", flag.ExitOnError)

	cloneUrl := cloneAwsCommand.String("url", "", "Repositorie URL (Required)")
	cloneUser := cloneAwsCommand.String("user", "", "Codecommit username (Required)")
	clonePass := cloneAwsCommand.String("password", "", "Codecommit password (Required)")
	cloneProjectName := cloneAwsCommand.String("projectName", "", "Folder name to copy repo (optional)")

	if len(os.Args) < 2 {
		fmt.Println("list or count subcommand is required")
		os.Exit(1)
	}

	cloneAwsCommand.Parse(os.Args[2:])

	if *cloneUser == "" {
		cloneAwsCommand.PrintDefaults()
		os.Exit(1)
	}

	if *clonePass == "" {
		cloneAwsCommand.PrintDefaults()
		os.Exit(1)
	}

	if *cloneUrl == "" {
		cloneAwsCommand.PrintDefaults()
		os.Exit(1)
	}

	userEncode := url.QueryEscape(*cloneUser)
	passwordEncode := url.QueryEscape(*clonePass)

	repoUrl := getRepoPath(*cloneUrl)

	if repoUrl == "" {
		fmt.Println("Somthing is wrong with the repo URL")
		os.Exit(1)
	}

	fullUrlEncode := "https://" + userEncode + ":" + passwordEncode + "@" + repoUrl

	fmt.Println(" ")
	fmt.Println(fullUrlEncode)

	cloneRepo(fullUrlEncode, *cloneProjectName)
}

func getRepoPath(url string) string {
	urlSplit := strings.Split(url, "//")

	return urlSplit[1]
}

func cloneRepo(url string, projectName string) {
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
