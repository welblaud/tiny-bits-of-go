// scope-rules.go
// Another demonstration of scoping.

package main
import (
  "fmt"
  "math/rand"
)

var era = "AD" // Package global var (can't be initialized with short notation.

func main() {
  year := 2018 // Era and year are in scope.

  switch month := rand.Intn(12) + 1; month { // Era, year, and month are in scope.
  case 2:
    day := rand.Intn(28) + 1 // Era, year, month, and day are in scope.
    fmt.Println(era, year, month, day)
  case 4, 6, 9, 11:
    day := rand.Intn(30) + 1 // It's a new day.
    fmt.Println(era, year, month, day)
  default:
    day := rand.Intn(31) + 1 // It's a new day.
    fmt.Println(era, year, month, day)
  } // Month and day are out of scope.

} // Year is no longer in scope.
