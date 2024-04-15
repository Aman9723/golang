package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in Golang")
	greeter()

	proAdderResult, _ := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Result is: ", proAdderResult)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum, "Hi Pro result function"

}

func greeter() {
	fmt.Println("Namastey from golang")
}
