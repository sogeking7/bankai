package main

import "fmt"

func operators() {

	x, y := 5, 6

	fmt.Println("x + y, ", x + y)
	fmt.Println("x - y, ", x - y)
	fmt.Println("x * y, ", x * y)
	fmt.Println("x / y, ", x / y)
	fmt.Println("x mod y, ", x % y)

	isBool := true // var isBool bool = true
	hate := false

	fmt.Println(isBool && hate)
	fmt.Println(isBool || hate)
	fmt.Println(!isBool)

	
}