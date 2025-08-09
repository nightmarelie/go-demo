package main

import (
	"fmt"
	"math"
)

func main() {
	const IMP_POWER float64 = 2
	userHeight := 1.78
	userKg := 85.5

	fmt.Println("Calculate your IMT")
	fmt.Println("Enter your height:")
	fmt.Scan(&userHeight)
	fmt.Println("Enter your weight:")
	fmt.Scan(&userKg)

	imt := userKg / math.Pow(userHeight, IMP_POWER)

	fmt.Println("Your IMT is:", imt)
}
