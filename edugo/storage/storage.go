package storage

import (
	"encoding/json"
	"os"
)

type Data struct {
	Users   []models.User   `json:"users"`
	Courses []models.Course `json:"courses"`
	Reviews []models.Review `json:"reviews"`
}

var data Data

func LoadData() error {
	file, err := os.ReadFile("storage/data.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func SaveData() error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("storage/data.json", file, 0644)
}
