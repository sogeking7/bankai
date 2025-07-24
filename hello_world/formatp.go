package main

import "fmt"

func formatp() {
	var name string = "Bankai"
	var pi float64 = 3.14123453
	var win = true
	var x int = 1000

	fmt.Println(len(name))
	fmt.Println(name + " - concat strings")

	fmt.Printf("pi=%f \n", pi) // float
	fmt.Printf("pi=%.3f \n", pi) // with floating point
	fmt.Printf("%T \n", name) // Type
	fmt.Printf("%s \n", name) // String
	fmt.Printf("%t \n", win) // Boolean
	fmt.Printf("%d \n", x) // Int
	fmt.Printf("%b \n", 4134) // Binary
	fmt.Printf("%c \n", 65) // ASCII Code
	fmt.Printf("%x \n", 50) // Hex Code
	fmt.Printf("%e \n", pi) // Scientific notation
}