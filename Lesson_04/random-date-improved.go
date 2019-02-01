// random-date-improved.go
// Upgrading the random-date.go
// This version generates 10 random days for years between 2000 and 2045.

package main
import (
  "fmt"
  "math/rand"
)

var era = "AD"

func main() {

  for dates := 0; dates < 10; dates++ {

    year := rand.Intn(45) + 2001
    month := rand.Intn(12) + 1
    daysInMonth := 31

    // Here, the switch decides and assigns
    // new value appropriately.
    switch month {
    case 2:
      if year % 400 == 0 || (year % 4 == 0 && year % 100 != 0) {
        daysInMonth = 28
      } else {
        daysInMonth = 29
      }
    case 4, 6, 9, 11:
      daysInMonth = 30
    }

    day := rand.Intn(daysInMonth) + 1 // Smart!
    fmt.Println(era, year, month, day)
  }

}
