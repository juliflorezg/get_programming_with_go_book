package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("You leave the cave:", exit)

	fmt.Println("There is a sign near the entrance that reads 'No Minors'.")
	var age = 26
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)

	fmt.Println("banana is greater than apple? ", "apple" < "banana")

	// if statement
	command = "go east"
	if command == "go east" {
		fmt.Println("You head further up the mountain.")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}

	// determine if 2100 will be a leap year

	var year = 2100
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	if leap {
		fmt.Println("Leap yearrrr")
	} else {
		fmt.Println("NO leap yearr :(")

	}

	fmt.Println("There is a cavern entrance here and a path to the east.")
	command = "go inside"
	switch command {
	case "go east":
		fmt.Println("You head further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}

	// use of fallthrough keyword
	var room = "lake"

	switch room {
	case "cave":
		fmt.Println("Didn't quite get that.")
	case "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough
	case "underwater":
		fmt.Println("The water is freezing cold.")
	}

	// for loop

	// var count = 5
	// for count > 0 {
	// 	fmt.Println(count)
	// 	time.Sleep(time.Second)
	// 	count--
	// }
	// fmt.Println("Liftoff!")

	// for i := 0; i < 50; i++ {
	// 	fmt.Println(rand.Intn(2))
	// }

	//infinite loop with break
	var degrees = 0

	for {
		fmt.Println("Degrees:::", degrees)
		degrees++
		if degrees >= 360 {
			degrees = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}

	//* Quick check 3.6 Not every launch goes smoothly. Implement a countdown where every second there’s a 1 in 100 chance that the launch fails and the countdown stops.

	// count = 10
	// for count > 0 {
	// 	fmt.Println(count)
	// 	if (rand.Intn(99) + 1) == 1 {
	// 		break
	// 	}
	// 	time.Sleep(time.Second)
	// 	count--
	// }
	// if count == 0 {
	// 	fmt.Println("Another liftoff!!!")
	// } else {
	// 	fmt.Println("Launch failed :c")

	// }

	guess()

}
