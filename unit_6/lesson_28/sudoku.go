package main

import "errors"

const rows, columns = 9, 9

var (
	ErrBounds = errors.New("The position is out of bounds")
	ErrDigit  = errors.New("The value is not a valid digit")
)

type grid [rows][columns]int8

func (g *grid) set(r, c int, value int8) error {
	if !inBounds(r, c) {
		return ErrBounds
	}

  if !isValidDigit(value){
    return ErrDigit
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
