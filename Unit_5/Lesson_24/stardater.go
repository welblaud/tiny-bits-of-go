// stardater.go
package main

import (
	"fmt"
	"time"
)

type stardater interface {
	YearDay() int
	Hour() int
}

// sol will satisfy the stardater interface
type sol int

func (s sol) YearDay() int {
	return int(s % 668)
}
func (s sol) Hour() int {
	return 0
}

// stardate returns a fictional measure of time
func stardate(t stardater) float64 {
	doy := float64(t.YearDay())
	h := float64(t.Hour()) / 24.0
	return 1000 + doy + h
}

func main() {
	// time.Date satisfies the stardater interface
	day := time.Date(2012, 8, 6, 5, 17, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	// sol satisfies the stardate interface, too
	s := sol(1422)
	fmt.Printf("%.1f Happy birthday\n", stardate(s))
}
