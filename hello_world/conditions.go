package main

import "fmt"

func conditions () {
	age := 20

	if age >= 20  {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	switch age {
	case 16: 
		fmt.Println("Prepare for college")
	case 18: 
		fmt.Println("Don't run after girls")
	case 20: 
		fmt.Println("Get yourself a job!")
	default:
		fmt.Println("Are you even alive?")
	}
}