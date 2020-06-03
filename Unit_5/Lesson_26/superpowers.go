// superpowers.go
package main

import "fmt"

func main() {

	/* as with structures, composite literals for array can also by prefixed
	 * with the address operator (&) to create a new pointer; arrays also
	 * provide automatic dereferencing */
	superpowers := &[3]string{"flight", "invisibility", "super strength"}
	fmt.Println(superpowers[0])   // array
	fmt.Println(superpowers[1:2]) // slice
}
