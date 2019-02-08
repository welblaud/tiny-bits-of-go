// concise-switch.go
// demonstrating the switch loop

package main
import "fmt"

func main() {

  fmt.Println("There is a cavern entrance here and a path to the east.")
  var command = "go inside"

  switch command {
  case "go east":
    fmt.Println("You head further up the mountain.")
  case "enter cave", "go inside": // It is possible to test more conditions!
    fmt.Println("You find yourself in a dimly lit cavern.")
  case "read sign":
    fmt.Println("The sign reads 'No Minors'.")
  default: // Default value.
    fmt.Println("Didn't quite get that.")
  }

}
