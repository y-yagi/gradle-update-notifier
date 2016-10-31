package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/codegangsta/cli"
	"github.com/pkg/errors"
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

	return nil
}

func commandFlags() []cli.Flag {
	return []cli.Flag{
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
			Name:  "report_file, f",
			Value: "",
			Usage: "The generated JSON file in Gradle Versions Plugin.",
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
}

func isRunDay(today, weekday string) bool {
	if weekday == "" || today == weekday {
		return true
	}
	return false
}

func main() {
	app := cli.NewApp()
	app.Name = "gradle-update-notifier"
	app.Usage = "notify gradle update"
	app.Version = "0.1.0"
	app.Flags = commandFlags()

	app.Action = func(c *cli.Context) error {
		weekday := c.String("weekday")
		today := time.Now().Weekday().String()
		if !isRunDay(today, weekday) {
			fmt.Printf("Today is %s. It is set to be executed in %s.\n", today, weekday)
			return nil
		}

		err := checkRequiredArguments(c)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		var reportData []byte
		if c.String("report_file") == "" {
			if c.String("circleci_api_token") == "" {
				return cli.NewExitError("Please set CircleCI API token.", 1)
			}

			reportData, err = readReportFileFromCircleCI(c.String("circleci_api_token"), c.String("user"), c.String("repository"))
			if err != nil {
				return cli.NewExitError(err.Error(), 1)
			}
		} else {
			reportData, err = ioutil.ReadFile(c.String("report_file"))
			if err != nil {
				return cli.NewExitError(errors.Wrap(err, "report file read error").Error(), 1)
			}
		}

		report, err := parse(reportData)
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		err = reportToGithub(report, c.String("github_access_token"), c.String("user"), c.String("repository"))
		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "build",
			Usage: "Triggers a new CircleCI build",
			Flags: commandFlags(),
			Action: func(c *cli.Context) error {
				weekday := c.String("weekday")
				today := time.Now().Weekday().String()
				if !isRunDay(today, weekday) {
					fmt.Printf("Today is %s. It is set to be executed in %s.\n", today, weekday)
					return nil
				}

				err := checkRequiredArguments(c)
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}

				if c.String("circleci_api_token") == "" {
					return cli.NewExitError("Please set CircleCI API token.", 1)
				}

				err = triggerBuild(c.String("circleci_api_token"), c.String("user"), c.String("repository"))
				if err != nil {
					return cli.NewExitError(err.Error(), 1)
				}

				return nil
			},
		},
	}
	app.Run(os.Args)
}
