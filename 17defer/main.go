package main

import "fmt"

func main()  {
	/*
	In Golang, the defer keyword is used to delay the execution of a function
	until the surrounding function completes. The deferred function 
	calls are executed in Last-In-First-Out (LIFO) order.
	*/
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")

	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}