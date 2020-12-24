package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// Maps
var m map[string]Vertex

// Map literals
// Map literals need the keys to be declared
var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// More map literals
// If the value is a type name, it can be ommited from the literal
var m3 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	// The zero value of a map is nil, nil maps don't have keys and keys cannot be added
	// The notation for a map is map[KEY_TYPE]VALUE_TYPE
	// Creating a map with make returns an array ready for use, initialized
	// Accessing a non-existing key will generate a 0-valued object
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Bell Labs2"])

	// Map literals
	fmt.Println(m2)

	// More map literals
	fmt.Println(m3)

	// Mutating maps
	// Inserting/updating value --> m[key] = elem
	// Accessing value --> elem = m[key]
	// Deleting value --> delete(m, key)
	// Testing if map contains key --> elem, ok = m[key]
	// If key in the map, ok is true, false otherwise
	// If key in the map, elem is the value stored, if not, it is a zero value for the map type
	m4 := make(map[string]int)

	m4["Answer"] = 42
	fmt.Println("The value:", m4["Answer"])

	m4["Answer"] = 48
	fmt.Println("The value:", m4["Answer"])

	delete(m4, "Answer")
	fmt.Println("The value:", m4["Answer"])

	v, ok := m4["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
