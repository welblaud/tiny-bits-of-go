// whoops.go
package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}
	// maps aren't copied!
	planetsMarkII := planets
	planets["Earth"] = "whoops"
	fmt.Println(planetsMarkII)
	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
}
