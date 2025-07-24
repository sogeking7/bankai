package main

import "fmt"

func pointers () {
	x := 10

	fmt.Println(x)
	fmt.Println(&x) // address where x stores, 

	changeValue(&x)

	fmt.Println(x)
}

// Basic example 
func changeValue(x *int) { 
	*x = 7
}