// time.go
// 64-bit integers

package main

import (
  "fmt"
  "time"
)

func main() {

  future := time.Unix(12622780800, 0)
  fmt. Println(future) // Prints 2370-01-01 00:00:00 +0000 CET

}
