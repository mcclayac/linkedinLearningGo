package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2009, time.November,
		10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launch date at : %s\n", t)

	now := time.Now()
	fmt.Printf("The time now : %s\n", now)
	fmt.Println("The month is :", now.Month())
	fmt.Println("The Day is :", now.Day())
	fmt.Println("The Week Day is :", now.Weekday())

	tomarrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomarrow is %v, %v, %v, %v\n",
		tomarrow.Weekday(), tomarrow.Month(), tomarrow.Day(), tomarrow.Year())

	longFormat := "Monday, January 2, 2006"
	fmt.Println("Tomarrow format is:", tomarrow.Format(longFormat))
	shortFormat := "01/2/06"
	fmt.Println("Tomarrow short format is:", tomarrow.Format(shortFormat))

}
