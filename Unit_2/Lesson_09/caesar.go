// caesar.go
package main

import (
	"fmt"
)

func main() {

	// Caesar's quote
	message := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(message); i++ {
		char := message[i]
		if char >= 'c' && char <= 'z' {
			char -= 3
			if char < 'a' { // restart again, abc has to become xyz
				char += 26
			}
		} else if char >= 'C' && char <= 'Z' {
			char -= 3
			if char < 'A' { // restart again, ABC has to become XYZ
				char += 26
			}
		}
		fmt.Printf("%c", char)
	}
	fmt.Println()
}
