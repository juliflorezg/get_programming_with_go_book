package main

import (
	"errors"
	"strings"
)

const rows, columns = 9, 9

var (
	ErrBounds = errors.New("The position is out of bounds")
	ErrDigit  = errors.New("The value is not a valid digit")
)

type grid [rows][columns]int8

type SudokuError []error

func (se SudokuError) Error() string{
  var s []string

  for _, err := range se {
    s = append(s, err.Error())
  }

  return strings.Join(s, ", ")
}

func (g *grid) set(r, c int, value int8) error {
  var errs SudokuError
	if !inBounds(r, c) {
    errs = append(errs, ErrBounds)
		// return ErrBounds
	}
  
  if !isValidDigit(value){
    errs = append(errs, ErrDigit)
    // return ErrDigit
  }

  if len(errs) > 0 {
    return errs
  }

	g[r][c] = value
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
