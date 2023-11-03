package main

import (
	"fmt"
	"math"
)

func main() {
	var days float64 = 365.2425
	var answer float64 = 42.0
	fmt.Println(days, answer)

	var pi32 float32 = math.Pi
	var pi64 = math.Pi

	fmt.Println(pi32)
	fmt.Println(pi64)

	// zero value
	var price float64
	fmt.Println(price)

	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)     // 0.3333333333333333
	fmt.Printf("%f\n", third)     // 0.3333333333333333
	fmt.Printf("%.3f\n", third)   // 0.333
	fmt.Printf("%4.2f\n", third)  // 0.33
	fmt.Printf("%06.3f\n", third) // 00.333
	fmt.Printf("%06.2f\n", third) // 000.33
	fmt.Printf("%06.1f\n", third) // 0000.3

	// floating point inaccuracies
	third = 1.0 / 3.0
	fmt.Println(third + third + third) // 1 wtf

	piggyBank := 0.1
	piggyBank += 0.2

	fmt.Println(piggyBank) // 0.30000000000000004

	// best results by multiplicate first instead of divide first
	celsius := 21.0
	fmt.Println((celsius/5.0*9.0)+32, "°F") // 69.80000000000001 °F
	fmt.Println((9.0/5.0*celsius)+32, "°F") // 69.80000000000001 °F

	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Println(fahrenheit, "°F")                                                           //69.8 °F
	fmt.Println(0.10 + 0.10 + 0.10 + 0.10 + 0.10 + 0.10 + 0.10 + 0.10 + 0.10 + 0.10 + 0.10) // 1.1

	piggyBank = 0.0
	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}
	fmt.Println(piggyBank) // 1.0999999999999999

	piggy()

}
