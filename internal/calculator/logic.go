package calculator

import "fmt"

type PriceModifier func(float64) float64

/* 
	* Checks the validity og the input
*/
func (pc *PriceCalculator) Validate() error {
	if len(pc.Prices) == 0 || len(pc.TaxRates) == 0 {
		return fmt.Errorf("Prices or Tax Rates cannot be empty.")
	}

	for _, price := range pc.Prices {
		if price < 0 {
			return fmt.Errorf("There cannot be a negative price: %v", price)
		}
	}

	for _, tax := range pc.TaxRates {
		if tax < 0 {
			return fmt.Errorf("There cannot be a negative tax rate: %v", tax)
		}
	}
	return nil
}

/*
	* Computes the tax-inclusive prices
*/

func (pc *PriceCalculator) Calculate() map[float64][]float64 {
    result := make(map[float64][]float64)
    for _, tax := range pc.TaxRates {
        for _, price := range pc.Prices {
            newPrice := price + (price * tax / 100)
            // Check if the key exists, initialize if not
            if _, exists := result[tax]; !exists {
                result[tax] = []float64{}
            }
            result[tax] = append(result[tax], newPrice)
        }
    }
    return result
}

/* 
 *ApplyModifier applies a price modifier function to all prices
*/

func (pc *PriceCalculator) ApplyModifier (modifier PriceModifier) {
	for i, price := range pc.Prices {
		pc.Prices[i] = modifier(price)
	}
}