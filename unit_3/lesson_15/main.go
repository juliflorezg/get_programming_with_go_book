// CAPSTONE: TEMPERATURE TABLES
// for full exercise details refer to pg 117 in the book

package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

func (c celsius) toKelvin() kelvin {
	return kelvin(c + 273.15)
}
func (c celsius) toFahrenheit() fahrenheit {
	f := ((c * 9.0) / 5.0) + 32.0
	return fahrenheit(f)
}

func (k kelvin) toCelsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) toFahrenheit() fahrenheit {
	return fahrenheit((k.toCelsius()*9)/5 + 32)
}

func (f fahrenheit) toCelsius() celsius {
	return celsius((f - 32) / (9 / 5))
}
func (f fahrenheit) toKelvin() kelvin {
	return kelvin((f + 459.67) * (9 / 5))
}

type getRowFn func(v int) (string, string)

// todo write function to get a header for table  ✅
// =======================
// | ºC       | ºF       |
// =======================

func drawHeader(f, s string) {
	fmt.Println("=======================")
	fmt.Printf("| %-9v| %-9v|\n", f, s)
	fmt.Println("=======================")
}

var fromCelsiusToF = func(c celsius) fahrenheit {
	return c.toFahrenheit()
}

func ctof(v int) (string, string) {
	v1 := celsius(v*5 - 40)
	v2 := celsius(v1).toFahrenheit()

	cell1 := fmt.Sprintf("%.1f", v1)
	cell2 := fmt.Sprintf("%.1f", v2)

	return cell1, cell2
}

func ftoc(v int) (string, string) {
	v1 := fahrenheit(v*5 - 40)
	v2 := fahrenheit(v1).toCelsius()

	cell1 := fmt.Sprintf("%.1f", v1)
	cell2 := fmt.Sprintf("%.1f", v2)

	return cell1, cell2
}

// todo write a function to get a row of dat to get F degrees for a given C temperature

//this function draws a table for the given parameters and gets temperature conversion
func drawTable(fistTitle, secondTitle string, getData getRowFn, numOfRows int) {
	drawHeader(fistTitle, secondTitle)
	for i := 0; i < numOfRows; i++ {
		cell1, cell2 := getData(i)
		fmt.Printf("| %-9v| %-9v|\n", cell1, cell2)
	}
	fmt.Println("=======================")
}

func main() {
	// drawHeader
	drawTable("°C", "°F", ctof, 29)
	fmt.Println()
	drawTable("°F", "°C", ftoc, 29)

}
