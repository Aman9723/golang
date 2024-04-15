package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Web requests in golang")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)

	defer response.Body.Close() // Close the response body

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data: ", string(databytes))
}
