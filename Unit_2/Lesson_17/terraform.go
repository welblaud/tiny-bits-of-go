// terraform.go --> my own + tiny improvement acc.
package main

import "fmt"

// Planets is the main type for every planet
type Planets []string

func (planets Planets) terraform() {
	for planetPos := range planets {
		planets[planetPos] = "New " + planets[planetPos]
	}
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	Planets(planets[3:4]).terraform()
	Planets(planets[6:]).terraform()
	fmt.Println(planets)
}
