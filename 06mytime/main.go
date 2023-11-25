package main

import (
	"fmt"
	"time"
)

func main()  {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	// Weird thing in Go: we must use this exact time to get the current time and format it
	// https://pkg.go.dev/time#Time.Format
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	// Create a custom date
	createdDate := time.Date(2020, time.December, 9, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}