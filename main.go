package main

import (
	"flag"
	"log"
	"os"
	"time"
)

func main() {
	weekday := os.Getenv("WEEKDAY")
	if weekday != "" {
		if string(time.Now().Weekday()) != weekday {
			// Do nothing.
			return
		}
	}

	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	if githubAccessToken == "" {
		log.Fatal("Please set 'GITHUB_ACCESS_TOKEN' env")
	}

	circleciApiToken := os.Getenv("CIRCLECI_API_TOKEN")
	if circleciApiToken == "" {
		log.Fatal("Please set 'CIRCLECI_API_TOKEN' env")
	}

	var userName = flag.String("user", "", "GitHub user name")
	var repositoryName = flag.String("repository", "", "GitHub repository name")
	flag.Parse()

	if *userName == "" {
		log.Fatal("Please specifiy user name.")
	}

	if *repositoryName == "" {
		log.Fatal("Please specifiy repository name.")
	}

	reportData, err := readReportFileFromCircleCI(circleciApiToken, *userName, *repositoryName)
	if err != nil {
		log.Fatal("Report fetch error: ", err)
	}

	report, err := parse(reportData)
	if err != nil {
		log.Fatal("JSON parse error: ", err)
	}

	err = reportToGithub(report, githubAccessToken, *userName, *repositoryName)
	if err != nil {
		log.Fatal("Report error: ", err)
	}
}
