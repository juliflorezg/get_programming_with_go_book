// nil, lesson goals:

// Do something with nothing
// Understand the trouble with nil
// See how Go improves on nil’s story

package main

import (
	"fmt"
	"sort"
)

type person struct {
	age int
}

func (p *person) birthday() {
	//guarding against nil pointer
	if p == nil {
		return
	}
	p.age++
}

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		// less = func(i, j int) bool { return s[i] < s[j] }
		// Write a line of code to sort food from the shortest to longest string in listing 27.6.
		less = func(i, j int) bool { return len(s[i]) < len(s[j]) }
	}
	sort.Slice(s, less)
}

// "Whenever you write a function that accepts a slice, ensure that a nil slice has the same behavior as an empty slice."
func mirepoix(s []string) []string {
	return append(s, "celery", "carrot", "onion")

}
func main() {
	// nil means nothing or zero
	// it's the zero value for pointers, slices, maps & interfaces

	// nil leads to panic
	var nowhere *int
	fmt.Println(nowhere) // <nil>
	// fmt.Println(*nowhere) // ! panic: runtime error: invalid memory address or nil pointer dereference

	if nowhere != nil {
		fmt.Println(*nowhere)

	}

	var nobody *person
	fmt.Println(nobody) // <nil>

	nobody.birthday()

	// nil function values
	var fn func(a, b int) int
	fmt.Println(fn == nil) // true

	// fn(1, 4) // ! panic: runtime error: invalid memory address or nil pointer dereference
	food := []string{"onion", "carrot", "celery"}

	sortStrings(food, nil)
	fmt.Println(food)

	// nil slices
	var soup []string
	fmt.Println(soup == nil) // true
	for _, v := range soup {
		fmt.Println(v)
	}
	fmt.Println(len(soup)) // 0
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(len(soup)) // 3
	fmt.Println(soup)      // 3

	//functions can accept nil if they expect a slice
	soup2 := mirepoix(nil)
	fmt.Println(soup2)
	knights()

}
