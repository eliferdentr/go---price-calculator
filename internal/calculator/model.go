package calculator

type PriceCalculator struct {
	Prices [] float64 `json:"prices"`
	TaxRates []float64 `json:"tax_rates"`
}

type Validator interface {
	Validate () error
}

type Calculator interface{
	Calculate() map[float64][] float64
}