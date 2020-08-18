package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	//	Add Values to a slice
	colors = append(colors, "Purple")
	fmt.Println(colors)

	// remove the first itm from an array
	// start with color index = 1
	// to the end of the array len(colors)
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	// the default :[value] is the len of the array
	colors = append(colors[1:])
	fmt.Println(colors)

	// Remove the last item of the slice
	// colors -1
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// Declare Slice type int, iitial size 5, grow by 5
	numbers := make([]int, 5, 5)
	numbers[0] = 126
	numbers[1] = 75
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156

	fmt.Println(numbers)

	// Append a number to slice
	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println("Capacity :", cap(numbers))

	// sort the slice
	sort.Ints(numbers)
	fmt.Println(numbers)

}
