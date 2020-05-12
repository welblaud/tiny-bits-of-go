// capacity.go
/* official version:
 * s := []string{}
 * lastCap := cap(s)
 * for i := 0; i < 10000; i++ {
 *     s = append(s, "An element")
 *     if cap(s) != lastCap {
 *         fmt.Println(cap(s))
 *         lastCap = cap(s)
 *     }
 * }
 */
package main

import "fmt"

func main() {
	animals := []string{"dog", "cat", "bat"}
	for i := 0; i < 10; i++ {
		animals = append(animals, "unknown species!")
		fmt.Printf("length: %v, capacity: %v\n", len(animals), cap(animals))
	}
}
