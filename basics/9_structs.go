package main

import "fmt"

// Structs are collections of fields
type Vertex struct {
	X int
	Y int
}

// Struct literals
// Fields can be initialized using named parameters
var (
	v1 = Vertex{1, 2}           // has type Vertex
	v2 = Vertex{X: 1}           // Y:0 is implicit
	v3 = Vertex{}               // X:0 and Y:0
	v4 = Vertex{Y: 150, X: 250} // X:250 and Y:150
	_p = &Vertex{1, 2}          // has type *Vertex
)

func main() {
	// Structs
	fmt.Println(Vertex{1, 2})

	// Struct fields accessed using dot operator
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// Pointers to structs
	// Fields can be accessed through a pointer to the struct
	// There is no need to dereference the pointer, it makes the code easier to read
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// Struct literals
	fmt.Println(v1, _p, v2, v3, v4)
}
