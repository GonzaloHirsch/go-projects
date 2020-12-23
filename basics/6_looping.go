package main

import "fmt"

// The for loop has init, condition and post parts
// Init: i := 0
// Condition: i < 10
// Post: i++
func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

// The Init and Post condition can be ommited, transforming the "for" into a "while"
func whileLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// Ommiting the Condition creates a loop that runs forever
// Calling it hangs the programs
func foreverLoop() {
	for {

	}
}

func main() {
	forLoop()
	whileLoop()
}
