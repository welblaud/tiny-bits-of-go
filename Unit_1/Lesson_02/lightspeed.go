// How long does it take to get to Mars?
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	const hoursPerDay = 24
	var speed = 100800  // km/h
	distance = 96300000 // Earth from Mars
	fmt.Println(distance/speed/hoursPerDay, "days")

	distance = rand.Intn(345000000) + 56000000 // Generates random distance (rand diff + min dist)
	fmt.Println(distance/speed/hoursPerDay, "days")
	// Random numbers remain the same (beware the Seeding and caching (more advanced stuff))
	var numRand = rand.Intn(10)
	fmt.Println(numRand)
	numRand = rand.Intn(10)
	fmt.Println(numRand + 2)
}

//  Declarations of multiple variables:
//  var one = 1
//  var two = 2
//  or
//  var one, two = 1, 2
//  or
//  var (
//    one = 1
//    two = 2
//  )
