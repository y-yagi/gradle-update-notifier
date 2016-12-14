package main

import (
	"io/ioutil"
	"testing"
)

func TestGenerateIssueBody(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("testdata/normal.json")
	report, _ := parse(jsonData, nil)
	releaseNotes := []ReleaseNote{ReleaseNote{Package: "com.google.firebase:firebase-core", URL: "https://firebase.google.com/support/releases"}, ReleaseNote{Package: "com.github.ben-manes:gradle-version-plugin", URL: ""}}

	issueBody := generateIssueBody(report, releaseNotes)
	expected := "* [ ] `com.github.ben-manes:gradle-versions-plugin:0.13.0`([Release Note](https://firebase.google.com/support/releases))\n* [ ] `com.google.firebase:firebase-core:9.4.0`\n"

	if issueBody != expected {
		t.Errorf("Expect \n%v\n, but \n%v", expected, issueBody)
	}
}
