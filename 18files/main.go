package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")

	content := "This needs to go in a file"

	// Create a file
	file, err := os.Create("./myFile.txt")

	// if err != nil {
	// 	panic(err)
	// } 
	checkNilError(err)

	// Write to the file
	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is: ", length)
	
	// Close the file. It is recommanded to use defer keyword
	defer file.Close()

	readFile("./myFile.txt")
}

func readFile(fileName string)  {
	dataByte, err := os.ReadFile(fileName)
	checkNilError(err)

	fmt.Println("Text data in the file is: \n", string(dataByte))
}


func checkNilError(err error)  {
	if err != nil {
		panic(err)
	}
}