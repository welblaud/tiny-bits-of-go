// struct-literal.go
package main

import "fmt"

func main() {

	type location struct {
		lat, long float64
	}

	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	/*
		Also possible as:
		insight := location{4.5, 135.9}
		... but it is necessary to perserve the order and
		number of fields
	*/
	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)

	spirit := location{-14.5684, 175.472636}
	// printing the struct with field names
	fmt.Printf("%+v\n", spirit)
}
