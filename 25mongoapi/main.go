package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aman9723/mongoapi/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server is starting on port 4000...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is listening on port 4000...")
}
