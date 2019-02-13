// inspect.go
// Inspect a variable's type.

package main

import "fmt"

func main() {

  year := 2018
  days := 365.2425
  text := "My text test."
  boolean := true
  fmt.Printf("Type %T for %v\n", year, year)
  fmt.Printf("Type %T for %[1]v\n", days) // using argument, the first variable
  fmt.Printf("Type %T for %[1]v\n", text)
  fmt.Printf("Type %T for %[1]v\n", boolean)

}
