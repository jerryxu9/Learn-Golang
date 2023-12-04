package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to more web requests")

	// PerformGetRequests()
	// PerformPostJsonRequests()
	PerformPostFormRequest()
}

func PerformGetRequests() {
	const myurl = "http://localhost:8080/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}
	
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	// First way of reading thte content
	// content, _ := io.ReadAll(response.Body)
	// fmt.Println(string(content))

	// Second way
	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte Count is: ", byteCount)
	fmt.Println(responseString.String())
}


func PerformPostJsonRequests()  {
	const myurl = "http://localhost:8080/post"

	// Create request body
	requestBody := strings.NewReader(`
		{
			"user": "jerry",
			"location": {
				"country": "Canada",
				"school": "UBC",
				"age": 23
			}
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8080/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "Jerry")
	data.Add("lastname", "Xu")
	data.Add("email", "jerry@gmail.com")
	
	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}