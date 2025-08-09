package main

import (
	"fmt"
	"math"
)

func main() {
	const IMP_POWER float64 = 2
	userHeight := 1.78
	userKg := 85.5

	fmt.Println("Enter your height:")
	fmt.Scan(&userHeight)
	fmt.Println("Enter your weight:")
	fmt.Scan(&userKg)

	IMT := userKg / math.Pow(userHeight, IMP_POWER)

	fmt.Println(IMT)
}
