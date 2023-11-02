package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%-17v%-5v%-5v%-6v%v", "Spaceline", "Days", "Trip", "type", "Price\n")
	fmt.Println("======================================")

	//todo iterate 10 times for 10 random ticket pricing
	// [X] get random spaceline
	//todo calculate maximum amount of days bases on distance and minimum ship speed
	//todo get random value for trip type
	//todo associate speed with pricing, for every km/s of speed increase 1 mill dollars in the provided range
	//todo
	//todo

	spacelines := map[int]string{
		0: "Space Aventures",
		1: "SpaceX",
		2: "Virgin Galactic",
	}

	tripTypes := map[int]string{
		0: "One-way",
		1: "Round-trip",
	}

	for i := 0; i < 30; i++ {

		spaceline := spacelines[rand.Intn(3)]
		tripType := tripTypes[rand.Intn(2)]
		distance := 62_100_000 // km
		minimumSpeed := 16     //km/s
		maximumSpeed := 30     //km/s
		secondsInADay := 86400
		tripSpeed := rand.Intn(maximumSpeed-minimumSpeed+1) + minimumSpeed
		tripDuration := (distance / tripSpeed) / secondsInADay // days
		tripPrice := tripSpeed + 20

		if tripType == "Round-trip" {
			tripPrice *= 2
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline, tripDuration, tripType, tripPrice)
	}
}
