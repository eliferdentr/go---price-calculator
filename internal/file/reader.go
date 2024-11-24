package file

import (
	"encoding/json"
	"errors"
	"os"
)

/*
 * reads data from a json file and unmarshalls it into a given struct
 */

func ReadJSONFile (fileName string, v any) error {
	if fileName == "" {
		return errors.New("File name cannot be empty")
	}
	file , err := os.Open(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
}