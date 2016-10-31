package main

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Available is type for android library
type Available struct {
	Release     string
	Milestone   string
	Integration string
}

// Dependency is type for android library
type Dependency struct {
	Group     string
	Version   string
	Available Available
	Name      string
}

// Item is type for android library
type Item struct {
	Dependencies []Dependency
	count        int
}

// Report is type for android library status
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

// Pkg retrun full pkg name
func (dependency *Dependency) Pkg() string {
	return dependency.Group + ":" + dependency.Name
}
