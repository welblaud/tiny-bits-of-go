// gps.go
package main

import (
	"fmt"
	"math"
)

// gps tool
type gps struct {
	world
	from location
	to   location
}

// any wheel-based vehicle
type rover struct {
	gps
}

// any planet
type world struct {
	radius float64
}

// current location consisting of latitude and longitude
type location struct {
	name      string
	lat, long float64
}

// return formatted information about a location
func (l location) description() string {
	return fmt.Sprintf("%v, (%.2f, %.2f)", l.name, l.lat, l.long)
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

// distance measures the distance between two locations
func (g gps) distance() float64 {
	return g.world.distance(g.from, g.to)
}

// message prints the info about remaining distance
func (g gps) message() string {
	return fmt.Sprintf("Km remaining: %.2f\n", g.distance())
}

func main() {

	var mars = world{radius: 3389.5}
	bradbury := location{"Bradbury Landing", -4.5895, 137.4417}
	elysium := location{"Elysium Planitia", 4.5, 135.9}
	gps := gps{
		world: mars,
		from:  bradbury,
		to:    elysium,
	}
	curiosity := rover{
		gps: gps,
	}
	fmt.Print(curiosity.message())
}
