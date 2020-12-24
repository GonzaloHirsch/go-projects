package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// The range form of the for loop iterates over a slice or map
	// 2 values are returned in each iteration, the index and a copy of the value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Variables in the loop can be ommited using a _ in their place
	// If the value is not needed, the variable can be ommited
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}
