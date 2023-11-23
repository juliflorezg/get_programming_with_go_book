// for full information about this exercise, refer to lesson 25 in book (pg 196)
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type sanctuary struct {
	animals []animal
}

// type animal interface {
// 	move() string
// 	eat() string
// }

type animal struct {
	name  string
	foods []string
}

// func String makes animal satisfy the Stringer interface
func (a animal) String() string {
	return fmt.Sprintf("name:::%v", a.name)
}

// type squirrel struct {
// 	animalInfo
// }

func (a animal) move() string {
	sound := ""
	switch a.name {
	case "squirrel":
		sound = "wirr birr"
	case "cow":
		sound = "moo neigh"
	case "dog":
		sound = "woof roar"
	case "cat":
		sound = "meow hiss"
	}
	return strings.Repeat(sound+" ", rand.Intn(4)+1)
}
func (a animal) eat() string {
	return a.foods[rand.Intn(len(a.foods))]
}

// type cow struct {
// 	animalInfo
// }

// func (c cow) move() string {
// 	return strings.Repeat("moo, neigh", rand.Intn(4)+1)
// }
// func (c cow) eat() string {
// 	return c.foods[rand.Intn(len(c.foods))]
// }

// type dog struct {
// 	animalInfo
// }

// func (d dog) move() string {
// 	return strings.Repeat("woof, woof", rand.Intn(4)+1)
// }
// func (d dog) eat() string {
// 	return d.foods[rand.Intn(len(d.foods))]
// }

// type cat struct {
// 	animalInfo
// }

// func (c cat) move() string {
// 	return strings.Repeat("meow, hiss", rand.Intn(4)+1)
// }
// func (c cat) eat() string {
// 	return c.foods[rand.Intn(len(c.foods))]
// }

func (s sanctuary) dayCycle() {
	for i := 0; i < 24; i++ {
		fmt.Print(fmt.Sprintf("%2d:00", i) + " ")
		action := rand.Intn(2)
		switch i {
		case 0, 1, 2, 3, 4, 5, 18, 19, 20, 21, 22, 23:
			{
				fmt.Print("The animals are sleeping zzz\n")
			}
		default:
			{
				a := s.animals[rand.Intn(len(s.animals))]
				switch action {
				case 0:
					{
						fmt.Printf("%v is moving:: %v\n", a.name, a.move())
					}
				case 1:
					{
						fmt.Printf("%v is eating:: %v\n", a.name, a.eat())
					}
				}
			}
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {

	a := []animal{
		{
			name:  "squirrel",
			foods: []string{"nuts", "seeds", "roots"},
		},
		{
			name:  "cow",
			foods: []string{"pasture", "silage", "corn"},
		},
		{
			name:  "dog",
			foods: []string{"steak", "strawberry", "chicken"},
		},

		{
			name:  "dog",
			foods: []string{"beef", "tuna", "turkey"},
		},
	}

	s := sanctuary{
		animals: a,
	}

	s.dayCycle()
	s.dayCycle()
	s.dayCycle()
}
