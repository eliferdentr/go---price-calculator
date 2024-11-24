package file

import (
	"encoding/json"
	"errors"
	"os"
)

/*
 * writes data to a JSON file in a readable format
 */

func WriteJSONFile(fileName string, v any) error {
	if fileName == "" {
		return errors.New("File name cannot be empty")
	}
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(v); err != nil {
		return err
	}

	return nil

}
