// sturct.go
package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func main() {

	// pointer to a struct is very common
	timmy := &person{
		name: "Timothy",
		age:  10,
	}

	// it isn't necessary to dereference structures -> (*timmy).superpower
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)
}
