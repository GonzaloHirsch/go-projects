package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Seeding by hand, can be done with time package
	rand.Seed(10000)
	fmt.Println("My favorite number is", rand.Intn(10))
}
