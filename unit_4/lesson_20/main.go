package main

import (
	"fmt"
	"math/rand"
	"time"
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
	fmt.Println("\x0c")

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
	fmt.Println("percentage: ", percentage)
	var numOfCells float64 = float64(width*height) * (float64(percentage) / 100)
	fmt.Println("living cells::", numOfCells)

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

	for count <= int(numOfCells) {
		vertical := rand.Intn(height)
		horizontal := rand.Intn(width)
		if u[vertical][horizontal] {
			continue
		}
		u[vertical][horizontal] = true
		count++
	}

	// for count <= int(numOfCells) {

	// 	for i := 0; i < len(u)/2; i++ {
	// 		for j := 0; j < width; j++ {

	// 			// fmt.Println(i, j)
	// 			// fmt.Println("len", len(u[i]))

	// 			// decide if cell in first half will live
	// 			willLive := rand.Intn(2)
	// 			if willLive == 1 {
	// 				u[i][j] = true
	// 				count++
	// 			}
	// 			// decide if cell in second half will live
	// 			willLive = rand.Intn(2)
	// 			if willLive == 1 {
	// 				// fmt.Println(i, j)
	// 				// fmt.Println(len(u) - 1 - i)
	// 				// fmt.Println(len(u[i]) - 1 - j)

	// 				u[len(u)-1-i][len(u[i])-1-j] = true
	// 				count++
	// 			}
	// 		}
	// 	}
	// }

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

	// fmt.Println("horizontal", horizontal)
	// fmt.Println("vertical", vertical)
	// fmt.Println("horizontal", horizontal)

	return u[vertical][horizontal]

}

// function Neighbors return the amount of alive neighbors for a specified cell
func (u Universe) Neighbors(x, y int) int {
	count := 0
	horizontal, vertical := x, y

	// look from upper left to lower right adjacent positions to the cell
	for v := -1; v <= 1; v++ {
		for h := -1; h <= 1; h++ {
			// if it's not the cell in question and it's alive
			// if (v != 0 && h != 0) && u.Alive(horizontal+h, vertical+v) {
			if !(v == 0 && h == 0) && u.Alive(horizontal+h, vertical+v) {
				count++
			}
		}
	}

	return count
}

// Next determines if a cell must live the next generation
func (u Universe) Next(x, y int) bool {

	// isAlive := u.Alive(x, y)
	numOfNeighbors := u.Neighbors(x, y)
	// // A live cell with less than two live neighbors dies.
	// if isAlive && numOfNeighbors < 2 {
	// 	return false
	// }
	// // A live cell with two or three live neighbors lives on to the next generation
	// if isAlive && (numOfNeighbors == 2 || numOfNeighbors == 3) {
	// 	return true
	// }
	// // A dead cell with exactly three live neighbors becomes a live cell.
	// if !isAlive && numOfNeighbors == 3 {
	// 	return true
	// }
	// // A live cell with more than three live neighbors dies.
	// return false

	return numOfNeighbors == 3 || numOfNeighbors == 2 && u.Alive(x, y)

}

// Step goes on one universe and for each cell, defines it's state for the next gen, it does it in another sample Universe, so the current one doesn't get affected by changes
func Step(a, b Universe) {
	for v := range a { // vertical, rows
		for h := range a[v] { // horizontal, columns
			shouldLive := a.Next(h, v)
			b[v][h] = shouldLive
		}
	}

}

const (
	width  = 80
	height = 15
)

func main() {
	a, b := NewUniverse(), NewUniverse()
	// a := NewUniverse()
	// u.Show()
	a.Seed()
	// a.Show()
	fmt.Println("Is currently alive? (5, 2):", a.Alive(5, 2))
	fmt.Println("neighbors (5, 2):", a.Neighbors(5, 2))
	fmt.Println("Should Live Next Gen? (5, 2):", a.Next(5, 2))

	count := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if a[i][j] {
				count++
			}
		}
	}
	fmt.Println("count of live Cells:", count)
	// a.Next()
	
	for i := 0; i < 150; i++ {
		Step(a, b)
		a.Show()
		time.Sleep(time.Second / 20)
		a, b = b, a
	}
	fmt.Println("count of live Cells:", count)
	
	count = 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if a[i][j] {
				count++
			}
		}
	}
	fmt.Println("count of live Cells on last generation:", count)
	// u.Show()
	// fmt.Println(u[0][3])
	// fmt.Println(u.Alive(0, 3))
	// fmt.Println(u.Alive(80, 3))
	// fmt.Println(u.Alive(83, 3))
}
