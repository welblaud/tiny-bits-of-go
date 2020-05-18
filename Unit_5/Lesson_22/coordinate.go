// coordinate.go
package main

import "fmt"

func main() {
	// Bradbury Landing 4°35'22.2" S, 137°26'30.1" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}
	fmt.Println(lat.decimal(), long.decimal())

	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	fmt.Println(curiosity)
}

// coordinate in degrees, minutes, seconds in a N/S/E/W hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// location consisting of latitude and longitude
type location struct {
	lat, long float64
}

// decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}
