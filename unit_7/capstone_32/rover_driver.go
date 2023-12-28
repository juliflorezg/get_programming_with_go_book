package main

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"time"
)

type command int

const (
	right = command(0)
	left  = command(1)
	start = command(100)
	stop  = command(300)
)

type RoverDriver struct {
	commandc chan command
}

func NewRoverDriver(mg *MarsGrid) *RoverDriver {
	rd := &RoverDriver{
		commandc: make(chan command),
	}

	go rd.drive(mg)

	return rd
}

func (rd *RoverDriver) drive(mg *MarsGrid) {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}

	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	isMoving := false

	for {
		select {
		case c := <-rd.commandc:
			fmt.Println("c:", c)
			fmt.Println("isMoving:", isMoving)
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			case start:
				isMoving = true
			case stop:
				isMoving = false
			}
			log.Printf("New direction %v", direction)
		case <-nextMove:
			fmt.Println(isMoving)
			if isMoving {
				nextPosition := pos.Add(direction)
				isNextPosFree := true
				// out of bounds
				if nextPosition.X < mg.firstLimit.X || nextPosition.X > mg.secondLimit.X || nextPosition.Y < mg.firstLimit.Y || nextPosition.Y > mg.secondLimit.Y {
					randDir := rand.Intn(2)
					if randDir == int(left) {
						rd.Left()
					} else if randDir == int(right) {
						rd.Right()
					}
				}
				// if position that rover is about to move in is already occupied
				for i := 0; i < len(mg.occupiedPositions); i++ {
					occ := &mg.occupiedPositions[i]
					occ.mu.Lock()
					defer occ.mu.Unlock()
					if occ.pos.Eq(nextPosition) {
						randDir := rand.Intn(2)
						if randDir == int(left) {
							rd.Left()
							isNextPosFree = false
							break
						} else if randDir == int(right) {
							rd.Right()
							isNextPosFree = false
							break
						}
					}
				}
				if isNextPosFree {

					pos = nextPosition
					log.Printf("moved to %v", pos)
				} else {
					log.Printf("Next position is occupied, changed direction")
				}
			}
			nextMove = time.After(updateInterval)
		}
	}
}

// Left turns the rover left (90° counter clockwise)
func (rd *RoverDriver) Left() {
	rd.commandc <- left
}

// Left turns the rover right (90° clockwise)
func (rd *RoverDriver) Right() {
	rd.commandc <- right
}

func (rd *RoverDriver) Start() {
	rd.commandc <- start
}

func (rd *RoverDriver) Stop() {
	rd.commandc <- stop
}
func worker() {
	// n := 0
	pos := image.Point{X: 10, Y: 10}
	direction := image.Point{X: 1, Y: 0}
	delay := 1.0
	next := time.After(time.Duration(delay) * time.Second)

	for {
		select {
		case ct := <-next:
			// n++
			// fmt.Println(n)
			pos = pos.Add(direction)
			fmt.Printf("%v\n", ct.Second())
			fmt.Println("current position is ", pos)
			delay += 0.5
			next = time.After(time.Duration(delay) * time.Second)
		}
	}
}
