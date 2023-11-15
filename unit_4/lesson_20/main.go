package main

import (
	"fmt"
	"math/rand"
)

type Universe [][]bool

func NewUniverse() Universe {
	newUniverse := make(Universe, height)

	for i := range newUniverse {
		newUniverse[i] = make([]bool, width)
	}

	return newUniverse
}

func (u Universe) Show() {
	// r is for row in the universe
	for _, r := range u {
		// c is for each individual cell
		for _, c := range r {
			if c {
				fmt.Print("*")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}

func (u Universe) Seed() {
	// get random percentage from 23 to 27%
	percentage := 23 + rand.Intn(5)
	fmt.Println(percentage)
	var numOfCells float64 = float64(width*height) * (float64(percentage) / 100)
	fmt.Println(numOfCells)

	//count goes from 0 to numOfCells
	count := 0

	// firstLoop:
	// for _, r := range u {
	// 	for j := range r {
	// 		if count == int(numOfCells) {
	// 			break firstLoop
	// 		}

	// 		willLive := rand.Intn(2)
	// 		if willLive == 1 {
	// 			r[j] = true
	// 			count++
	// 		}
	// 	}
	// }

firstLoop:
	for i := 0; i < len(u)/2; i++ {
		for j := 0; j < width; j++ {
			if count == int(numOfCells) {
				break firstLoop
			}
			fmt.Println(i, j)
			fmt.Println("len", len(u[i]))

			willLive := rand.Intn(2)
			if willLive == 1 {
				u[i][j] = true
				count++
			}
			willLive = rand.Intn(2)
			if willLive == 1 {
				// fmt.Println(i, j)
				fmt.Println(len(u) - 1 - i)
				fmt.Println(len(u[i]) - 1 - j)
				u[len(u)-1-i][len(u[i])-1-j] = true
				count++
			}
		}
	}

}

func (u Universe) Alive(x, y int) bool {
	horizontal, vertical := x, y

	if x < 0 {
		horizontal = width + x
	}
	if x > 79 {
		horizontal = x % width
	}
	if y < 0 {
		vertical = height + y
	}
	if y > 14 {
		vertical = y % height
	}

	return u[horizontal][vertical]

}

const (
	width  = 80
	height = 15
)

func main() {
	u := NewUniverse()
	u.Show()
	u.Seed()
	u.Show()
	fmt.Println(u[0][3])
	fmt.Println(u.Alive(0, 3))
	fmt.Println(u.Alive(80, 3))
	fmt.Println(u.Alive(83, 3))
}
