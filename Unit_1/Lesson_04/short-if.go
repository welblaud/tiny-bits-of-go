// short-if.go
// Demonstrating the short declaration again.

package main
import (
  "fmt"
  "math/rand"
)

func main() {
  if num := rand.Intn(3); num == 0 {
    fmt.Println("Space Adventures")
  } else if num == 1 {
    fmt.Println("SpaceX")
  } else {
    fmt.Println("Virgin Galactic")
  } // num is no longer ni scope
}
