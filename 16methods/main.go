package main

import "fmt"

func main()  {
	fmt.Println("Struct in Golang")
	jerry := User{"Jerry", "jerry@gmail.com", true, 24}
	fmt.Println(jerry)
	fmt.Printf("Jerry's details are: %+v\n", jerry)
	fmt.Printf("Name is %v and Email is %v\n", jerry.Name, jerry.Email)

	jerry.GetStatus()
	jerry.NewMail()
}

// The first letters of the variables are capital because 
// we want anybody to access them
type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus () {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "newEmail@gmail.com"
	fmt.Println("Email of this user is: ", u.Email)
}
