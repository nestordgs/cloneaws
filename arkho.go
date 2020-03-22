package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// import "os/exec"
// import "log"
// import "bytes"
// import "strings"

func main() {
	//  Subcommands
	cloneAwsCommand := flag.NewFlagSet("clone", flag.ExitOnError)

	cloneUser := cloneAwsCommand.String("user", "", "Codecommit username (Required)")
	clonePassword := cloneAwsCommand.String("password", "", "Codecommit password (Required)")
	cloneUrl := cloneAwsCommand.String("url", "", "Repositorie URL (required)")

	flag.ErrHelp.Error()
	if len(os.Args) < 2 {
		fmt.Errorf("list or count subcommand is required")
		os.Exit(1)
	}

	cloneAwsCommand.Parse(os.Args[2:])

	if *cloneUser == "" {
		cloneAwsCommand.PrintDefaults()
		os.Exit(1)
	}

	if *clonePassword == "" {
		cloneAwsCommand.PrintDefaults()
		os.Exit(1)
	}

	if *cloneUrl == "" {
		cloneAwsCommand.PrintDefaults()
		os.Exit(1)
	}

	userEncode := url.QueryEscape(*cloneUser)
	passwordEncode := url.QueryEscape(*clonePassword)

	repoUrl := getRepoPath(*cloneUrl)

	if repoUrl == "" {
		fmt.Println("Somthing is wrong with the repo URL")
		os.Exit(1)
	}

	fullUrlEncode := "https://" + userEncode + ":" + passwordEncode + "@" + repoUrl

	fmt.Println(" ")
	fmt.Println(fullUrlEncode)

	// cloneRepo(fullUrlEncode)
}

func getRepoPath(url string) string {
	urlSplit := strings.Split(url, "//")

	return urlSplit[1]
}

// func cloneRepo(url string)  {

// 	cmd := exec.Command("git", "clone", url)

// 	// cmd.Stdin = strings.NewReader()
// 	var out byte.Buffer
// 	cmd.Stdout = &out

// 	err := cmd.Run()

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }p
