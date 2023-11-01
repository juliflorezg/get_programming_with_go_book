package main

import "fmt"

func malacandra() {

	const distance = 56_000_000                    // km/h
	var expectedTripDuration, hoursPerDay = 28, 24 // days, h

	var speedToGetThere float64 = float64(distance) / float64((expectedTripDuration * hoursPerDay))
	fmt.Println(speedToGetThere, "km/h")

}
