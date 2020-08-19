package main

import "fmt"

func main() {

	doSomething()

	result := addValues(3, 5)
	fmt.Println("add values result :", result)

	result = addValues2(3, 12)
	fmt.Println("add values 2 result :", result)

	result = addLongValues(1, 3, 5, 7)
	fmt.Println("add longValues result :", result)

}

func addValues(value1 int, value2 int) int {

	return value1 + value2
}

func addValues2(value1, value2 int) int {

	return value1 + value2
}
func doSomething() {

	fmt.Println("Inside do something")
}

func addLongValues(values ...int) int {
	sum := 0
	fmt.Printf("The type of values : %T\n", values)
	for i := range values {
		sum = sum + values[i]
	}
	return sum
}
