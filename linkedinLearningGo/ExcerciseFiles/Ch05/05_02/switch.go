package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	//dow := rand.Intn(6) + 1
	//// rand.Intn value 0-6
	//fmt.Println("Day", dow)

	result := ""
	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "It's a Sunday"
	case 7:
		result = "It's a Saturday"
	default:
		result = "It's a weekday"
	}
	//fmt.Println("dow :", dow, "Result : ", result)
	fmt.Println("Result : ", result)

	// case inside the switch
	switch x := -42; {
	case x < 0:
		result = "less than 0"
		//fallthrough
	case x > 0:
		result = "greater than 0"
	default:
		result = "equal to 0"
	}
	fmt.Println("Result : ", result)
}
