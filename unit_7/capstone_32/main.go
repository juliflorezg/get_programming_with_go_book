package main

import (
	"image"
	"sync"
)

type MarsGrid struct {
	mu                      sync.Mutex
	firstLimit, secondLimit image.Point
	occupiedPositions       []Occupier
}

func (mg *MarsGrid) Occupy(p image.Point) *Occupier {
	mg.mu.Lock()
	defer mg.mu.Unlock()

	// out of bounds
	if p.X < mg.firstLimit.X || p.Y > mg.secondLimit.Y {
		return nil
	}

	//! for _, occ := range mg.occupiedPositions {
	// point is already occupied
	for i := 0; i < len(mg.occupiedPositions); i++ {
		occ := &mg.occupiedPositions[i]
		occ.mu.Lock()
		defer occ.mu.Unlock()
		if occ.pos.Eq(p) {
			return nil
		}
	}
	// newOccPoint := Occupier{
	// 	mu:  sync.Mutex{},
	// 	pos: p,
	// }
	// newOccPoint.mu.Lock()
	// defer newOccPoint.mu.Unlock()

	mg.occupiedPositions = append(mg.occupiedPositions, Occupier{
		mg:  mg,
		mu:  sync.Mutex{},
		pos: p,
	})

	return &mg.occupiedPositions[len(mg.occupiedPositions)-1]
}

type Occupier struct {
	mg  *MarsGrid
	mu  sync.Mutex
	pos image.Point
}

func (o *Occupier) Move(p image.Point) bool {
	if p.X < o.mg.firstLimit.X || p.Y > o.mg.secondLimit.Y {
		return false
	}

	for i := 0; i < len(o.mg.occupiedPositions); i++ {
		occ := &o.mg.occupiedPositions[i]
		occ.mu.Lock()
		defer occ.mu.Unlock()
		if occ.pos.Eq(p) {
			return false
		}
	}

	o.pos.X = p.X
	o.pos.Y = p.Y
	return true
}
