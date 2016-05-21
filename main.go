package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s PATTERN\n", os.Args[0])
}

func main() {
	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if githubAccessToken == "" {
		log.Fatal("Please set 'GITHUB_ACCESS_TOKEN' env")
	}

	var userName = flag.String("user", "", "GitHub user name")
	var repositoryName = flag.String("repository", "", "GitHub repository name")
	var reportPath = flag.String("report", "", "report file path")
	flag.Parse()

	if *userName == "" {
		log.Fatal("Please specifiy user name.")
	}

	if *repositoryName == "" {
		log.Fatal("Please specifiy repository name.")
	}

	if *reportPath == "" {
		log.Fatal("Please specifiy report file path.")
	}

	report, err := parse(*reportPath)
	if err != nil {
		log.Fatal(err)
	}

	err = reportToGithub(report, githubAccessToken, *userName, *repositoryName)
	if err != nil {
		log.Fatal(err)
	}
}
