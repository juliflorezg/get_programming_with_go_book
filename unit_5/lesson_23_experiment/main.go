package main

import (
	"fmt"
	"math"
)

type location struct {
	name      string
	lat, long float64
}

type world struct {
	radius float64 // in km
}

type gps struct {
	current     location
	destination location
	world
}
type rover struct {
	gps
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

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	d := fmt.Sprint("There're still ", g.distance(), " km to destination")
	return d
}

func (l location) description() string {
	lat := fmt.Sprintf("%f", l.lat)
	long := fmt.Sprintf("%f", l.long)

	return l.name + " latitude: " + lat + " longitude: " + long
}

func main() {
	mars := world{radius: 3389.5}
	bradbury := location{lat: -4.5895, long: 137.4417, name: "Bradbury Landing"}
	elysium := location{lat: 4.5, long: 135.9, name: "Elysium Planitia"}
	marsGPS := gps{world: mars, current: bradbury, destination: elysium}

	curiosity := rover{marsGPS}

	fmt.Println(bradbury.description())
	fmt.Println(elysium.description())

	fmt.Println(curiosity.message())
}
