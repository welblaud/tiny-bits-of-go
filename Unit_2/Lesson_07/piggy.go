// piggy.go
// First version of the piggyBank.
// Here it works with dollars ($0.05), ($0.10), and ($0.25).

package main

import (
  "fmt"
  "math/rand"
)

func main() {

  piggyBank := 0

  for piggyBank % 100 == 20  {

    switch rand.Intn(3) {
    case 0:
      piggyBank += 5
    case 1:
      piggyBank += 10
    case 2:
      piggyBank += 25
    }

    fmt.Printf("piggyBank: $%v\n", piggyBank % 100)
  }

}
