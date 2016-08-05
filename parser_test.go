package main

import (
	"io/ioutil"
	"testing"
)

func TestParseSuccess(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("testdata/normal.json")
	report, err := parse(jsonData)
	if err != nil {
		t.Errorf("Get error: %v", err)
	}

	if len(report.Outdated.Dependencies) != 2 {
		t.Errorf("Expect %v, but %v", 2, len(report.Outdated.Dependencies))
	}
}

func TestParseFailed(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("testdata/invalid.json")
	_, err := parse(jsonData)
	if err == nil {
		t.Errorf("Expected parse error. But not")
	}
}
