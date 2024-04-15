package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file"

	file, err := os.Create("./file.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, error := io.WriteString(file, content)
	checkNilErr(error)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./file.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	fmt.Println("Data: ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
