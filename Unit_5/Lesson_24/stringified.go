// stringified.go
package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long coordinate
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// because we want to print for the whole location
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}
	// unnecessary to use String(), it is run by Println (Stringer!)
	fmt.Println("Elysium Planitia is at", elysium)

}
