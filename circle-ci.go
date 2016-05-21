package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Artifact struct {
	Path       string
	PrettyPath string
	NodeIndex  int
	Url        string
}

func artifactApiUrl(apiToken, userName, repositoryName string) string {
	return fmt.Sprintf("https://circleci.com/api/v1/project/%v/%v/latest/artifacts?circle-token=%v",
		userName, repositoryName, apiToken)
}

func readReportFileFromCircleCI(apiToken, userName, repositoryName string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", artifactApiUrl(apiToken, userName, repositoryName), nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var artifacts []Artifact

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&artifacts)
	if err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	resp, err = http.Get(artifacts[0].Url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
