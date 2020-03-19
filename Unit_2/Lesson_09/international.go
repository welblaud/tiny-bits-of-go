// international.go
package main

import (
	"fmt"
)

func main() {

	message := "Hola Estación Espacial Internacional"

	for _, char := range message {
		if char >= 'c' && char <= 'z' {
			char += 13
			if char > 'z' { // restart again, n-z becomes a-m
				char -= 26
			}
		} else if char >= 'C' && char <= 'Z' {
			char += 13
			if char > 'Z' { // restart again, N-Z becomes A-M
				char -= 26
			}
		}
		fmt.Printf("%c", char)
	}
	fmt.Println()
}
