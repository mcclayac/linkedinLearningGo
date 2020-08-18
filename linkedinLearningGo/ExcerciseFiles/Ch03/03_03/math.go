package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	r1, r2, r3 := 1, 3, 5
	intSum := i1 + i2 + i3
	intSumr := r1 + r2 + r3
	fmt.Println("Integer sum: ", intSum)
	fmt.Println("Integer sumR: ", intSumr)
	f1, f2, f3 := 23.32, 23.88, 55.22
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(23.45)
	b2.SetFloat64(44.6432)
	b3.SetFloat64(102.23456)
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("BigSum = %10g\n", &bigSum)

	circleRadium := 15.5
	circleCircumference := circleRadium * math.Pi
	fmt.Printf("Circumfrance : %.2f\n", circleCircumference)

}
