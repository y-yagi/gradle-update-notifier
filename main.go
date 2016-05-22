package main

import (
	"flag"
	"fmt"
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
		fmt.Println("Please set 'GITHUB_ACCESS_TOKEN' env")
		os.Exit(1)
	}

	circleciApiToken := os.Getenv("CIRCLECI_API_TOKEN")
	if circleciApiToken == "" {
		fmt.Println("Please set 'CIRCLECI_API_TOKEN' env")
		os.Exit(1)
	}

	var userName = flag.String("user", "", "GitHub user name")
	var repositoryName = flag.String("repository", "", "GitHub repository name")
	flag.Parse()

	if *userName == "" {
		fmt.Println("Please specifiy user name.")
		os.Exit(1)
	}

	if *repositoryName == "" {
		fmt.Println("Please specifiy repository name.")
		os.Exit(1)
	}

	reportData, err := readReportFileFromCircleCI(circleciApiToken, *userName, *repositoryName)
	if err != nil {
		fmt.Println("Report fetch error: ", err)
		os.Exit(1)
	}

	report, err := parse(reportData)
	if err != nil {
		fmt.Println("JSON parse error: ", err)
		os.Exit(1)
	}

	err = reportToGithub(report, githubAccessToken, *userName, *repositoryName)
	if err != nil {
		fmt.Println("Report error: ", err)
		os.Exit(1)
	}
}
