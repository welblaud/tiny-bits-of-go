// sort.go
package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	// converting each planet to StringSlice and sorting
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}
