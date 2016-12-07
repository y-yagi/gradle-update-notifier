package main

import (
	"io/ioutil"
	"regexp"
	"testing"
)

func TestParseSuccess(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("testdata/normal.json")
	report, err := parse(jsonData, nil)
	if err != nil {
		t.Errorf("Get error: %v", err)
	}

	if len(report.Outdated.Dependencies) != 2 {
		t.Errorf("Expect %v, but %v", 2, len(report.Outdated.Dependencies))
	}
}

func TestParseFailed(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("testdata/invalid.json")
	_, err := parse(jsonData, nil)
	if err == nil {
		t.Errorf("Expected parse error. But not")
	}
}

func TestParseExcludePkgs(t *testing.T) {
	jsonData, _ := ioutil.ReadFile("testdata/exclude_pkgs.json")
	excludePattern := regexp.MustCompile("com.android.databinding")
	report, err := parse(jsonData, excludePattern)
	if err != nil {
		t.Errorf("Get error: %v", err)
	}

	if len(report.Outdated.Dependencies) != 2 {
		t.Errorf("Expect %v, but %v", 2, len(report.Outdated.Dependencies))
	}

	pkg := report.Outdated.Dependencies[0]
	if pkg.Pkg() != "com.github.ben-manes:gradle-versions-plugin" {
		t.Errorf("Expect 'com.github.ben-manes:gradle-versions-plugin', but %v", pkg.Pkg())
	}

	pkg = report.Outdated.Dependencies[1]
	if pkg.Pkg() != "com.google.firebase:firebase-core" {
		t.Errorf("Expect 'com.google.firebase:firebase-core', but %v", pkg.Pkg())
	}
}
