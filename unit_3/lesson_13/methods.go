package main

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
