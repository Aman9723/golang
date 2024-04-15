package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	// no inheritance in Go, No super or parent
	user := User{"aman", "aman.dev", true, 22}
	fmt.Printf("User: %+v\n", user)
	fmt.Printf("Name is: %v\n", user.Name)
	user.GetStatus()
	user.NewMail()
	fmt.Println(user)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of user is: ", u.Email)
}
