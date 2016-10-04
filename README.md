# gradle-update-notifier

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

`gradle-update-notifier` creates a GitHub issue with result of [Gradle Versions Plugin](https://github.com/ben-manes/gradle-versions-plugin/).

## Installation

```sh
$ go get github.com/y-yagi/gradle-update-notifier
```

## Usage

```
gradle-update-notifier -u <GitHub user name> -r <GitHub repository> -f <Generated file in Gradle Vesions Plugin(JSON format)> --github_access_token <GitHub access token>
```

Examples of the generated issue is [here](https://github.com/y-yagi/TravelBase/issues/89).

In addition to the local file, you can also specify the artifact of [CircleCI](https://circleci.com/).

Search for artifcat from the specified github user name and github repository. Instead of the file, you must specify the CircleCI API token.

```
gradle-update-notifier -u <GitHub user name> -r <GitHub repository> --github_access_token <GitHub access token> --circleci_api_token <CiecleCI API token>
```
In addition, GitHub access token and CircleCI API token can specify in the environment variable. For details, refer to help.


### Automatic execution

`gradle-update-notifier` works in [Heroku](https://www.heroku.com/). By using in conjunction with ,[Heroku Scheduler](https://devcenter.heroku.com/articles/scheduler) it is possible to run on a regular basis.

Heroku Scheduler supports 10 minutes, every hour, or every day only. If you want to run only once week, you can specify the day of the week to run in `weekday` option.

```
gradle-update-notifier -u <GitHub user name> -r <GitHub repository> --weekday Monday
```

