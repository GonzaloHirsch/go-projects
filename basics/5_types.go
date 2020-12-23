/*
BASIC TYPES IN GO:

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// Var declaration block
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Initial value for variables
	var i int     // --> 0
	var f float64 // --> 0
	var b bool    // --> false
	var s string  // --> ""
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// Type conversions
	// Expression is TYPE(v), converts the value v to type TYPE
	var x, y int = 3, 4
	var _f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(_f)
	fmt.Println(x, y, z)
}
