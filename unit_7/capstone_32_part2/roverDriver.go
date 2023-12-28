package main

import (
	"image"
	"log"
	"math/rand"
	"time"
)

// Message holds a message as sent from Mars to Earth.
type Message struct {
	Pos       image.Point
	LifeSigns int
	Rover     string
}

type command int

const (
	right command = 0
	left  command = 1
)

func startDriver(name string, grid *MarsGrid, marsToEarth chan []Message) *RoverDriver {
	var o *Occupier
	// Try a random point; continue until we've found one that's not currently occupied.
	for o == nil {
		startPoint := image.Point{X: rand.Intn(grid.Size().X), Y: rand.Intn(grid.Size().Y)}
		o = grid.Occupy(startPoint)
	}
	return NewRoverDriver(name, o, marsToEarth)
}

// Radio represents a radio transmitter that can send
// message to Earth.
type Radio struct {
	fromRover chan Message
}

// SendToEarth sends a message to Earth. It always succeeds immediately - the actual message may be buffered and actually transmitted later.
func (r *Radio) SendToEarth(m Message) {
	r.fromRover <- m
}

// NewRadio returns a new Radio instance that sends messages on the toEarth channel.
func NewRadio(toEarth chan []Message) *Radio {
	r := &Radio{
		fromRover: make(chan Message),
	}
	go r.run(toEarth)
	return r
}

// run buffers messages sent by a rover until they can be sent to Earth.
func (r *Radio) run(toEarth chan []Message) {
	var buffered []Message
	for {
		toEarth1 := toEarth
		if len(buffered) == 0 {
			toEarth1 = nil
		}
		select {
		case m := <-r.fromRover:
			buffered = append(buffered, m)
		case toEarth1 <- buffered:
			buffered = nil
		}
	}
}

// RoverDriver drives a rover around the surface of Mars.
type RoverDriver struct {
	commandc chan command
	occupier *Occupier
	name     string
	radio    *Radio
}

// NewRoverDriver starts a new RoverDriver and returns it.
func NewRoverDriver(name string, occupier *Occupier, marsToEarth chan []Message) *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
		occupier: occupier,
		name:     name,
		radio:    NewRadio(marsToEarth),
	}
	go r.drive()
	return r
}

// drive is responsible for driving the rover. It is expected to be started in a goroutine.
func (r *RoverDriver) drive() {
	log.Printf("%s initial position %v", r.name, r.occupier.Pos())
	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
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
			}
			log.Printf("%s new direction %v", r.name, direction)
		case <-nextMove:
			nextMove = time.After(updateInterval)
			newPos := r.occupier.Pos().Add(direction)
			if r.occupier.MoveTo(newPos) {
				log.Printf("%s moved to %v", r.name, newPos)
				r.checkForLife()
				break
			}
			log.Printf("%s blocked trying to move from %v to %v", r.name, r.occupier.Pos(), newPos)
			// Pick one of the other directions randomly.
			// Next time round, we'll try to move in the new direction.
			dir := rand.Intn(3) + 1 // 1-4
			for i := 0; i < dir; i++ {
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("%s new random direction %v", r.name, direction)
		}
	}
}
func (r *RoverDriver) checkForLife() {
	// Successfully moved to new position.
	sensorData := r.occupier.Sense()
	if sensorData.LifeSigns < 900 {
		return
	}
	r.radio.SendToEarth(Message{
		Pos:       r.occupier.Pos(),
		LifeSigns: sensorData.LifeSigns,
		Rover:     r.name,
	})
}

// Left turns the rover left (90° counterclockwise).
func (r *RoverDriver) Left() {
	r.commandc <- left
}

// Right turns the rover right (90° clockwise).
func (r *RoverDriver) Right() {
	r.commandc <- right
}
