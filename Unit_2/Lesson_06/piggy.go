// piggy.go

package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var piggyBank float64
	var nickel = 0.05
	var dime = 0.10
	var quarter = 0.25

	for piggyBank < 20.0 {

		switch rand.Intn(3) {
		case 0:
			piggyBank += nickel
		case 1:
			piggyBank += dime
		case 2:
			piggyBank += quarter
		}

		fmt.Printf("$ %05.2f\n", piggyBank)
	}
}
