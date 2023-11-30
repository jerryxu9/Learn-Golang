package main

import "fmt"

func main()  {
	fmt.Println("Welcome to loops in Golang\n")
	
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)
	fmt.Println()

	fmt.Println("First way of for loops")
	// 1.
	for i := 0; i < len(days); i++ {
		fmt.Print(days[i] + " ")
	}

	fmt.Println("\n\nSecond way of for loops")
	// 2.
	for i := range days {
		fmt.Print(days[i] + " ")
	}

	fmt.Println("\n\nThird way of for loops")
	// 3.
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}
	// if we are not interested in index, use _:
	// for _, day := range days {
	// 	fmt.Printf("Value is %v\n", day)
	// }
	fmt.Println()

	// 4. break and continue
	rougueValue := 1

	for rougueValue < 5 {
		if rougueValue == 2{
			goto myLabel
		}

		if rougueValue == 3 {
			break

			// rougueValue++
			// continue
		}


		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}


	// 5. goto
	myLabel: 
		fmt.Println("Jumped to here")
}
