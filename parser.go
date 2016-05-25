package main

import (
	"encoding/json"

	"github.com/pkg/errors"
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

func parse(reportData []byte) (Report, error) {
	var report Report

	err := json.Unmarshal([]byte(reportData), &report)
	if err != nil {
		return report, errors.Wrap(err, "JSON parse failed")
	}

	return report, nil
}
