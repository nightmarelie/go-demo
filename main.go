package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight, userKg := getUserInput()

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

func getUserInput() (float64, float64) {
	var userHeight float64
	var userKg float64

	fmt.Println("Enter your height (m):")
	fmt.Scan(&userHeight)
	fmt.Println("Enter your weight (kg):")
	fmt.Scan(&userKg)

	return userHeight, userKg
}
