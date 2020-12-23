package main

import "fmt"

// Normal declaration
func add(x int, y int) int {
	return x + y
}

// Multiple params of same type
func _add(x, y int) int {
	return x + y
}

// Multiple return values, there is no limit to amount of return values
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// "Naked" return to make the function finish
	return
}

func main() {
	// Normal declaration
	fmt.Println(add(42, 13))

	// Multiple params of same type
	fmt.Println(_add(42, 13))

	// Multiple return values
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// Named return values
	fmt.Println(split(17))
}
