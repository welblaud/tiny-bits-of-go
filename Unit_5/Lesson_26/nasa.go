// nasa.go
package main

import "fmt"

// type pointer
var administrator *string

func main() {

	scolese := "Christopher J. Scolese"
	// administrator becomes a pointer
	administrator = &scolese
	fmt.Println(administrator)  // prints the address
	fmt.Println(*administrator) // prints the value

	bolden := "Charles F. Bolden"
	// administartor points to another address
	administrator = &bolden
	fmt.Println(*administrator) // prints the new value

	bolden = "Charles Frank Bolden Jr."
	fmt.Println(*administrator) // prints the changed value

	major := administrator                            // major is becoming a pointer
	*major = "Major General Charles Frank Bolden Jr." // changes the value of 'bolden'
	fmt.Println(bolden)                               // prints the new value of major [-> bolden]

	fmt.Println(administrator == major) // true, the same address

	lightfoot := "Robert M. Lightfoot Jr."
	administrator = &lightfoot          // new address for administrator, major remains the same
	fmt.Println(administrator == major) // false, the address differs

	charles := *major         // new var equals the value of bolden
	*major = "Charles Bolden" // changes the value of bolden
	fmt.Println(charles)      // prints the initial bolden
	fmt.Println(bolden)       // print the new value of bolden

	charles = "Charles Bolden"
	fmt.Println(charles == bolden)   // true, both values are the same
	fmt.Println(&charles == &bolden) // false, the address differs

}
