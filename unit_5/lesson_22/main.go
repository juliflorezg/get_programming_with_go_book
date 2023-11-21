package main

import (
	"fmt"
	"math"
)

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere. (4°35'22.2" S, 137°26'30.12" E)
type coordinate struct {
	d, m, s float64
	h       rune // 'N', 'n', 'S', 's', 'E', 'e', 'W', 'w'
}

// -14.568388888888888 175.4726388888889
type location struct {
	lat, long float64
}

type world struct {
	radius float64 // km
}

// decimal returns a decimal coordinate for a dms coordinate
func (c coordinate) decimal() float64 {
	sign := 1.0

	switch c.h {
	case 'S', 's', 'W', 'w':
		sign = -1
	}

	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func (w world) distance(l1, l2 location) float64 {
	s1, c1 := math.Sincos(rad(l1.lat))
	s2, c2 := math.Sincos(rad(l2.lat))
	clong := math.Cos(rad(l1.long - l2.long))

	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radians (360° are 2π radians)
func rad(deg float64) float64 {
	return ((deg * math.Pi) / 180)
}

func main() {
	//Bradbury landing: 4°35'22.2" S, 137°26'30.12" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())

	bradburyLocation := location{lat.decimal(), long.decimal()}
	fmt.Println(bradburyLocation)

	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity)

	//world type
	var mars = world{radius: 3389.5}

	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}

	dist := mars.distance(spirit, opportunity)

	fmt.Printf("%.2f km\n", dist)

	getLocations()
	getDistances()
	getDistances2()
}
