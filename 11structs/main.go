package main

import "fmt"

func main()  {
	fmt.Println("Struct in Golang")
	jerry := Usesr{"Jerry", "jerry@gmail.com", true, 24}
	fmt.Println(jerry)
	fmt.Printf("Jerry's details are: %+v\n", jerry)
	fmt.Printf("Name is %v and Email is %v\n", jerry.Name, jerry.Email)
}

// The first letters of the variables are capital because 
// we want anybody to access them
type Usesr struct {
	Name string
	Email string
	Status bool
	Age int
}