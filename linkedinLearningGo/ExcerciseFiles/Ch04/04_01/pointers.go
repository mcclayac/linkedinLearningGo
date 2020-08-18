package main

import "fmt"

func main() {
	var p *int

	if p != nil {
		fmt.Println("Write the value of p:", *p)
	} else {
		fmt.Println("The String is nil")
	}

	var v int = 42
	//# Connect the pointer to this variable
	p = &v
	if p != nil {
		fmt.Println("Write the value of p:", *p)
	} else {
		fmt.Println("The String is nil")
	}

	var value1 float64 = 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value 1:", *pointer1)
	fmt.Println("Value 1:", value1)
	fmt.Printf("Value : %.3f\n", value1)
	fmt.Printf("Value : %.3f\n", *pointer1)

}
