package main

import "fmt"

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog. These nutz"

	aNumber := 42
	isTrue := true

	strln, _ := fmt.Println(str1, str2, str3)

	//if err == nil {
	fmt.Println("The length is:", strln)
	//}

	fmt.Printf("Value of a number: %v\n", aNumber)
	fmt.Printf("Value isTrue: %v\n", isTrue)
	fmt.Printf("Value of a float: %0.2f\n", float64(aNumber))

	fmt.Printf("Data Types %T,  %T, %T, %T and %T\n",
		str1, str2, str3, isTrue, aNumber)

	outText := fmt.Sprintf("val Data Types %T,  %T, %T, %T and %T\n",
		str1, str2, str3, isTrue, aNumber)

	fmt.Println(outText)
}
