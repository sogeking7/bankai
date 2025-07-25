package main

import "fmt"

func maps () {

	StudentAge := make(map[string] int)

	StudentAge["Sogeking"] = 20
	StudentAge["Luffy"] = 21
	StudentAge["Chopper"] = 18

	fmt.Printf("%v \n",StudentAge)
	fmt.Println(len(StudentAge))

	superHero := map[string]map[string]string {
		"Superman": map[string] string{
			"Realname": "Clark Kent",
			"City": "Metropolis",
		},
		"Batman": map[string] string {
			"Realname": "Bruce Wayne",
			"City": "Gotham City",
		},
	}

	if temp, exists := superHero["Superman"]; exists {
		fmt.Println(temp["Realname"], temp["City"])
	}
}
