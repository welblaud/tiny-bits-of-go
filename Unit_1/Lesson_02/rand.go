// Random numbers
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Random numbers remain the same (beware the Seeding and caching (more advanced stuff))
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
