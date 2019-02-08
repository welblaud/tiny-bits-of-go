// countdown-with-failure.go
// Demonstrating the for loop with a timer!

package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {

  var count = 10
  for count > 0 {
    if rand.Intn(5) == 3 {
      fmt.Println("Stop! Check the left engine!")
      break
    } else {
      fmt.Println(count)
    }
    time.Sleep(time.Second)
    count--
  }
  if count == 0 {
    fmt.Println("Liftoff!")
  } else {
    fmt.Println("Launch failed.")
  }

}
