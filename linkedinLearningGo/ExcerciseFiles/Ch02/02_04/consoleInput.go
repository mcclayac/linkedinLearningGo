package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//var s string
	//fmt.Scan(&s)
	//fmt.Println(s)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text :")
	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	myString := fmt.Sprint(str)
	fmt.Printf("myString = %s \n", myString)

	fmt.Print("Enter Number :")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err == nil {
		fmt.Printf("the value entered : %0.2f", f)
	}
}
