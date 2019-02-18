// hexadecimals.go

package main

import "fmt"

func main() {

// var red, green, blue uint8 = 0, 151, 213
  var red, green, blue uint8 = 0x00, 0x8d, 0xd5

  // %x for printing hexadecimal values
  fmt.Printf("%x %x %x\n", red, green, blue)
  // output for CSS
  fmt.Printf("color: #%02x%02x%02x;\n", red, green, blue)

}
