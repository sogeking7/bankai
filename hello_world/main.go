package main

// Numeric: int, float,
// Boolean: true, false
// String
// Derived: Pointer, Array, Structure, Map, Interface

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	// Variables
	var a int = 5
	var b float32 = 3.14
	const pi float64 = 3.14151394758

	var(
		c = 8
		d = 7
	)

	fmt.Println(a, b, c, d, pi)

	x, y := 14, 15

	fmt.Println(x)
	fmt.Println(y)

	// Operators
	operators()

	// Pointers
	pointers()

	// Format Printing
	formatp()

	// Loops 
	loops()
	
	// Conditions
	conditions()
	
	// Arrays
	arrays()

	// Map
	maps()

	// Functions
	functions()

	// Structures
	structures()

	// Interfaces
	interfaces()

	// I/O
	io()

	// WebServer
	webserver()
}