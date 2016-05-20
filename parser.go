package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Available struct {
	Release     string
	Milestone   string
	Integration string
}

type Dependency struct {
	Group     string
	Version   string
	Available Available
	Name      string
}

type Item struct {
	Dependencies []Dependency
	count        int
}

type Report struct {
	Current    Item
	Exceeded   Item
	Outdated   Item
	Unresolved Item
	Count      int
}

func parse(filePath string) (Report, error) {
	var report Report

	reader, err := os.Open(filePath)
	if err != nil {
		return report, err
	}

	jsonStr, err := ioutil.ReadAll(reader)
	if err != nil {
		return report, err
	}

	err = json.Unmarshal([]byte(jsonStr), &report)

	if err != nil {
		return report, err
	}

	return report, nil
}
