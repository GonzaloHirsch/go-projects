package main

import "fmt"

// All variables of same type, bool inits to false
// No initializers in these variables
var c, python, java bool

// With initializers, 1 per variable
var _i, _j int = 1, 2

// Constants, declared with =, not with :=
const Pi = 3.14

func main() {
	// No initializers
	// Integer variable, inits to 0
	var i int
	fmt.Println(i, c, python, java)

	// With initializers, 1 per variable
	var _c, _python, _java = true, false, "no!"
	fmt.Println(_i, _j, _c, _python, _java)

	// Implicit "var" declaration, implicit type from initializer
	// Short initializer only inside functions
	var __i, __j int = 1, 2
	__k := 3
	__c, __python, __java := true, false, "no!"
	fmt.Println(__i, __j, __k, __c, __python, __java)

	// Constants
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
