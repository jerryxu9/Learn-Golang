package main

import (
	"fmt"
	"net/url"
)
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"


func main() {
	fmt.Println("Welcome to handling URLs in Go")
	fmt.Println(myurl)

	// 1. Parsing the URL
	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println("")

	queryParams := result.Query()
	fmt.Printf("The type of queryParams is: %T\n", queryParams)

	// Get query param directly
	fmt.Println(queryParams["coursename"])
	fmt.Println(queryParams["paymentid"])

	// Loop through all query params
	for _, val := range queryParams {
		fmt.Println("Param value is: ", val)
	}

	// 2. Constructing an URL
	// Important: we need to use reference
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawQuery: "user=jerry",
	}

	constructedUrl := partsOfUrl.String()
	fmt.Println("The constructed URL is: ", constructedUrl)
}