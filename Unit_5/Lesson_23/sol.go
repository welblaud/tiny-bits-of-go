// sol.go
package main

import "fmt"

type sol int

type report struct {
	sol
	location
	temperature
}

// assigning to sol type
func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

// assigning and forwarding
func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

func main() {
	report := report{sol: 15}
	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))
}
