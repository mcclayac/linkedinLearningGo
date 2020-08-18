package main

import "fmt"

func main() {
	//Declare Array
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[1])

	// Declare array and initialize values
	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)

	// Length of arrays
	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))

}
