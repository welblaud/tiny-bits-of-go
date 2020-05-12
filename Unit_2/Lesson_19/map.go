// map.go
package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v° C.\n", temp)
	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

	// missing value is set to zero
	moon := temperature["Moon"]
	fmt.Println("Moon temp is missing:", moon)

	/* comma, ok syntax ('ok' could have any name),
	 * if the key is missing, ok will be false here */
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %v° C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
	if venus, ok := temperature["Venus"]; ok {
		fmt.Printf("On average the Venus is %v° C.\n", venus)
	} else {
		fmt.Println("Where is the Venus?")
	}
	temperature["Moon"] = 0
	// and here the zero is allowed
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %v° C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
}
