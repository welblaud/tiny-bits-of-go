// fmt.go

package main

import "fmt"

func main() {
	// It is possible to pass several arguments to the Print and Println functions, they will concatenated.
	// It is also possible to use format verbs (e.g., %v) in Printf function.
	fmt.Print("My weight on the surface of Mars is ", 72*0.3783, "kilograms, and I would be ", 38*365.2425/687, " years old.\n")
	fmt.Println("My weight on the surface of Mars is", 72*0.3783, "kilograms, and I would be ", 38*365.2425/687, "years old.")
	fmt.Printf("My weight on the surface of Mars is %v kilograms,", 72*0.3783) // Like Print, but with format verbs.
	fmt.Printf(" and I would be %v years old.\n", 38*365/687)
	fmt.Printf("My weight on the surface of %v is %v kilograms.\n", "Earth", 72) // Multiple format verbs.
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)                                     // Aligning text. Negative value pads the value to left, positive to right.
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
