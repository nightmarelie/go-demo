package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight = 70.0
	var userKg float64 = 86
	var IMT = userKg / math.Pow(userHeight, 2)

	fmt.Println(IMT)
}
