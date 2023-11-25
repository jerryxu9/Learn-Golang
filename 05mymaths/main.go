package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// // random number:
	// // rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) // exclusive

	// random from crypto
	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}