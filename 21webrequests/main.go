package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Performing GET request...")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "https://www.google.com"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("Response status:", response.Status)
	fmt.Println("Content length:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content:", string(content))
}

func PerformPostJsonRequest() {
	const myUrl = "https://www.google.com"

	// fake json payload
	requestBody := strings.NewReader(`
	{
		"coursename": "Let's go with golang",
		"price":0,
		"platform": "learCodeOnline.in"
	}`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content:", string(content))
}

func PerformPostFormRequest() {
	const myUrl = "https://www.google.com"

	// fake form payload
	data := url.Values{}
	data.Add("coursename", "Let's go with golang")
	data.Add("price", "0")
	data.Add("platform", "learCodeOnline.in")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content:", string(content))
}
