package main

import (
	"fmt"
	"math"
)

// Function values
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// Function closures
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// Functions are values too, they can be passed around and used as variables
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Function closures
	// A closure is a function value that references variables from outside it's body
	// The function can access and assing those variables, it's "bound" to them
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
