// struct.go
package main

import "fmt"

func main() {
	var curiosity struct {
		lat  float64
		long float64
	}
	// means 4°35'22.2" and 137°26'30.1" in degrees
	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}
