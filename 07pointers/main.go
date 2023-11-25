package main

import "fmt"

func main()  {
	fmt.Println("Welcome to a class on pointers")

	var ptr *int
	fmt.Println("Value of ptr is: ", ptr)

	myNumber := 23
	var ptr2 = &myNumber
	fmt.Println("Address that ptr2 points to is: ", ptr2)
	fmt.Println("Value of ptr2 is: ", *ptr2)

	*ptr2 = *ptr2 * 2
	fmt.Println("New value of myNumber is: ", myNumber)
}