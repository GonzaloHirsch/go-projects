package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// If
// Ifs do not have parenthesis
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// If with statement
// If statements can have a short statement to execute before the condition
// Here it is declared v before executing the if condition
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// If and else
// Variables declared in the if short statement can be used in the else block
func _pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// Switch
// The switch does not need a "break" statement, it is already included
func osDetector() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// Switch with no condition
// A switch with no condition is equivalent to switch true, so it's like a if-else with many else if
func timeGreeting() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Defer stacking
// Defer calls can be stacked, and will be performed in a LIFO order
func stackDefers() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	// Defer
	// The execution of a defer statement is done after the surrounding function returns
	defer fmt.Println("I'm done")

	// If
	fmt.Println(sqrt(2), sqrt(-4))

	// If with statement
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// If and else
	fmt.Println(
		_pow(3, 2, 10),
		_pow(3, 3, 20),
	)

	// Switch
	osDetector()

	// Switch with no condition
	timeGreeting()

	// Defer stacking
	stackDefers()
}
