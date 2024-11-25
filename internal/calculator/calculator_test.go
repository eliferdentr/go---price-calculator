package calculator

import (
    "testing"
)

func TestValidate(t *testing.T) {
    pc := PriceCalculator{
        Prices: [] float64{100, 200},
        TaxRates: [] float64{10, 20},
    }

    err := pc.Validate()
    if err != nil {
        t.Errorf("Unexpected error: %s", err)
    }

    pc.Prices = []float64{-100, 100}
    err = pc.Validate()
    if err == nil {
        t.Errorf("Expected error for negative price, got nil")
    }
}

func TestCalculate(t *testing.T){
    pc := PriceCalculator{
        Prices: []float64{100, 200},
        TaxRates: []float64{10},
    }

    result := pc.Calculate()

    if len(result[10]) != 2{
        t.Errorf("Expected 2 results, got %d", len(result[10]))
    }

    if result[10][0] != 110 {
		t.Errorf("Expected 110, got %f", result[10][0])
	}
}