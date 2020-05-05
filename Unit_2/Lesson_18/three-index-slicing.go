// three index slicing
package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	/* the last index means the capacity,
	 * it is safer! without it appending
	 * overwrites the underlaying array
	 * because it appends to the end of the slice,
	 * hence changes the next element in the array */
	terrestrial1 := planets[0:4:4]
	worlds := append(terrestrial1, "Ceres")
	fmt.Println(worlds)
	fmt.Println(planets)
	terrestrial2 := planets[0:4]
	worlds = append(terrestrial2, "Ceres")
	fmt.Println(worlds)
	fmt.Println(planets)
}
