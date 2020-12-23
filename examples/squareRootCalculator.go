package main

import (
	"fmt"
	"math"
)

// Calculates sqrt in 10 steps
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

const DELTA = 0.000000000000001

// Calculates sqrt using a delta
func Sqrt_(x float64) float64 {
	var z, z_, e float64 = 1.0, 1.0, 1.0
	var s int = 0
	for DELTA < e {
		// Aproximate
		z -= (z*z - x) / (2 * z)
		// Calculate change
		e = math.Abs(z - z_)
		z_ = z
		// Count steps
		s++
		fmt.Println(z)
	}
	fmt.Println("Used %d steps", s)
	return z
}

func main() {
	fmt.Println("Value is %v", Sqrt(50))
	fmt.Println("Value is %v", Sqrt_(50))
}
