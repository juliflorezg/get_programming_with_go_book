package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// join strings together
	countdownStr := "Launch is in T minus" + "10 seconds"
	// can't join string with number
	// countdownStr := "Launch is in T minus" + 10 + "seconds" //! invalid operation: "Launch is in T minus" + 10 (mismatched types untyped string and untyped int)

	fmt.Println(countdownStr)

	// type conversion
	age := 26
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge = marsAge * earthDays / marsDays
	fmt.Println("I am", marsAge, "years old on Mars.")

	fmt.Println(int(earthDays)) // 365

	fmt.Println(math.MinInt16) // -32768
	fmt.Println(math.MaxInt16) // 32767

	// string conversions
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Print(string(pi), string(alpha), string(omega), string(bang), "\n") // πάω!

	// use strcon.Itoa function from string conversion package to convert numbers to strings
	countdown := 10
	str := "Launch in T minus " + strconv.Itoa(countdown) + " seconds"
	fmt.Println(str)
	countdown = 9
	str = fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str)

	// now the other way, using the atoi function from strconv

	countdown, err := strconv.Atoi("11")
	if err != nil {
		// log error, stop app...
	}
	fmt.Println(countdown)

	//boolean conversion
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)

	// for i := 0; i < 6; i++ {
	// }
	fmt.Println(input())

}
