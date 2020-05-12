// sensor.go
package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// fakeSensor generates temperature between 150-300°K
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

// realSensor -> ready for plug
func realSensor() kelvin {
	return 0
}

func main() {
	// assigning of a function to variable
	sensor := fakeSensor
	// calling
	fmt.Println(sensor())
	sensor = realSensor
	fmt.Println(sensor())
}
