package main

import (
	"fmt"
	"os"
)

func main() {
	s, fv := NewSudoku([rows][columns]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	for _, v := range s {
		fmt.Println(v)
	}

	fmt.Println(fv)

	// err := s.set(0, 0, 4, fv)
	err := s.Clear(6, 6, fv)
	// err := s.set(7, 7, 2, fv)
	if err != nil {
		// fmt.Println(err)

		if errs, ok := err.(SudokuError); ok {
			// fmt.Println(errs)
			fmt.Printf("%d error(s) occurred: \n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}

		os.Exit(1)
	}

	for _, v := range s {
		fmt.Println(v)
	}
}
