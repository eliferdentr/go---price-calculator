package utils

import "fmt"

func HandleError (err error, context string) {
	if err != nil {
		fmt.Printf("Error in %s: %s \n", context, err)
		panic(err)
	}
}