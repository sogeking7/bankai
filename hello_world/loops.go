package main

import "fmt"

func loops () {
	// For loop
	for i := 1; i <= 10; i ++ {
		fmt.Println(i)
	}

	// While loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i ++
	}

	// Nested loop
	for i := 1; i <= 10; i ++ {
		for j := 1; j <= 10; j ++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}