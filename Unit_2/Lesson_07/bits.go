// bits.go
// Display the bits.

package main

import "fmt"

func main() {

  var green uint8 = 3
  fmt.Printf("%08b\n", green) // Prints 00000011
  green++
  fmt.Printf("%08b\n", green) // Prints 00000100

}
