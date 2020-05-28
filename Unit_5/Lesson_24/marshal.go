// marshal.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long coordinate
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v°%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

// because we want to print for the whole location
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

// decimal converts a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type locationJSON struct {
	Decimal    float64 `json:"decimal"`
	Dms        string  `json:"dms"`
	Degrees    float64 `json:"degrees"`
	Minutes    float64 `json:"minutes"`
	Seconds    float64 `json:"seconds"`
	Hemisphere string  `json:"hemisphere"`
}

func (l *location) MarshalJSON() {
	lat := locationJSON{
		l.lat.decimal(),
		l.lat.String(),
		l.lat.d,
		l.lat.m,
		l.lat.s,
		string(l.lat.h),
	}
	long := locationJSON{
		l.long.decimal(),
		l.long.String(),
		l.long.d,
		l.long.m,
		l.long.s,
		string(l.long.h),
	}
	json1, err := json.MarshalIndent(lat, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	json2, err := json.MarshalIndent(long, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json1))
	fmt.Println(string(json2))
}

func main() {
	elysium := location{
		lat:  coordinate{4, 30, 0.0, 'N'},
		long: coordinate{135, 54, 0.0, 'E'},
	}
	elysium.MarshalJSON()

}
