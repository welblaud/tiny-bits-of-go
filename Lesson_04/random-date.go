// random-date.go
// Refactoring of scope-rules.go

package main
import (
  "fmt"
  "math/rand"
)

var era = "AD"

func main() {
  year := 2018
  month := rand.Intn(12) + 1
  daysInMonth := 31

  // Here, the switch decides and assigns
  // new value appropriately.
  switch month {
  case 2:
    daysInMonth = 28
  case 4, 6, 9, 11:
    daysInMonth = 30
  }

  day := rand.Intn(daysInMonth) + 1 // Smart!
  fmt.Println(era, year, month, day)

}
