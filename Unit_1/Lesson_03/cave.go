// cave.go
package main

import "fmt"

func main() {

	var point = "entrance"

	if point == "cave" {
		fmt.Println("You are inside a dark cave.")
	} else if point == "entrance" {
		fmt.Println("You are before a dark entrance. It probably leads to a much darker cave.")
	} else if point == "mountain" {
		fmt.Println("You are standing on the top of the mountain of the Dead.")
	} else {
		fmt.Println("What the hell are you standing in front of?!")
	}
}
