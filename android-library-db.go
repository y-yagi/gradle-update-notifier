package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type ReleaseNote struct {
	Package string
	Url     string
}

const ANDROID_LIBRARY_DB_API = "https://android-library-db.herokuapp.com"

func getReleaseNotes(pkgs string) ([]ReleaseNote, error) {
	var releaseNotes []ReleaseNote

	client := &http.Client{}
	url := fmt.Sprintf(ANDROID_LIBRARY_DB_API+"/release_notes?packages=%v", pkgs)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "release_notes API failed")
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&releaseNotes)
	if err != nil {
		return nil, errors.Wrap(err, "release_notes API response parse failed")
	}

	return releaseNotes, nil
}
