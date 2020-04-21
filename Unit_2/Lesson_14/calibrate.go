// calibrate.go
package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

// ready for plugging in a real sensor
func realSensor() kelvin {
	return 0
}

// fakeSensor generates random temperature on Mars
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

// calibrate improves a sensor functionality
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 4
	sensor1 := calibrate(realSensor, offset)
	fmt.Println(sensor1())
	sensor2 := calibrate(fakeSensor, offset)

	for i := 0; i < 10; i++ {
		fmt.Println(sensor2())
	}
}
