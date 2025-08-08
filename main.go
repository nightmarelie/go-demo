package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight = 1.78
	var userKg float64 = 85.5
	var IMT = userKg / math.Pow(userHeight, 2)

	fmt.Println(IMT)
}
