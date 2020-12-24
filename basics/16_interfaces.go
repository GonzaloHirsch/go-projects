package main

import (
	"fmt"
	"math"
)

// Interfaces are a set of method signatures
type Abser interface {
	Abs() float64
}

// MyFloat has a receiver for the Abs() function
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex has a receiver for the abs function
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Implicit implementation
// Interfaces are implemented implicitly, they decouple definition from implementation
type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

// Interface values
// Interfaces can be thought as a tuple of (value, type), an interface value holds a value for a specific underlying concrete type
// Calling a method on an interface value executes the method with the same name on the type
type I2 interface {
	M2()
}

type T2 struct {
	S string
}

func (t *T2) M2() {
	fmt.Println(t.S)
}

type F2 float64

func (f F2) M2() {
	fmt.Println(f)
}

func describe2(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Nil underlying values
// If the concrete value is nil, the method is called with a nil receiver
// In Go it's common to write methods that handle the nil receiver
// The value itself cannot be a nil pointer, it can be a nil structure or type
type I3 interface {
	M3()
}

type T3 struct {
	S string
}

func (t *T3) M3() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe3(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Empty interface
// The empty interface can hold any value that implements at least 0 methods
// It is used for values of unknown type
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// Stingers
// It is the most present interface, many packages look for types that implement it when printing
// It is the equivalent to the toString() in Java
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// a = v --> This line is wrong on purpose

	fmt.Println(a.Abs())

	// Implicit implementation
	var i I = T{"hello"}
	i.M()

	// Interface values
	var i2 I2

	i2 = &T2{"Hello"}
	describe2(i2)
	i2.M2()

	i2 = F2(math.Pi)
	describe2(i2)
	i2.M2()

	// Nil underlying values
	var i3 I3

	var t3 *T3
	i3 = t3
	describe(i3)
	i3.M3()

	i3 = &T3{"hello"}
	describe(i3)
	i3.M3()

	// Empty interface
	var i4 interface{}
	describe(i4)

	i4 = 42
	describe(i4)

	i4 = "hello"
	describe(i4)

	// Stringers
	a2 := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a2, z)
}
