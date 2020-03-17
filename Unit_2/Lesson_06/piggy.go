// piggy.go
// First version of the piggyBank.
// Here it works with dollars.

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	piggyBank := 0

	nickel := 5
	dime := 10
	quarter := 25

	for piggyBank < 2000 {

		switch rand.Intn(3) {
		case 0:
			piggyBank += nickel
		case 1:
			piggyBank += dime
		case 2:
			piggyBank += quarter
		}

		dollars := piggyBank / 100
		cents := piggyBank % 100

		fmt.Printf("piggyBank: $%02d.%02d\n", dollars, cents)
	}

}
