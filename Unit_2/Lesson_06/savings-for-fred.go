// savings-for-fred
// Write a program that randomly places nickels ($0.05), dimes ($0.10),
// and quarters ($0.25) into an empty piggy bank until it contains
// at least $20.00. Display the running balance of the piggy bank after
// each deposit, formatting with an appropriate width and precision.

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	piggyBank := 0.0 // initial value

	for piggyBank < 20.00 {

		// Also possible just: switch rand.Intn(2) {
		switch num := rand.Intn(2); num {
		case 0:
			piggyBank += 0.05
		case 1:
			piggyBank += 0.10
		case 2:
			piggyBank += 0.25
		}

		fmt.Printf("In the piggy is: $%05.2f\n", piggyBank)

	}

}
