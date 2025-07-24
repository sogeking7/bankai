package main

import "fmt"

func arrays () {

	var EvenNum[5] int 

	EvenNum[0] = 0
	EvenNum[1] = 2
	EvenNum[2] = 4
	EvenNum[3] = 6
	EvenNum[4] = 8

	fmt.Println(EvenNum[2])

	OddNum := [5] int {1, 3, 5, 7, 9}

	for i := 1; i < len(OddNum); i ++ {
		fmt.Printf("%d\n", OddNum[i])
	}
	fmt.Println()
	for _, value := range EvenNum {
		fmt.Println(value)
	}

	numSlice := [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliced := numSlice[:7]
	fmt.Println(sliced)
	
	sliced2 := make([] int, 5, 10)
	fmt.Printf("sliced2=%v, len=%d, cap=%d \n", sliced2, len(sliced2), cap(sliced2))

	copy(sliced2, sliced)

	fmt.Printf("sclied2=%v \n", sliced2)

	sliced3 := append(numSlice, 11, 12)

	fmt.Printf("sliced3=%v \n", sliced3)
}