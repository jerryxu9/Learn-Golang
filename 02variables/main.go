package main

import "fmt"

// jwtToken := "adfaf" // this is not allowed
var jwtToken1 = 3000	// option 1
var jwtToken2 int = 3000	// option 2

const LoginToken string = "adsfasf" 	// Public variable: first letter is captial letter


func main()  {	
	var username string = "jerry"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn);
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallFloat float32 = 234.12412325
	fmt.Println(smallFloat);
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var anotherIntVariable int
	fmt.Println(anotherIntVariable);
	fmt.Printf("Variable is of type: %T \n", anotherIntVariable)

	var anotherStrVariable string
	fmt.Println(anotherStrVariable);
	fmt.Printf("Variable is of type: %T \n", anotherStrVariable)

	// implict type
	var website = "adsfsa"
	fmt.Println(website)

	// no var style, this is only allowed in a function, not outside
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)

	// public variable 
	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}