package main

import "fmt"

func main()  {
	fmt.Println("Welcome to array in Golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("The fruit list is: ", fruitList)
	fmt.Println("The length of the fruit list array is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("The veggie list is: ", vegList)
	fmt.Println("The length of the veggie list array is: ", len(vegList))
}