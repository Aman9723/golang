package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "RB")

	// loops in maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
