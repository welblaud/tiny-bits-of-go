// malacandra.go

package main

import "fmt"

func main() {
  const days = 28
  var distance = 56000000
  var speed = 56000000 / 28 / 24
  // or 56000000/(28*24)

  fmt.Printf("For the given distance of %v km, the ship has to travel %v km/h for reaching Mars in %v days.", distance, speed, days)
}
