package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza")

	// comma ok || err ok
	// in Golang, we treat errors like variables, there is no try-catch
	// input, err := reader.ReadString('\n') // if we care about the error
	input, _ := reader.ReadString('\n') // keep on reading until there is a \n
	fmt.Println("Thanks for rating, ", input) // we can also use '+' instead of ','
	fmt.Printf("Type of the rating is %T ", input)
}
