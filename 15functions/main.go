package main

import "fmt"

func main()  {
	fmt.Println("Welcome to functions in Golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 5, 8, 7)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro message is: ", myMessage)
}


func greeter()  {
	fmt.Println("Hello from Golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// If the last parameter of a function has type ...T, 
// it can be called with any number of trailing arguments of type T.
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "another return value"
}