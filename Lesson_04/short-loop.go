// short-loop.go

package main

import "fmt"

func main() {

  for count := 10; count > 0; count-- {
    fmt.Println(count) // New local scope.
  } // Count is no longer in scope.

}
