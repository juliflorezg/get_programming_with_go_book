package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const lightSpeed = 299_792  // km/s
	const spaceXSpeed = 100_800 // km/h
	const hoursPerDay, minutesPerHour = 24, 60
	var distance = 56_000_000                   // km
	fmt.Println(distance/lightSpeed, "seconds") //s
	distance = 401_000_000
	fmt.Println(distance/lightSpeed, "seconds") //s
	distance = 96_300_000
	fmt.Println((distance / spaceXSpeed / hoursPerDay), "days") //s
	fmt.Println(minutesPerHour)
	var weight = 100
	weight -= 2
	fmt.Println(weight)

	// randomness

	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)

	/* The distance between Earth and Mars varies from nearby to opposite sides
	of the sun. Write a program that generates a random distance from 56,000,000 to
	401,000,000 km.
	*/
	{
		const nearbyDistance = 56_000_000
		const oppositeSidesDistance = 401_000_000
		var num = rand.Intn(oppositeSidesDistance-nearbyDistance) + nearbyDistance

		fmt.Println(num)
	}

	//malacandra
	malacandra()

}
