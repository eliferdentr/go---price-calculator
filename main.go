package main

import (
	"fmt"
	"pricecalculator.com/internal/calculator"
	"pricecalculator.com/internal/file"
    "pricecalculator.com/internal/utils"
)

func main () {
	fmt.Println("Price calculator is starting...")

	//read the input file
	var pc calculator.PriceCalculator
	err := file.ValidateJSONFormat("input.json", &pc)
	utils.HandleError(err, "There was an error while reading the file.")

	//check the format
	fmt.Println("Validating the input...")
	err = pc.Validate()
	utils.HandleError(err, "File format is wrong")

	//do the calculation
	fmt.Println("Doing the calculation...")
	results := pc.Calculate ()

	//write the results to file
	err = file.WriteJSONFile("output.json", results)
	utils.HandleError(err, "There was an error while writing the results to file")

	fmt.Println("Calculation is done. Results can be found in 'output.json' file...")


}