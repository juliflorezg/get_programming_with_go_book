// Canis Major Dwarf is the closest known galaxy to Earth at 236,000,000,000,000,000 km from our Sun (though some dispute that it is a galaxy). Use constants to convert this distance to light years

package main

import "fmt"

func canis() {
	const distanceFromEarthToCanis = 236_000_000_000_000_000 // km
	// todo get equivalence from 1km to light year

	const lightSpeed = 299_972 //km/s
	// todo get how many km does light travel a year
	const secondsPerDay = 86400
	const secondsPerYear = secondsPerDay * 365

	const lightYear = lightSpeed * secondsPerYear // 9_459_916_992_000

	fmt.Println(lightYear)
	fmt.Println("this is the distance from earth to canis galaxy in light years:", distanceFromEarthToCanis/lightYear)
}
