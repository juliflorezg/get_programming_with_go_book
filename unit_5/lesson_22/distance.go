package main

import "fmt"

func getDistances2() {
	earth := world{6371.0}
	mars := world{3389.5}

	// Find the distance from London, England (51°30’N 0°08’W)
	// to Paris, France(48°51’N 2°21’E).

	londonLocation := newLocation(coordinate{51, 30, 0, 'N'}, coordinate{0, 8, 0, 'W'})

	parisLocation := newLocation(coordinate{48, 51, 0, 'N'}, coordinate{2, 21, 0, 'E'})

	londonParisDistance := earth.distance(londonLocation, parisLocation)

	// Find the distance from your city to the capital of your country.
	bogotaLocation := location{4.624335, -74.063644}

	tocancipaLocation := location{4.96531, -73.91301}

	tocBogDistance := earth.distance(bogotaLocation, tocancipaLocation)

	// Find the distance between Mount Sharp (5°4' 48"S, 137°51’E) and
	// Olympus Mons (18°39’N, 226°12’E) on Mars.

	mountSharp := newLocation(coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0, 'E'})
	olympusMons := newLocation(coordinate{18, 39, 0, 'N'}, coordinate{226, 12, 0, 'E'})
	sharpMonsDistance := mars.distance(mountSharp, olympusMons)

	fmt.Printf("Distance from London to Paris is: %.2f km\n", londonParisDistance)
	fmt.Printf("Distance from Bogotá to Tocancipá is: %.2f km\n", tocBogDistance)
	fmt.Printf("Distance from Mount Sharp to Olympus Mons on Mars is: %.2f km\n", sharpMonsDistance)

}
