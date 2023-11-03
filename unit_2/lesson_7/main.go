package main

import (
	"fmt"
	"time"
)

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	days := 365.2425
	// use the first argument for the second format verb
	fmt.Printf("Type %T for %[1]v\n", days)
	animal := "lion"
	fmt.Printf("Type %T for %[1]v\n", animal)
	third := 1.0 / 3.0
	fmt.Printf("Type %T for %[1]v\n", third)

	fmt.Printf("Type %T for %[1]v\n", true)

	//uint8 for colors in css (red, green, blue)

	var red, green, blue uint8 = 0, 141, 213
	//* go prefixes hexadecimal numbers with 0x, the line below is equivalente to the one previous to this
	// var red, green, blue uint8 = 0x00, 0x8d, 0xd5
	fmt.Println(red, green, blue)
	fmt.Printf("%x %X %x\n", red, green, blue)
	fmt.Printf("color: #%02x%02x%02x;\n", red, green, blue)

	//> integers wrap around
	// when an integer rage is exceeded, it goes to the beginning of the range
	{
		var red uint8 = 255
		red++
		fmt.Println(red) // 0
		var number int8 = 127
		number++
		fmt.Println(number) // -128
		var number2 uint16 = 65535
		number2++
		fmt.Println(number2) // 0

	}

	//avoid wrapping with time
	{
		future := time.Unix(12_622_780_800, 0)
		fmt.Println(future)
	}

	piggy()
}
