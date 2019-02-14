// integers.go
// Integers wrap around.

package main

import "fmt"

func main() {

  var red uint8 = 255
  red++
  fmt.Println(red) // Prints 0

  var number int8 = 127
  number++
  fmt.Println(number) // Prints -128

  // Add a number larger than 1
  red += 3
  fmt.Println(red) // Prints 3

  // Wrap the other way
  red = 0
  red--
  fmt.Println(red) // Prints 255

  number = -128
  number--
  fmt.Println(number) // Prints 127 

  // Wrapping with a 16-bit unsigned integer
  var bigger uint16 = 65535
  bigger++
  fmt.Println(bigger) // Prints 0

}
