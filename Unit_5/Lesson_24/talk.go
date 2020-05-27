// talk.go
package main

import (
	"fmt"
	"strings"
)

// intergace
var t interface {
	talk() string
}

type martian struct{}

// talk function - martian satisfies the interface t
func (m martian) talk() string {
	return "nack nack"
}

type laser int

// talk function - laser satisfies the interface t
func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

// interface names usually end with -er
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	// polymorphism
	t = martian{} // based on the t interface
	fmt.Println(t.talk())
	t = laser(3) // based on the t interface
	fmt.Println(t.talk())
	shout(martian{})
	shout(laser(2))
}
