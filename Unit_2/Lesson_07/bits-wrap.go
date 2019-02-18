// bits-wrap.go
// The bits when integers wrap.

package main

import "fmt"

func main() {

  var blue uint8 = 255
  fmt.Printf("%08b\n", blue)
  blue++ // here bit wrap, 255 becomes 0
  fmt.Printf("%08b\n", blue)

}
