package main

import "fmt"

// type kelvin float64
// type celsius float64
// type fahrenheit float64

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// func (k kelvin) toCelsius() celsius {
// 	return celsius(k - 273.15)
// }

// func (f fahrenheit) toCelsius() celsius {
// 	return celsius((f - 32.0) * 5.0 / 9.0)
// }

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {

	const degrees = 20

	var temperature celsius = 20
	temperature += degrees
	fmt.Println(temperature)
	var warmUp float64 = 10

	temperature += celsius(warmUp) // we cannot use warmup directly bc it is not of type celsius, we have to convert it first  (mismatched types celsius and float64)

	// types can't be mixed
	{
		type celsius float64
		type fahrenheit float64
		var c celsius = 20
		var f fahrenheit = 20
		fmt.Println(c, f)
		//! (mismatched types celsius and fahrenheit)
		// if c == f {
		// }
		// c += f
	}

	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	c = k.toCelsius()
	fmt.Println(k, "º K is ", c, "º C")
	fmt.Print(127, "º C is ", celsiusToKelvin(127), "º K")

}
