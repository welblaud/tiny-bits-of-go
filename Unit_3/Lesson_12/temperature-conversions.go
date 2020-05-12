// kelvin.go
package main

import "fmt"

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

// celsiusToFahrenheit converts °C to °F
func celsiusToFahrenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32.0
	return f
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin)
	fahrenHeit := celsiusToFahrenheit(celsius)
	fmt.Println(kelvin, "°K is", celsius, "°C")
	fmt.Println(celsius, "°C is", fahrenHeit, "°F")
}
