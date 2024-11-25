package file

import (
	"os" 
	"testing"
)

func TestReadJSONFile (t *testing.T) {
	data := `{"prices": [100, 200], "tax_rates": [10, 20]}`
	err := os.WriteFile("temp.json", []byte(data), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %s", err)
	}

	defer os.Remove("temp.json")

	var pc struct {
		Prices   []float64 `json:"prices"`
		TaxRates []float64 `json:"tax_rates"`
	}

	err = ReadJSONFile("temp.json", &pc)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}

	if len(pc.Prices) != 2 || len(pc.TaxRates) != 2 {
		t.Errorf("Expected 2 prices and 2 tax rates, got %d and %d", len(pc.Prices), len(pc.TaxRates))
	}

}

func TestWriteJSONFile(t *testing.T) {
	data := map[string]interface{}{
		"prices":   []float64{100, 200},
		"tax_rates": []float64{10, 20},
	}

	err := WriteJSONFile("output.json", data)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	defer os.Remove("output.json")

	
	if _, err := os.Stat("output.json"); os.IsNotExist(err) {
		t.Errorf("Expected file 'output.json' to exist")
	}
}