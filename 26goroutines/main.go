package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func main() {
	// go greeter("Hello")
	// greeter("World")
	websiteList := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.github.com",
	}

	for _, website := range websiteList {
		go getStatusCode(website)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(0)
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Println(result.StatusCode, endpoint)
	}
}
