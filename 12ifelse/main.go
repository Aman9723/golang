package main

import "fmt"

func main() {
	fmt.Println("If Else in Go")

	loginCount := 23

	if loginCount < 10 {
		fmt.Println("Regular User")
	} else if loginCount > 10 {
		fmt.Println("Watch out")
	} else {
		fmt.Println("Exactly 10 logins")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

	// if err != nil {

	// }
}
