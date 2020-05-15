// location.go
package main

import "fmt"

func main() {
	type location struct {
		lat  float64
		long float64
	}
	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var oportunity location
	oportunity.lat = -1.9462
	oportunity.long = 354.4734

	fmt.Println(spirit, oportunity)
}
