package main

import (
	"fmt"
	"strings"
)

func main() {
	myString := "An implicitly created variable"
	fmt.Printf("My String and Type %s:%T\n", myString, myString)

	var myString2 string
	myString2 = "An explicitly created variable"
	fmt.Printf("My String2 and Type %s:%T\n", myString2, myString2)
	fmt.Println(strings.ToUpper(myString2))
	fmt.Println(strings.Title(myString2))

	lValue := "hello"
	uValue := "HELLO"

	fmt.Println("Equal?", (lValue == uValue))
	fmt.Println("Equal non-case sensitive?", strings.EqualFold(lValue, uValue))
	fmt.Println("Contains 'exp'?", strings.Contains(myString, "exp"))
	fmt.Println("Contains 'exp'?", strings.Contains(myString2, "exp"))

}
