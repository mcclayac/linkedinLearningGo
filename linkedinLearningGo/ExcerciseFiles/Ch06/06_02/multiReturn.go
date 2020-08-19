package main

import "fmt"

func main() {

	nm, le := FullName("Tony", "McClay")

	fmt.Printf("full name :%v length: %v", nm, le)
}

func FullName(f, l string) (string, int) {

	full := f + " " + l
	length := len(full)
	return full, length
}
