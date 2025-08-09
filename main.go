package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.78
	userKg := 85.5

	fmt.Println("Calculate your IMT")
	fmt.Println("Enter your height (m):")
	fmt.Scan(&userHeight)
	fmt.Println("Enter your weight (kg):")
	fmt.Scan(&userKg)

	imt := calculateIMT(userHeight, userKg)
	outputResult(imt)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Your IMT is: %f\n", imt)

	fmt.Println(result)
}

func calculateIMT(height float64, weight float64) float64 {
	const IMT_POWER float64 = 2

	return weight / math.Pow(height, IMT_POWER)
}
