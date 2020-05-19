// landing.go
package main

import "fmt"

func main() {
	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'})

	fmt.Printf("Spirit: %.3f\n", spirit)
	fmt.Printf("Opportunity: %.3f\n", opportunity)
	fmt.Printf("Curiosity: %.3f\n", curiosity)
	fmt.Printf("InSight: %.3f\n", insight)
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
