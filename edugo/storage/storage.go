package storage

import (
	"edugo/models"
	"os"
)

type Data struct {
	Users []models.User `json:"users"`	
	Courses []models.ICourse `json:"courses"`
	Reviews []models.Review `json:"reviews"`

}

var filePath = "storage/data.json"

func LoadData() (*Data, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	func SaveData(data *Data) error {
	bytes, err := json.Marshal(data)
	if err != nil {	
		return err
	}


	err = os.WriteFile(filePath, bytes, 0644)
}
