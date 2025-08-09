package main

import (
	"fmt"
	"math"
)

func main() {
	const IMP_POWER float64 = 2
	userHeight := 1.78
	userKg := 85.5
	IMT := userKg / math.Pow(userHeight, IMP_POWER)

	fmt.Println(IMT)
}
