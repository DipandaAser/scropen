package main

import (
	"encoding/json"
)

type Project struct {
	Url string `json:"url"`
	Name string `json:"name"`
	Description string `json:"description"`
	Stars string `json:"stars"`
}

type Projects []Project

func toJson(project * Project) string {
	byteArray, err := json.MarshalIndent(project, "", "")
	if err != nil {
		return ""
	}
	return string(byteArray)
}