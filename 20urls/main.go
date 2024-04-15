package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://google.com:3000/search?q=golang"

func main() {
	fmt.Println("Handling urls in golang")

	// parsing
	result, err := url.Parse(myUrl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("RawQuery: ", result.RawQuery)
	fmt.Println("Port: ", result.Port())
	fmt.Println("Fragment: ", result.Fragment)

	qParams := result.Query()
	fmt.Println(qParams["q"])

	// building
	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "google.com",
		Path:     "/search",
		RawQuery: "q=golang",
	}
	fmt.Println(newUrl.String())
}
