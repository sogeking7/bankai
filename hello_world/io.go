package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func io() {
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hello World! \nNew line!")
	file.Close()

	stream, err := ioutil.ReadFile("sample.txt")


	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream)

	fmt.Println(s1)
}