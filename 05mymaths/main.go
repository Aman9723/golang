package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// var myNumberOne int = 2
	// var myNumberTwo float64 = 4.5

	// fmt.Println("Addition: ", myNumberOne+int(myNumberTwo))

	// random from math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	// random from crypto/rand
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
