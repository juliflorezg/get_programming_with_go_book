package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)
		count++
	}
	// fmt.Println(num) // ! error

	//short declaration for variables in for loop, count is not accessible out of the for loop scope
	{
		for count1 := 10; count1 > 0; count1-- {
			fmt.Println(count1)
		}
		// fmt.Println(count1) // ! error
	}

	//short declaration inside an if statement

	if num := rand.Intn(3) + 1; num == 1 {
		fmt.Println("Space Adventures")
	} else if num == 2 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}

	// fmt.Println(num) //! error, num is no longer in scope

	// short declaration inside a switch statement
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}

	// Experiment: dandom-dates.go
	generateRandomDates()

}
