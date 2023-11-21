package main

import "fmt"

//? composition and forwarding

// naive attempt of representing a weather report
// type report struct {
// 	sol       int
// 	high, low float64 //* high & low what, it is not clear what this fields are referencing to
// 	lat, long float64
//   //! it gets worse when adding other fields like wind speed and direction, pressure, humidity, season, sunrise & sunset
// }

type report struct {
	sol         int // current day
	temperature temperature
	location    location
}

type sol int

//> embed a type in a struct (makes method forwarding possible as well as accessing the inner type fields)
// we can embed any type in structures, not only other structures (embeding sol type)
type report1 struct {
	sol
	temperature
	location
}

type temperature struct {
	high, low celsius
}

type celsius float64

type location struct {
	lat  float64
	long float64
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

//method that forwards average to report
func (r report) average() celsius {
	return r.temperature.average()
}

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}

	return days
}

// name collisions
func (l location) days(l2 location) int {
  // todo: complicated distance calculation (lesson 22)
  return 5
}

func (r report1) days(s2 sol) int {
  return r.sol.days(s2)
}

func main() {

	// create a weather report, composing it from related data
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}

	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)

	fmt.Printf("A balmy %v°C \n", report.temperature.high)

	fmt.Printf("average %v°C \n", t.average())

	// we can access average temperature by chaining it from the report
	fmt.Printf("average temperature from report: %v°C\n", report.temperature.average())

	report1 := report1{
		sol:         15,
		temperature: temperature{high: -1.0, low: -78.0},
		location:    location{-4.5895, 137.4417},
	}
	fmt.Printf("average (forwarding): %v°C \n", report1.average())
	// we can access the inner structures fields
	fmt.Printf("%v°C \n", report1.high)
	// we can also modify inner struct fields
	report1.high = 32

	fmt.Println(report1.temperature.high)
	fmt.Println(report1.sol)
	fmt.Println(report1.sol.days(1446)) //* this line and the one below are equivalent, due to method and type embedding
	// fmt.Println(report1.days(1446)) // now, this produces an error (ambiguous selector) since there are two methods we can use and go compiler doesn't know which one to use 
	fmt.Println(report1.sol.days(1446)) // to solve it, chain methods through the inner type or implement a custom days method which receiver is of type report1
  fmt.Println(report1.days(1446))

}
