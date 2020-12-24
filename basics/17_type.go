package main

import "fmt"

// Type switches
// These switches can assert multiple types at once
// The syntax is like a regular switch, but the cases are types
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	// Type assertions can be used to test if an interface holds a value of specific type
	// The 2 possible assertions are:
	// t := i.(T) --> Tests if the interface has a value T, if not, it panics
	// t, ok := i.(T) --> Tests if the interface has a value T and returns the type if present, with ok true, if not, zero value for the type and ok false
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)

	// Type switches
	do(21)
	do("hello")
	do(true)
}
