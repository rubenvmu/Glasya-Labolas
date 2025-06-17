package parsers

import (
	"encoding/csv"
	"os"
	"glasya-labolas/models"
)

func LoadCSV(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(records) < 2 {
		return nil
	}

	for _, row := range records[1:] {
		e := models.Entity{
			Nombre: row[0],
			IP:     row[1],
			Hash:   row[2],
		}
		e.Save()
	}
	return nil
}