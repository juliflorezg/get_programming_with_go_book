package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

type getRowFn func(row int) (string, string)

func drawTable(rows int, gr getRowFn) {}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func measureTemperature(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Printf("%v °K\n", k)
		time.Sleep(time.Second)
	}
}

var f = func() {
	fmt.Println("Dress up for the masquerade.")
}

// * closures; anonymous function encloses the value of s and offset even after calibrate fn returns
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {

	sensor := fakeSensor
	fmt.Println(sensor())
	// sensor = realSensor
	// fmt.Println(sensor())

	//> we can pass a function as a parameter to another function
	measureTemperature(3, sensor)

	//* anonymous function
	// can be assigned to a variable
	f()

	// we can declare them and invoke in one step
	func() {
		fmt.Println("Functions anonymous")
	}()

	var offset kelvin = 5

	sensor = calibrate(realSensor, offset)
	fmt.Println(sensor())
	offset = 6.9
	fmt.Println(sensor())

	sensor2 := calibrate(fakeSensor, offset)
	fmt.Println(sensor2())
	fmt.Println(sensor2())
	fmt.Println(sensor2())

}
