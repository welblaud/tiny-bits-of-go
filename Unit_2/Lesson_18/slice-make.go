// slice-make.go

package main

import "fmt"

func main() {
	/* make prevents additional allocations and copies to enlarge the
	 * underlaying array; the capacity arg is optional, without it
	 * the append func would add the 11th element */
	dwarfs := make([]string, 0, 10)
	dwarfs = append(dwarfs, "Ceres", "Pluton", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs)

}
