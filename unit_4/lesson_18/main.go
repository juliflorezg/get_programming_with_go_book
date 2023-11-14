package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range newWorlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}

	return newWorlds
}

// appending, length & capacity in slices
func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)
	fmt.Println(len(dwarfs))

	dump("dwarfs:", dwarfs)
	dump("dwarfs[1:2]:", dwarfs[1:2])

	// using append and investigating capacity

	dwarfs1 := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println("dwarfs1", len(dwarfs1), cap(dwarfs1))

	dwarfs2 := append(dwarfs1, "Orcus")
	fmt.Println("dwarfs2", len(dwarfs2), cap(dwarfs2))

	dwarfs3 := append(dwarfs2, "Salacia", "Quaoar", "Sedna")
	fmt.Println("dwarfs3", len(dwarfs3), cap(dwarfs3))

	dwarfs3[0] = "random"

	//* the change is only reflected on dwarfs3 & dwarfs2 since they point to another array, due to the increased capacity
	fmt.Println(dwarfs1) // [Ceres Pluto Haumea Makemake Eris]
	fmt.Println(dwarfs2) // [random Pluto Haumea Makemake Eris Orcus]
	fmt.Println(dwarfs3) // [random Pluto Haumea Makemake Eris Orcus Salacia Quaoar Sedna]

	//> the three index slicing
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	// third index is for specifying the slice capacity, if it's not provided, attempts to append to this new slice will overwrite the original slice (see block in line  54)
	terrestrial := planets[0:4:4]
	fmt.Println("terrestrial", len(terrestrial), cap(terrestrial), terrestrial)

	worlds := append(terrestrial, "Ceres")
	fmt.Println(worlds)
	fmt.Println(planets)

	//! wrong, without 3rd index
	{
		fmt.Println("-----wrong, without 3rd index!!!")
		planets := []string{
			"Mercury", "Venus", "Earth", "Mars",
			"Jupiter", "Saturn", "Uranus", "Neptune",
		}
		terrestrial := planets[0:4]
		fmt.Println("terrestrial", len(terrestrial), cap(terrestrial), terrestrial) // 4, 8

		worlds := append(terrestrial, "Ceres")
		fmt.Println(worlds)
		// Ceres overwrites Jupiter in underlying array
		fmt.Println(planets) // [Mercury Venus Earth Mars Ceres Saturn Uranus Neptune]
	}

	//* allocating capacity using make function

	{
		dwarfs := make([]string, 0, 10)
		dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
		fmt.Println("dwarfs", len(dwarfs), cap(dwarfs), dwarfs)

	}

	// working with variadic functions
	{
		planets := []string{"Venus", "Mars", "Jupiter"}
		newPlanets := terraform("New", planets...)

		fmt.Println(newPlanets)
	}

	slice := []string{"First one"}
	newSlice := capacity(slice, 20)
	fmt.Println(newSlice)
}
