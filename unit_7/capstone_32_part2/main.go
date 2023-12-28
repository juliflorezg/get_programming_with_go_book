package main

import (
	"fmt"
	"image"
	"math/rand"
	"sync"
	"time"
)

func main() {
	marsToEarth := make(chan []Message)

	go earthReceiver(marsToEarth)
	gridSize := image.Point{X: 25, Y: 25}
	grid := NewMarsGrid(gridSize)
	rover := make([]*RoverDriver, 5)
	for i := range rover {
		rover[i] = startDriver(fmt.Sprint("rover", i), grid, marsToEarth)
	}
	time.Sleep(60 * time.Second)
}

const (
	// The length of a Mars day.
	dayLength = 24 * time.Second
	// The length of time per day during which messages can be transmitted from a rover to Earth.
	receiveTimePerDay = 2 * time.Second
)

// MarsGrid represents a grid of some of the surface of Mars. It may be used concurrently by different goroutines.
type MarsGrid struct {
	bounds image.Rectangle
	mu     sync.Mutex
	cells  [][]cell
}

// SensorData holds information about what's in a point in the grid.
type SensorData struct {
	LifeSigns int
}
type cell struct {
	groundData SensorData
	occupier   *Occupier
}

// NewMarsGrid returns a new MarsGrid of the given size.
func NewMarsGrid(size image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{
			Max: size,
		},
		cells: make([][]cell, size.Y),
	}
	for y := range grid.cells {
		grid.cells[y] = make([]cell, size.X)
		for x := range grid.cells[y] {
			cell := &grid.cells[y][x]
			cell.groundData.LifeSigns = rand.Intn(1000) + 1
		}
	}
	return grid
}

// Size returns a Point representing the size of the grid.
func (g *MarsGrid) Size() image.Point {
	return g.bounds.Max
}

// Occupy occupies a cell at the given point in the grid. It returns nil if the point is already occupied or the point is outside the grid. Otherwise it returns a value that can be used to move to different places on the grid.
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	cell := g.cell(p)
	if cell == nil || cell.occupier != nil {
		return nil
	}
	cell.occupier = &Occupier{
		grid: g,
		pos:  p,
	}
	return cell.occupier
}

func (g *MarsGrid) cell(p image.Point) *cell {
	if !p.In(g.bounds) {
		return nil
	}
	return &g.cells[p.Y][p.X]
}

// Occupier represents an occupied cell in the grid.
type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

// MoveTo moves the occupier to a different cell in the grid.
// It reports whether the move was successful
// It might fail because it was trying to move outside the grid or because the cell it's trying to move into is occupied. If it fails, the occupier remains in the same place.
func (o *Occupier) MoveTo(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	newCell := o.grid.cell(p)
	if newCell == nil || newCell.occupier != nil {
		return false
	}
	o.grid.cell(o.pos).occupier = nil
	newCell.occupier = o
	o.pos = p
	return true
}

// Sense returns sensory data from the current cell.
func (o *Occupier) Sense() SensorData {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	return o.grid.cell(o.pos).groundData
}

// Pos returns the current grid position of the occupier.
func (o *Occupier) Pos() image.Point {
	return o.pos
}
