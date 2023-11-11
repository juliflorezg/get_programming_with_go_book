package main

import "fmt"

func terraform(arr [8]string) [8]string {
	for i, v := range arr {
		arr[i] = "New " + v
	}
	return arr
}

func main() {

	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2]

	fmt.Println(earth)

	fmt.Println(len(planets))
	fmt.Println(planets[5] == "")

	// planets[8] = "Pluto" //! invalid argument: index 8 out of bounds

	// i := 8
	// planets[i] = "Pluto" //! panic: runtime error: index out of range [8] with length 8
	// pluto := planets[i]

	// fmt.Println(pluto)

	//>>> composite literal
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	planets = [...]string{"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Println(dwarfs)
	fmt.Println(len(planets))

	//two ways to loop
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}
	fmt.Println()

	for i, v := range dwarfs {
		fmt.Println(i, v)
	}

	//* arrays are copied
	planetsMarkII := planets

	planets[2] = "whoops"
	fmt.Println(planets)
	fmt.Println(planetsMarkII)
	planets = terraform(planets)

	fmt.Println(planets)
	// fmt.Println(terraform(planets))

	getInitialPositions()
	displayBoard()

}
