package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	var map1 map[string]int
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
		// "Moon":  0,
	}
	fmt.Println(map1)
	fmt.Println(temperature)
	temp := temperature["Earth"]
	fmt.Printf("On average, the Earth is %v °C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

	moon := temperature["Moon"]
	fmt.Println(moon) // 0, since that key doesn't exist on temperature map

	if moon, ok := temperature["Moon"]; ok {
		fmt.Println(ok)
		fmt.Printf("On average the moon is %vº C.\n", moon)
	} else {
		fmt.Println(ok)
		fmt.Println("Where the hell is the Moon?")
	}

	// when we pass a map we don't pass a copy ::
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	planetsMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets)       // map[Earth:whoops Mars:Sector ZZ9]
	fmt.Println(planetsMarkII) // map[Earth:whoops Mars:Sector ZZ9]

	delete(planets, "Earth")
	fmt.Println(planetsMarkII) // map[Mars:Sector ZZ9]

	//* using make for declare a map
	temperature2 := make(map[string]int, 8)
	fmt.Println(temperature2)

	// using a map to count things

	temperatures := []float64{-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0}

	frequency := make(map[float64]int)

	for _, t := range temperatures {
		frequency[t]++
	}

	for t, count := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, count)
	}

	//grouping data with maps & slices

	groups := make(map[float64][]float64)

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10

		groups[g] = append(groups[g], t)
	}

	fmt.Println(groups)

	for g, temps := range groups {
		//g is for group of temperatures
		fmt.Printf("%v: %v\n", g, temps)
	}

	// making a set from a map

	set := make(map[float64]bool)

	for _, t := range temperatures {
		set[t] = true
	}

	fmt.Println(set)

	// set to slice and sorting

	unique := make([]float64, 0, len(set))

	for k := range set {
		unique = append(unique, k)
	}

	// sort.Float64s(unique)
	slices.Sort(unique)
	fmt.Println(unique) // [-33 -31 -29 -28 -23 32]

  words()
}
