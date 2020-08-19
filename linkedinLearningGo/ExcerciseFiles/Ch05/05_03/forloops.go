package main

import "fmt"

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	for i := 0; i < len(colors); i++ {
		fmt.Printf("Color is : %v\n", colors[i])
	}

	for i := range colors {
		fmt.Printf("Range Color is : %v\n", colors[i])
	}

	// while loop using for
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
	}

}
