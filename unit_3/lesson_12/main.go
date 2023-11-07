package main

import "fmt"

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	f := ((c * 9.0) / 5.0) + 32.0
	return f
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fmt.Println(kelvin, "º K is ", celsius, "º C")
	kelvin = 233.0
	fmt.Println(kelvin, "º K is ", kelvinToCelsius(kelvin), "º C")
	kelvin = 0
	fmt.Println(kelvin, "º K is ", kelvinToFahrenheit(kelvin), "º F")

}
