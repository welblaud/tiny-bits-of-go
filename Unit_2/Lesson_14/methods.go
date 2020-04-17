// methods.go
package main

import "fmt"

type celsius float64
type kelvin float64
type fahrenheit float64

// celsiusToKelvin °C to °K
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// celsiusToFahrenheit converts °C to °F
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// kelvinToCelsius converts °K to °C
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// kelvinToFahrenheit converts °K to °F
func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

// fahrenheitToCelsius converts °F to °C
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5 / 9)
}

// fahrenheitToKelvin converts °F to °K
func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func main() {
	var k kelvin = 294.0
	c := k.celsius()
	f := k.fahrenheit()
	fmt.Println(k, "°K is", c, "°C.")
	fmt.Println(k, "°K is", f, "°F.")
}
