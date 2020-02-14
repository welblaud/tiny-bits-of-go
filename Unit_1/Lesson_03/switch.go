// concise-switch.go
// demonstrating the switch loop

package main

import "fmt"

func main() {

	fmt.Println("There is a cavern entrance here and a path to the east.")
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough // falls through to the next case, default in JS and Java
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}

}
