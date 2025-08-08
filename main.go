package main

import "fmt"

func main() {
	var userHeight = 70.0
	var userKg float64 = 86
	var IMT = userKg / userHeight

	fmt.Println(IMT)
}
