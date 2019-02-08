// torch.go
// Demonstration of NOT operator.
package main
import "fmt"

func main() {

  var haveTorch = true;
  var litTorch = false

  if !haveTorch || !litTorch {
    fmt.Println("Nothing to see here.")
  }

}
