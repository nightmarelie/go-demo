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

func calculateIMT(height float64, weight float64) (result float64) {
	const IMT_POWER float64 = 2

	result = weight / math.Pow(height, IMT_POWER)

	return
}

func getUserInput() (userHeight float64, userKg float64) {

	fmt.Println("Enter your height (m):")
	fmt.Scan(&userHeight)
	fmt.Println("Enter your weight (kg):")
	fmt.Scan(&userKg)

	return
}
