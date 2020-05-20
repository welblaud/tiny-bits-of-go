// distance.go
package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
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

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

// newLocation from latitude, longitude d/m/s coordinates.
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {
	var mars = world{radius: 3389.5}
	var martians = map[string]location{}
	martians["spirit"] = newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	martians["opportunity"] = newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	martians["curiosity"] = newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.12, 'E'})
	martians["insight"] = newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'})

	martians["olympusMons"] = newLocation(coordinate{5, 4, 48, 'S'}, coordinate{137, 51, 0.0, 'E'})
	martians["mountSharp"] = newLocation(coordinate{18, 39, 0.0, 'N'}, coordinate{226, 12, 0.0, 'E'})

	fmt.Printf("%-22v %.2f km\n", "Spirit–Opportunity:", mars.distance(martians["spirit"], martians["opportunity"]))
	fmt.Printf("%-22v %.2f km\n", "Spirit–Curiosity:", mars.distance(martians["spirit"], martians["Curiosity"]))
	fmt.Printf("%-22v %.2f km\n", "Spirit–InSight:", mars.distance(martians["spirit"], martians["insight"]))
	fmt.Printf("%-22v %.2f km\n", "Opportunity–Curiosity:", mars.distance(martians["opportunity"], martians["curiosity"]))
	fmt.Printf("%-22v %.2f km\n", "Opportunity–InSight:", mars.distance(martians["opportunity"], martians["insight"]))
	fmt.Printf("%-22v %.2f km\n", "Curiosity–InSight:", mars.distance(martians["curiosity"], martians["insight"]))

	fmt.Printf("%-22v %.2f km\n", "Ol. Mons–Mt. Sharp:", mars.distance(martians["olympusMons"], martians["mountSharp"]))

	var earth = world{radius: 6371.0}
	var cities = map[string]location{}
	cities["london"] = newLocation(coordinate{51, 30, 0.0, 'N'}, coordinate{0, 0.8, 0.0, 'W'})
	cities["paris"] = newLocation(coordinate{48, 51, 0.0, 'N'}, coordinate{2, 21, 0.0, 'E'})
	cities["prague"] = newLocation(coordinate{50, 5.0, 19.0, 'N'}, coordinate{14, 25, 17.0, 'E'})
	cities["rez"] = newLocation(coordinate{50, 10.0, 23.7, 'N'}, coordinate{14, 21, 48.0, 'E'})

	fmt.Printf("%-22v %.2f km\n", "London–Paris:", earth.distance(cities["london"], cities["paris"]))
	fmt.Printf("%-22v %.2f km\n", "Řež–Praha:", earth.distance(cities["rez"], cities["prague"]))
}
