// canis.go
package main

import "fmt"

func main() {
	const lightSpeed = 299792
	const distance = 236000000000000000 // using 236e0 would convert the outcome to float!
	const days = 365
	const secondsPerDay = 24 * 3600
	const years = distance / lightSpeed / secondsPerDay / days

	fmt.Println("The journey from Earth to Canis Major Dwarf would take", years, "years.")
}
