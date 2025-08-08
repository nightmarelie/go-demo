package main

import "fmt"

func main() {
	var userHeight = 70.0
	var userKg = 86
	var IMT = float64(userKg) / userHeight

	fmt.Println(IMT)
}
