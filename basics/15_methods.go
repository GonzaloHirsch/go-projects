package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Go does not have classes, but methods can be defined for types
// Functions can take a receiver argument, which is the type assigned to the method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods are functions
// Methods are functions, they only take a special extra argument
// Here it is written as a function, but the implementation is the same
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Methods on non-structs
// Methods can be declared for non-struct types declared in the same package as the method
// Mathods cannot be declared for built-in types
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Pointer receivers
// Receivers can be pointers or values, the pointer is to a type, not a built-in type
// If the receiver is a pointer, the method can alter the original object, if not, only a copy is edited
// Same principle applies to functions receiving objects
// Why choose pointer receivers over value receivers?
// Pointer receivers can change the original object, and are more efficient over large structures to avoid copying the object to memory
// Types should have all value or pointer receivers
func (v *Vertex) ScaleOriginal(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleCopy(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// Methods
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	// Methods are functions
	fmt.Println(Abs(v))

	// Methods on non-structs
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Pointer receivers
	v2 := Vertex{3, 4}
	v2.Scale(10)
	fmt.Println(v2.Abs())
}
