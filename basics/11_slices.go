package main

import (
	"fmt"
	"strings"
)

// Slice length and capacity
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// Creating slices with make
func printMakeSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func main() {
	// Slices are dynamically-sized structures
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// They are declared like a[low : high], they need a low bound and high bound
	// Low is included, but high is excluded
	var s1 []int = primes[1:4]
	fmt.Println(s1)

	// Slices are references to arrays, if the slice is modified, the array is too
	// Slices sharing the same instance of array will see the changes
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	// Slice literals
	// To build a slice literal, it's like an array but without the size
	// It generates the array and then a slice to reference it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s2)

	// Slice defaults
	// Low and high bounds can be ommited
	// The default for low is 0, and the default for high is the size of the array
	// For array var a [10]int, all these are equivalent
	// a[0:10]
	// a[:10]
	// a[0:]
	// a[:]
	s3 := []int{2, 3, 5, 7, 11, 13}

	s3 = s3[1:4]
	fmt.Println(s3)

	s3 = s3[:2]
	fmt.Println(s3)

	s3 = s3[1:]
	fmt.Println(s3)

	// Slice length and capacity
	// The length is the number of elements in the slice --> len(s)
	// The capacity is the number of element in the underlying array --> cap(s)
	// Slices can be extended, only if the array has enough capacity
	s4 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s4)

	// Slice the slice to give it zero length.
	s4 = s4[:0]
	printSlice(s4)

	// Extend its length.
	s4 = s4[:4]
	printSlice(s4)

	// Drop its first two values.
	s4 = s4[2:]
	printSlice(s4)

	// Nil slices
	// Nil slices have a length of 0 and a capacity of 0
	// They have no underlying array
	var s5 []int
	fmt.Println(s5, len(s5), cap(s5))
	if s5 == nil {
		fmt.Println("nil!")
	}

	// Creating slices with make
	// Using make you can create dynamically sized arrays
	// A third argument for capacity can be used
	// make(type, len, cap)
	_a := make([]int, 5)
	printMakeSlice("a", _a)

	_b := make([]int, 0, 5)
	printMakeSlice("b", _b)

	_c := _b[:2]
	printMakeSlice("c", _c)

	_d := _c[2:5]
	printMakeSlice("d", _d)

	// Slices of slices
	// Slices can be of any type, including other slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to a slice
	// One can append values to a slice using the "append" function
	// func append(s []T, vs ...T) []T
	// It takes the slice as parameter, and the value/s to be inserted
	// It returns the pointer to the slice
	// If the underlying array is too small, the function allocates a new one and returns the reference to the new array
	var s6 []int
	printSlice(s6)

	// append works on nil slices.
	s6 = append(s6, 0)
	printSlice(s6)

	// The slice grows as needed.
	s6 = append(s6, 1)
	printSlice(s6)

	// We can add more than one element at a time.
	s6 = append(s6, 2, 3, 4)
	printSlice(s6)
}
