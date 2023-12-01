package main

import (
	"errors"
	"fmt"
	"strings"
)

const rows, columns = 9, 9

var (
	ErrBounds     = errors.New("The position is out of bounds")
	ErrDigit      = errors.New("The value is not a valid digit")
	ErrRow        = errors.New("The value is already present in the specified row")
	ErrColumn     = errors.New("The value is already present in the specified column")
	ErrSubgrid    = errors.New("The value is already present in the corresponding sub-grid")
	ErrFixedValue = errors.New("The value in this position is an initial value provided by the program and can't be overwritten/cleared")
)

type FixedValues map[string]bool

type SudokuGrid [rows][columns]int8

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string

	for _, err := range se {
		s = append(s, err.Error()) // -> converts error to string
	}

	return strings.Join(s, ", ")
}

func (g *SudokuGrid) set(r, c int, value int8, fv FixedValues) error {
	var errs SudokuError
	if !inBounds(r, c) {
		errs = append(errs, ErrBounds)
	}

	if !isValidDigit(value) {
		errs = append(errs, ErrDigit)
	}

	//check if digit is in row
	if isInRow(g, r, value) {
		errs = append(errs, ErrRow)
	}
	//check if digit is in column
	if isInColumn(g, c, value) {
		errs = append(errs, ErrColumn)
	}
	//check if digit is in subgrid (3x3 subregion)
	if isInSubgrid(g, r, c, value) {
		errs = append(errs, ErrSubgrid)
	}

	// check if position is a fixed value from grid initialization
	if isFixedValuePos(r, c, fv) {
		errs = append(errs, ErrFixedValue)
	}

	if len(errs) > 0 {
		return errs
	}

	g[r][c] = value
	return nil
}

func (g *SudokuGrid) Clear(r, c int, fv FixedValues) error {
	var errors SudokuError

	if !inBounds(r, c) {
		errors = append(errors, ErrBounds)
	}
	if isFixedValuePos(r, c, fv) {
		errors = append(errors, ErrFixedValue)
	}

	if len(errors) > 0 {
		return errors
	}

	g[r][c] = 0

	return nil
}

func inBounds(r, c int) bool {
	if r < 0 || r >= rows {
		return false
	}
	if c < 0 || c >= columns {
		return false
	}

	return true
}

func isValidDigit(v int8) bool {
	return v >= 1 && v <= columns
}

func isInRow(g *SudokuGrid, r int, v int8) bool {
	for i := 0; i < columns; i++ {
		if g[r][i] == v {
			return true
		}
	}

	return false
}

func isInColumn(g *SudokuGrid, c int, v int8) bool {
	for i := 0; i < rows; i++ {
		if g[i][c] == v {

			return true
		}
	}

	return false
}

func isInSubgrid(g *SudokuGrid, r, c int, value int8) bool {
	// // todo determine the corresponding subgrid for that position

	var rowStart, rowEnd, colStart, colEnd int
	// switch -> 9 subgrid based on r & c

	fmt.Println(r, c)

	switch true {
	// 1st subgrid:
	case (r >= 0 && r <= 2) && (c >= 0 && c <= 2):
		rowStart = 0
		rowEnd = 2
		colStart = 0
		colEnd = 2
		// 2nd subgrid:
	case (r >= 0 && r <= 2) && (c >= 3 && c <= 5):
		rowStart = 0
		rowEnd = 2
		colStart = 3
		colEnd = 5
		// 3rd subgrid:
	case (r >= 0 && r <= 2) && (c >= 6 && c <= 8):
		rowStart = 0
		rowEnd = 2
		colStart = 6
		colEnd = 8

		// 4th subgrid:
	case (r >= 3 && r <= 5) && (c >= 0 && c <= 2):
		rowStart = 3
		rowEnd = 5
		colStart = 0
		colEnd = 2
		// 5th subgrid:
	case (r >= 3 && r <= 5) && (c >= 3 && c <= 5):
		rowStart = 3
		rowEnd = 5
		colStart = 3
		colEnd = 5
		// 6th subgrid:
	case (r >= 3 && r <= 5) && (c >= 6 && c <= 8):
		rowStart = 3
		rowEnd = 5
		colStart = 6
		colEnd = 8

		// 7th subgrid:
	case (r >= 6 && r <= 8) && (c >= 0 && c <= 2):
		rowStart = 6
		rowEnd = 8
		colStart = 0
		colEnd = 2
		// 8th subgrid:
	case (r >= 6 && r <= 8) && (c >= 3 && c <= 5):
		rowStart = 6
		rowEnd = 8
		colStart = 3
		colEnd = 5
		// 9th subgrid:
	case (r >= 6 && r <= 8) && (c >= 6 && c <= 8):
		rowStart = 6
		rowEnd = 8
		colStart = 6
		colEnd = 8
	}

	fmt.Printf("rowStart:%v, rowEnd:%v\n", rowStart, rowEnd)
	fmt.Printf("colStart:%v, colEnd:%v\n", colStart, colEnd)
	for i := rowStart; i <= rowEnd; i++ {
		for j := colStart; j <= colEnd; j++ {
			fmt.Printf("i:%v, j:%v\n", i, j)
			fmt.Printf("value in grid:%v, value to set:%v\n", g[i][j], value)
			if g[i][j] == value {
				return true
			}
		}
	}
	return false
}

func isFixedValuePos(r, c int, fv FixedValues) bool {
	for k := range fv {
		pos := fmt.Sprintf("%v,%v", r, c)

		// if position given by set/clear functions is already in fixedValues map, can't be set/cleared
		if k == pos {
			return true
		}
	}

	return false
}

func NewSudoku(g SudokuGrid) (SudokuGrid, FixedValues) {
	// // todo: initialize the grid with the given values, ie:
	// s := NewSudoku([rows][columns]int8{
	// 	{5, 3, 0, 0, 7, 0, 0, 0, 0},
	// 	{6, 0, 0, 1, 9, 5, 0, 0, 0},
	// 	{0, 9, 8, 0, 0, 0, 0, 6, 0},
	// 	{8, 0, 0, 0, 6, 0, 0, 0, 3},
	// 	{4, 0, 0, 8, 0, 3, 0, 0, 1},
	// 	{7, 0, 0, 0, 2, 0, 0, 0, 6},
	// 	{0, 6, 0, 0, 0, 0, 2, 8, 0},
	// 	{0, 0, 0, 4, 1, 9, 0, 0, 5},
	// 	{0, 0, 0, 0, 8, 0, 0, 7, 9},
	// })

	var initialGrid SudokuGrid
	var fv = make(FixedValues)

	for i := range initialGrid {
		for j := 0; j < len(initialGrid[i]); j++ {
			// initialGrid[i][j] = g[i][j]
			// verify that digit is between 0 and 9

			initialGrid.set(i, j, g[i][j], fv)

			// checks if its a fixed value and store it in fv(fixedValues map)
			if g[i][j] > 0 {
				mapValue := fmt.Sprintf("%v,%v", i, j)
				fmt.Println(mapValue)
				fv[mapValue] = true
			}
		}
	}

	return g, fv

}
