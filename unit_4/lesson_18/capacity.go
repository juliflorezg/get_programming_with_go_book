package main

import "fmt"

func capacity(s []string, count int) []string {
	for i := 0; i < count; i++ {
		c := cap(s)
		s = append(s, "New"+string(i))
		if c != cap(s) {
			fmt.Println("capacity changed!!!")
			fmt.Printf("length: %v, capacity: %v\n", len(s), cap(s))
		}
	}

	return s

}
