// compose.go
package main

import "fmt"

/* unorganized
type report struct {
	sol       int
	high, low float64
	lat, long float64
}
*/

type report struct {
	sol         int
	temperature temperature
	location    location
}
type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}
type celsius float64

// binding method to smaller part
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// binding method of the smaller part to the bigger part
func (r report) average() celsius {
	return r.temperature.average()
}

// days count the distance between two places (see lesson 22)
func (l location) days(l2 location) int {
	// To-do
	return 5
}

func main() {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v° C\n", report.temperature.high)
	// using the temperature's method
	fmt.Printf("average %v° C\n", report.temperature.average())
	// using the report's method of temprature
	fmt.Printf("average %v° C\n", report.average())
}
