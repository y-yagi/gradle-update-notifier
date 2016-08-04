package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

func generateIssueBody(report Report, releaseNotes []ReleaseNote) string {
	var body string
	var dependency Dependency

	for i := 0; i < len(report.Outdated.Dependencies); i++ {
		dependency = report.Outdated.Dependencies[i]
		body += fmt.Sprintf("* [ ] `%v:%v`", dependency.Pkg(), dependency.Available.Release)

		if len(releaseNotes[i].Url) > 1 {
			body += fmt.Sprintf("([Release Note](%v))\n", releaseNotes[i].Url)
		} else {
			body += "\n"
		}
	}
	return body
}

func reportToGithub(report Report, githubAccessToken, userName, repositoryName string) error {
	var pkgs string
	for _, dependency := range report.Outdated.Dependencies {
		pkgs += dependency.Pkg() + ","
	}
	pkgs = strings.TrimRight(pkgs, ",")

	releaseNotes, err := getReleaseNotes(pkgs)
	if err != nil {
		return err
	}

	body := generateIssueBody(report, releaseNotes)

	if len(body) == 0 {
		// No libraries need to update
		return nil
	}

	currentTime := time.Now()
	title := "dependency-updates-" + currentTime.Format("20060102150405")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	issueRequest := &github.IssueRequest{Title: &title, Body: &body}

	_, _, err = client.Issues.Create(userName, repositoryName, issueRequest)
	if err != nil {
		return errors.Wrap(err, "issue create failed")
	}
	return nil
}
