package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Learning slices!")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:len(fruitList)])
	fmt.Println(fruitList)

	highScore := make([]int, 4)
	highScore[0] = 123
	highScore[1] = 82141
	highScore[2] = 2123
	highScore[3] = 999
	fmt.Println(highScore)

	// highScore[4] = 134	// this will give "panic: runtime error: index out of range [4] with length 4"
	// fmt.Println(highScore)

	highScore = append(highScore, 111, 222) // append method works without error
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println("After sorting: ", highScore)


	// Learn how to remove a value from slices based on index
	courses := []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	index := 2
	// remove swift from courses
	courses = append(courses[:index], courses[index + 1:]...)
	fmt.Println(courses)
}