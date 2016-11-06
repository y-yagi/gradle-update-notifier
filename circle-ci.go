package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Artifact is type for path info
type Artifact struct {
	Path       string
	PrettyPath string
	NodeIndex  int
	URL        string
}

// APIEndpoint is circleci endpoint
const APIEndpoint = "https://circleci.com/api/v1"

func readReportFileFromCircleCI(apiToken, userName, repositoryName string) ([]byte, error) {
	client := &http.Client{}
	artifactAPIURL := fmt.Sprintf(APIEndpoint+"/project/%v/%v/latest/artifacts?circle-token=%v",
		userName, repositoryName, apiToken)
	req, err := http.NewRequest("GET", artifactAPIURL, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "artifacts API failed")
	}

	var artifacts []Artifact

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&artifacts)
	if err != nil {
		resp.Body.Close()
		return nil, errors.Wrap(err, "artifacts API response parse failed")
	}
	resp.Body.Close()

	resp, err = http.Get(artifacts[0].URL)
	if err != nil {
		return nil, errors.Wrap(err, "artifacts get failed")
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func triggerBuild(apiToken, userName, repositoryName string) error {
	client := &http.Client{}
	triggerBuildAPIURL := fmt.Sprintf(APIEndpoint+"/project/%v/%v/tree/master?circle-token=%v",
		userName, repositoryName, apiToken)
	req, err := http.NewRequest("POST", triggerBuildAPIURL, nil)
	_, err = client.Do(req)
	if err != nil {
		return errors.Wrap(err, "trigger build API failed")
	}
	return nil
}
