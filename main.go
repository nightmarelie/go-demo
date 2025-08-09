package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64 = 1.78
	userKg := 85.5
	IMT := userKg / math.Pow(userHeight, 2)

	fmt.Println(IMT)
}
