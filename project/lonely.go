package project

import (
	"encoding/json"
	"fmt"
	"os"
)

type project struct {
	Version   string `json:"version"`
	BuildDate string `json:"buildDate"`
}

func GetProjectData() (project, error) {

	d, err := os.ReadFile("project/lonely.json")

	if err != nil {
		return project{}, fmt.Errorf("error reading project json data: %w", err)
	}
	projectData := project{}

	if err = json.Unmarshal(d, &projectData); err != nil {
		return project{}, fmt.Errorf("error unmarshaling project json data: %w", err)
	}

	return projectData, nil
}
