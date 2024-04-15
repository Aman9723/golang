package main

import "fmt"

func main() {
	fmt.Println("Pointers in Go")

	// var ptr *int
	// fmt.Println("Value of ptr is: ", ptr)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Address of myNumber is: ", &ptr)
	fmt.Println("Value of myNumber is: ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of myNumber is: ", myNumber)
}
