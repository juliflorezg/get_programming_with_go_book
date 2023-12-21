package main

import (
	"fmt"
	"image"
	"log"
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

func NewRoverDriver() *RoverDriver {
	rd := &RoverDriver{
		commandc: make(chan command),
	}

	go rd.drive()

	return rd
}

func (rd *RoverDriver) drive() {
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
				pos = pos.Add(direction)
				log.Printf("moved to %v", pos)
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
