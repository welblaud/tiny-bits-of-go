// ticket-to-mars.go
// Capstone: Ticket to Mars

package main

import (
	"fmt"
	"math/rand"
)

const secondsPerDay = 86400

func main() {

	distance := 62100000
	company := ""
	tripTip := ""

	fmt.Println("Spaceline       Days Trip type  Price")
	fmt.Println("=====================================")

	for ticket := 0; ticket < 10; ticket++ {

		switch num := rand.Intn(3); num {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "Space X"
		case 2:
			company = "Virgin Galactic"
		}

		speed := rand.Intn(16) + 15                  // 16-30 km/s
		price := speed + 20.0                        // millions
		duration := distance / speed / secondsPerDay // days

		if rand.Intn(2) == 0 {
			tripTip = "One-way"
		} else {
			tripTip = "Round-trip"
			price *= 2
		}

		fmt.Printf("%-17v %v %-10v $ %v\n", company, duration, tripTip, price)

	}

}
