package main

import "fmt"

func main() {

	// x float64 = 0
	var result string

	if x := -42; x < 0 {
		result = "Less than 0"
	} else if x == 0 {
		result = "equal to 0"
	} else {
		result = "Greater than 0"
	}

	fmt.Println(result)
	// x is not in scope
	//fmt.Printf("Value of x: %v", x)
}
