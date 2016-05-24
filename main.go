package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

func checkRequiredArguments(c *cli.Context) error {
	if c.String("user") == "" {
		return errors.New("Please set user name.")
	}
	if c.String("repository") == "" {
		return errors.New("Please set repository name.")
	}
	if c.String("github_access_token") == "" {
		return errors.New("Please set GitHub access token.")
	}
	if c.String("circleci_api_token") == "" {
		return errors.New("Please set CircleCI API token.")
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "gradle-update-notifier"
	app.Usage = "notify gradle update"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Value: "",
			Usage: "GitHub user name",
		},
		cli.StringFlag{
			Name:  "repository, r",
			Value: "",
			Usage: "GitHub repository name",
		},
		cli.StringFlag{
			Name:   "weekday",
			Value:  "",
			Usage:  "Weekday of run the command('Sunday', 'Monday', ...)",
			EnvVar: "WEEKDAY",
		},
		cli.StringFlag{
			Name:   "github_access_token",
			Value:  "",
			Usage:  "GitHub access token",
			EnvVar: "GITHUB_ACCESS_TOKEN",
		},
		cli.StringFlag{
			Name:   "circleci_api_token",
			Value:  "",
			Usage:  "CiecleCI API token",
			EnvVar: "CIRCLECI_API_TOKEN",
		},
	}

	app.Action = func(c *cli.Context) error {
		weekday := c.String("weekday")
		today := time.Now().Weekday().String()
		if weekday != "" {
			if today != weekday {
				fmt.Printf("Today is %s. It is set to be executed in %s.\n", today, weekday)
				return nil
			}
		}
		err := checkRequiredArguments(c)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		reportData, err := readReportFileFromCircleCI(c.String("circleci_api_token"), c.String("user"), c.String("repository"))
		if err != nil {
			return cli.NewExitError("Report fetch error: "+err.Error(), 1)
		}

		report, err := parse(reportData)
		if err != nil {
			return cli.NewExitError("JSON parse error: "+err.Error(), 1)
		}

		err = reportToGithub(report, c.String("github_access_token"), c.String("user"), c.String("repository"))
		if err != nil {
			return cli.NewExitError("Report error: "+err.Error(), 1)
		}
		return nil

	}

	app.Run(os.Args)
}
