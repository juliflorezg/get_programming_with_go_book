// lesson goals:

// learn how to use pointers
// understand relation between pointers and RAM (random access memory)
// know when to use (and not use) pointers

package main

import (
	"fmt"
	"time"
)

// * pointers for structures
type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	//we don't need to dereference the value like (*p).age++
	p.age++
}

func (p *person) birthday1() {
	p.age++
}

// interior pointers
type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name string
	stats
}

func reclassify(planets *[]string){
	*planets = (*planets)[0:8]
}

func main() {
	// use & (address operator to get memory address of a value)

	answer := 69

	fmt.Println(&answer)
	// fmt.Println(&42) // ! error: invalid operation: cannot take address of 42 (untyped int constant)
	// fmt.Println(&"we can't use & for getting memory address or literals (string, number, boolean)") //! error

	// dereferencing (getting the value of a memory address)

	address := &answer
	fmt.Println(*address)

	// type of pointer
	fmt.Printf("type of address: %T\n", address) // *int

	// another example
	{
		canada := "Canada"

		var home *string
		fmt.Printf("type of home: %T\n", home) // *string

		home = &canada
		// derefence a pointer
		fmt.Println(*home) // Canada
	}

	//> NASA administrator example
	var administrator *string
	scolese := "Christopher J. Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	// Scolese was succeeded by Bolden in 2009

	bolden := "Charles F. Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	//changes made to Bolden will be reflected on administrator
	bolden = "Charles Frank Bolden Jr."
	fmt.Println(*administrator)

	// also possible the other way round (change variable value through dereferencing)
	*administrator = "Maj. Gen. Charles Frank Bolden Jr."
	fmt.Println(bolden)

	// if we assign administrator to other variable, it will also point to the underlying value (bolden in this case)
	major := administrator
	*major = "Major General Charles Frank Bolden Jr."
	fmt.Println(bolden)

	//since administrator & major are pointing to the same value in the same memory address, they are the same (for now)
	fmt.Println(administrator == major) // true

	// Bolden was succeeded by Lightfoot in 2017
	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot
	// admin is pointing to other value now, bc of that admin & major are no longer the same
	fmt.Println(administrator == major) // false

	// when we assign a dereferenced value to another variable, it creates a copy of that value so changes made to the original doesn't affect the clone

	charles := *major
	*major = "Charles Bolden"

	fmt.Println(charles)
	fmt.Println(*major)

	// two strings are the same if the value is the same, even if they have different memory addresses

	charles = "Charles Bolden"
	fmt.Println(charles == bolden)   // true
	fmt.Println(&charles == &bolden) // false

	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	// we can use . syntax as usual since go automatically dereference the value so, instead of having to write (*timmy).superpower we can write the more readable timmy.superpower

	// (*timmy).superpower = "flying" //!!!
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy) // &{name:Timothy superpower:Flying age:10}

	//* pointing to arrays
	// same as with structs
	superpowers := &[3]string{"flying", "invisibility", "super strength"}

	fmt.Println(superpowers[2])
	fmt.Println(superpowers[1:2])

	birthday(timmy)
	fmt.Printf("%+v\n", timmy) // &{name:Timothy superpower:Flying age:11}

	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)
	fmt.Printf("%+v\n", rebecca) // {name:Rebecca superpower:imagination age:15}

	//* using pointer receivers
	terry := &person{
		name:       "Terry",
		superpower: "Super intelligence",
		age:        15,
	}

	terry.birthday1()

	fmt.Printf("%+v\n", terry)

	// we can also call birthday1 method from a person type, go will take the memory address and dereference it inside the function
	nathan := person{
		name: "Nathan",
		age:  17,
	}

	nathan.birthday1()
	fmt.Printf("%+v\n", nathan)

	// time.Time doesn't use pointer receivers

	{
		const layout = "Mon, Jan 2, 2006"
		day := time.Now()
		tomorrow := day.Add(24 * time.Hour)
		fmt.Println(day.Format(layout))
		fmt.Println(tomorrow.Format(layout))
	}

	// interior pointers
	{
		player := character{
			name: "Lol",
		}

		fmt.Printf("%+v\n", player.stats)
		levelUp(&player.stats) // interior pointer, this one is pointing to the memory address of field 'stats' inside player structure (character)
		fmt.Printf("%+v\n", player.stats)
	}

	//* maps are pointers themselves, they act as pointers when passed to functions or when assigned, so pointing to a map is redundant, the following is incorrect
	// func demolish (planets *map[string]string)

	// slices also behave like pointers, bc they point to an underlying array, so when a slice is passed to a function, the pointer allows to change or modify the array

	// if a function wants to modify any of these two types: structures and arrays, they must be passed a pointer to that value

	// we can use a pointer to a slice if we want to modify the slice capacity, length or starting offset
	{
		planets := []string{
			"Mercury", "Venus", "Earth", "Mars",
			"Jupiter", "Saturn", "Uranus", "Neptune",
			"Pluto",
		}

		reclassify(&planets)
		fmt.Println(planets)
	}


	// turtle exercise 
	t := turtle{
		x: 10,
		y: 10,
	}

	up(&t)
	left(&t)
	left(&t)
	left(&t)
	up(&t)
	up(&t)
	up(&t)
	down(&t)
	right(&t)
	right(&t)
	right(&t)
	right(&t)
	down(&t)
	right(&t)
	fmt.Printf("%+v", t)


}
