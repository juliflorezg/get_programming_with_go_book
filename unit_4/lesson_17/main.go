package main

import (
	"fmt"
	"sort"
	"strings"
)

func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

// slices
func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(planets)
	fmt.Println(terrestrial, gasGiants, iceGiants)

	//slices can be accessed using the index like an array
	fmt.Println(terrestrial[3])

	// we can slice an array an then slice the resulting slice
	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]

	fmt.Println(giants, gas, ice)

	//> when we modify an element present in one sub-slice, the change will be reflected on other slices and the original array (if any) where that element is present

	iceGiantsMarkII := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println(planets)
	fmt.Println(iceGiants, ice, iceGiantsMarkII)

	//default indices por slicing
	terrestrial = planets[:4] // from 0 to 4 (not including 4)
	gasGiants = planets[4:6]
	iceGiants = planets[6:] // from 6 to length of planets (last element)

	allPlanets := planets[:]
	fmt.Println(allPlanets)

	//* slicing also works on strings
	{
		neptune := "Neptune"
		tune := neptune[3:]
		fmt.Println(tune) // tune

		//when assigning other value to the original doesn't change the sliced string
		neptune = "Poseidon"
		fmt.Println(tune) // tune

		// index indicate the number of bytes, not runes
		question := "¿Cómo estás?"
		fmt.Println(question[:6]) // ¿Cóm

		// declaring slices
		dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

		dwarfSlice := dwarfArray[:]
		fmt.Println(dwarfSlice)

		dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

		fmt.Printf("type of dwarfsArray: %T\n", dwarfArray)
		fmt.Printf("type of dwarfs slice: %T\n", dwarfs)

	}

	planets2 := []string{"  Venus   ", "Earth   ", " Mars"}
	fmt.Println(planets2)

	// removing space surrounding worlds
	hyperspace(planets2)
	fmt.Println(planets2)
	fmt.Println(strings.Join(planets2, ""))
	// sorting a slice
	fmt.Println("sorting a slice")
	fmt.Println(planets)
	sort.StringSlice(planets[:]).Sort()
	fmt.Println(planets)

	// even simpler
	sort.Strings(planets[:])

	//* to terraform
	samplePlanets := []string{"Mars", "Uranus", "Neptune"}
	terraform(samplePlanets)
	fmt.Println(samplePlanets)

	samplePlanets2 := Planets{"Mercury", "Jupiter", "Saturn"}
	samplePlanets2.terraform()
	fmt.Println(samplePlanets2)
}
