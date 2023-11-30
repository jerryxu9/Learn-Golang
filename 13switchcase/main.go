package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	fmt.Println("Switch and case in Golang")

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Dice value is 2")
	case 3:
		fmt.Println("Dice value is 3")
	case 4:
		fmt.Println("Dice value is 4")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5")
	case 6:
		fmt.Println("Dice value is 6")
	default: 
		fmt.Println("Default message")
	}
}