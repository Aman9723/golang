package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	// no inheritance in Go, No super or parent
	user := User{"aman", "aman.dev", true, 22}
	fmt.Printf("User: %+v\n", user)
	fmt.Printf("Name is: %v\n", user.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
