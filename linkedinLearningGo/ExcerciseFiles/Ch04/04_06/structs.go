package main

import (
	"fmt"
)

type Dog struct {
	Breed  string
	weight int
}

func main() {
	poodle := Dog{"Poodle", 13}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed :%v\nWeight :%v\n", poodle.Breed, poodle.weight)

}
