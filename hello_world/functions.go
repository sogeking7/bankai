package main

import "fmt"

func functions() {
	x, y := 5, 6
	fmt.Println(add(x, y))
	f := 5
	fmt.Printf("Factorial of %d = %d \n", f, fact(f))
	
	defer FirstRun()
	SecondRun()

	fmt.Println(div(3, 0))
	fmt.Println(div(10, 5))
	demPanic()

	fmt.Println(addemup(10, 20, 30, 40, 50))

}

func addemup(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}
	return sum
}

func add (x, y int) int {
	return x + y
}

func fact (x int) int {
	if x == 0 {
		return 1
	}
	return x * fact(x - 1)
}

func FirstRun () {
	fmt.Println("I executed First")
}

func SecondRun () {
	fmt.Println("I executed Second")
}

func div (num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

func demPanic () {
	defer func() {
		fmt.Println(recover())
	}()
	panic("Panic")
}


// Defer 
// Defer statemenet defers the execution of a function until the surrounding function returns.
// Multiple defers are pushed into stack and executes in LIFO
// Defer generally used to cleanup resources like file, db connection etc.

// Recover
// Recover is another built-in function in go.
// It helps to regain the normal flow of execution after a panic.
// Generally it used with defer statement to recover panic in goroutine.

// Panic
// Panic is similar to throwing exception like other language
// Generally when a panic is called the normal execution flow is stops immediately
// The deffered functions are executed normally.