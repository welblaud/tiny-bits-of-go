// rounding-error.go
// Division first

package main

import "fmt"

func main() {

  celsius := 21.0 // division first makes rouding problems worse
  fmt.Print((celsius / 5.0 * 9.0) + 32, "° F\n")
  fmt.Print((9.0 / 5.0 * celsius) + 32, "° F\n")

}
