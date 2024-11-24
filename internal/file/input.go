package file


import (
	"encoding/json"
	"fmt"
	"os"
)

/*
  * check if the JSON file is valid and matches the expected format

*/

func ValidateJSONFormat (fileName string, v any) error {
	file, err:= os.Open(fileName)
	if err != nil {
		return fmt.Errorf("Cannot open the file: %v", err)
	}

	defer file.Close()

	//decode to check the format
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(v); err != nil {
		return fmt.Errorf("Invalid JSON format: %v", err)
	}
	return nil
}