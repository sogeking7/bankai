package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Beginner Go!")

	// Variables

	var name string = "sogeking"
	fmt.Printf("This is my name %s\n", name)

	age := 20
	fmt.Printf("This is my age %d\n", age)

	var city string
	city = "Zhezkazgan"
	fmt.Printf("This is my city %s\n", city)

	country, continent := "Kazakhstan", "Asia"
	fmt.Printf("This is my continent %s and this is my country %s\n", continent, country)

	var (
		isEmployed bool   = true
		salary     int    = 5000
		position   string = "developer"
	)

	fmt.Printf("is employed: %t\n, this is my salary: %d\n, this is my position: %s\n", isEmployed, salary, position)

	// Zero Values

	var defaultInt int       // 0
	var defaultString string // ""
	var defaultFloat float64 // 0.0
	var defaultBool bool     // false

	fmt.Printf("int: %d, float: %f, string %s, bool: %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	const pi = 3.14

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	fmt.Printf("Monday: %d - Tuesday: %d - Wednesday: %d \n", Monday, Tuesday, Wednesday)

	const typedAge int = 25
	const untypedAge = 25

	fmt.Println(typedAge == untypedAge)

	// There is no Enums in Go
	const (
		Jan int = iota + 1 // 1
		Feb                // 2
		Mar                // 3
		Apr                // 4
	)

	fmt.Printf("jan - %d feb - %d mar - %d apr - %d", Jan, Feb, Mar, Apr)

	result := add(3, 4)
	fmt.Printf("This is the result: %d \n", result)

	sum, product := calculateSumAndProduct(10, 10)
	fmt.Printf("This is sum: %d, this is product: %d\n", sum, product)

	myAge := 20

	if myAge >= 18 {
		fmt.Println("you are an adult")
	} else if myAge >= 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a chld")
	}

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("its the weekend")
	}

	for i := 0; i < 5; i++ {
		fmt.Println("This is", i)
	}

	counter := 0
	for counter < 3 {
		fmt.Println("This is the counter", counter)
		counter++
	}

	// Infinite loop
	// for {
	// }

	// Arrays and Slices
	numbers := [5]int{10, 20, 30, 40, 50}
	// numbers[5] = 60 // adding more is impossible
	fmt.Printf("This is out array %v \n", numbers)
	fmt.Printf("This is last element in array %d \n", numbers[len(numbers)-1])

	// numbersAtInt := [...]int{}

	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Printf("This is matrix %v \n", matrix)

	// Slices - dynamic array, changes on the fly

	// allNumbers := numbers[:] // slice copy of array
	// firstThree := numbers[0:3]

	fruits := []string{"apple", "banana", "straberry"} // this is slice, no size
	fmt.Printf("these are my fruits %v \n", fruits)

	fruits = append(fruits, "kiwi")

	fmt.Printf("these are my fruits with kiwi %v \n", fruits)

	fruits = append(fruits, "mango", "pineapple")

	fmt.Printf("these are my fruits with more fruits %v \n", fruits)

	moreFruits := []string{"tomato", "blueberries"}

	fruits = append(fruits, moreFruits...)

	fmt.Printf("these are my fruits with more fruits %v \n", fruits)

	for index, value := range fruits {
		fmt.Printf("index %d and value %s \n", index, value)
	}

	capitalCities := map[string]string{
		"USA":        "Washington D.C.",
		"India":      "New Delhi",
		"Kazakhstan": "Astana",
	}
	fmt.Printf("This is map: %v \n", capitalCities)
	fmt.Println(capitalCities["USA"])
	fmt.Println(capitalCities["Germany"]) // "" zero variable

	capital, exists := capitalCities["Germany"]
	if exists {
		fmt.Println("This is the capital", capital)
	} else {
		fmt.Println("Does not exists")
	}

	delete(capitalCities, "Kazakhstan")
	fmt.Println("This is new deleted map: %v \n", capitalCities)

	// Structs
	// Datatype can hold data
	person := Person{
		Name: "Sogeking",
		Age:  20,
	}
	fmt.Printf("This is person %+v \n", person)

	// anonymous struct
	employee := struct {
		name string
		id   int
	}{
		name: "Alice",
		id:   12,
	}
	fmt.Printf("This is employee %v \n", employee)

	contact := Contact{
		Name: "Jhon",
		Address: Address{
			street: "V15",
			city:   "Astana",
		},
	}
	fmt.Printf("This is contact %v \n", contact)

	fmt.Println("Name before: ", person.Name)
	// modifyPersonName(&person)
	person.modifyPersonName("Tom")
	person.modifyPersonAge(24)

	fmt.Println("Name after: ", person.Name)

	x := 20
	ptr := &x

	fmt.Printf("The value of x: %d and address of x: %p \n", x, ptr)
	*ptr = 30
	fmt.Printf("The value of new x: %d and address of new x: %p \n", x, ptr)
}

// func modifyPersonName(person *Person) {
// 	person.Name = "Richard"
// 	fmt.Println("inside scope: ", person.Name)
// }

func (p *Person) modifyPersonName(name string) {
	p.Name = name
	fmt.Println("inside scope: ", p.Name)
}

func (p *Person) modifyPersonAge(age int) {
	p.Age = age
	fmt.Println("inside scope: ", p.Age)
}

// defined struct
type Person struct {
	Name string
	Age  int
}

type Address struct {
	street string
	city   string
}

type Contact struct {
	Name    string
	Address Address
	Phone   string
}

// is not exported, private
func add(a, b int) int {
	return a + b
}

func calculateSumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}
