package main

import "fmt"

func main() {
	// LIFO
	defer fmt.Println("Two")
	defer fmt.Println("One")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
