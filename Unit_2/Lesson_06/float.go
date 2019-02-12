// float.go
// Floating-point inaccuracies

package main

import (
  "fmt"
  "math"
)

func main() {

  third := 1.0 / 3.0
  fmt.Println(third + third + third)

  piggyBank := 0.1
  piggyBank += 0.2
  fmt.Println(piggyBank) // Big inaccuracy!

  fmt.Println( "22 nul:", math.Abs( piggyBank - 0.3 ) < 0.00000000000000000000001 )
  fmt.Println( "21 nul:", math.Abs( piggyBank - 0.3 ) < 0.0000000000000000000001 )
  fmt.Println( "20 nul:", math.Abs( piggyBank - 0.3 ) < 0.000000000000000000001 )
  fmt.Println( "19 nul:", math.Abs( piggyBank - 0.3 ) < 0.00000000000000000001 )
  fmt.Println( "18 nul:", math.Abs( piggyBank - 0.3 ) < 0.0000000000000000001 )
  fmt.Println( "17 nul:", math.Abs( piggyBank - 0.3 ) < 0.000000000000000001 )
  // The difference is comparable until the end of the inaccuracy is reached.
  fmt.Println( "16 nul:", math.Abs( piggyBank - 0.3 ) < 0.00000000000000001 )
  fmt.Println( "15 nul:", math.Abs( piggyBank - 0.3 ) < 0.0000000000000001 )
  fmt.Println( "14 nul:", math.Abs( piggyBank - 0.3 ) < 0.000000000000001 )
  fmt.Println( "13 nul:", math.Abs( piggyBank - 0.3 ) < 0.00000000000001 )
  fmt.Println( "12 nul:", math.Abs( piggyBank - 0.3 ) < 0.0000000000001 )
  fmt.Println( "11 nul:", math.Abs( piggyBank - 0.3 ) < 0.000000000001 )
  fmt.Println( "10 nul:", math.Abs( piggyBank - 0.3 ) < 0.00000000001 )
  fmt.Println( "9 nul:", math.Abs( piggyBank - 0.3 ) < 0.0000000001 )
  fmt.Println( "8 nul:", math.Abs( piggyBank - 0.3 ) < 0.000000001 )
  fmt.Println( "7 nul:", math.Abs( piggyBank - 0.3 ) < 0.00000001 )
  fmt.Println( "6 nul:", math.Abs( piggyBank - 0.3 ) < 0.0000001 )
  fmt.Println( "5 nul:", math.Abs( piggyBank - 0.3 ) < 0.000001 )
  fmt.Println( "4 nul:", math.Abs( piggyBank - 0.3 ) < 0.00001 )
  fmt.Println( "3 nuly:", math.Abs( piggyBank - 0.3 ) < 0.0001 )
  fmt.Println( "2 nuly:", math.Abs( piggyBank - 0.3 ) < 0.001 )
  fmt.Println( "1 nuly:", math.Abs( piggyBank - 0.3 ) < 0.01 )
  fmt.Println( "žádná nula:", math.Abs( piggyBank - 0.3 ) < 0.1 )

  fmt.Println( "absolutní hodnota:", math.Abs( piggyBank - 0.3 ) )

}
