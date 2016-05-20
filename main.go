package main

import (
	"log"
	"os"
)

func main() {

	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if githubAccessToken == "" {
		log.Fatal("Please set 'GITHUB_ACCESS_TOKEN' env")
	}

	repository := os.Getenv("REPOSITORY")
	if repository == "" {
		log.Fatal("Please set 'REPOSITORY' env")
	}

	report, err := parse("report.json")
	if err != nil {
		log.Fatal(err)
	}

	err = reportToGithub(report, githubAccessToken, repository)
	if err != nil {
		log.Fatal(err)
	}
}
