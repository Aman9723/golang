package main

import "fmt"

const LoginToken string = "abcdefgh" // Public

func main() {
	var username string = "aman"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallFLoat float64 = 255.987654321
	fmt.Println(smallFLoat)
	fmt.Printf("Variable is of type: %T \n", smallFLoat)

	// default values and some aliases
	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	// implicit type
	var website = "learn golang"
	fmt.Println(website)

	// no var style
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
}
