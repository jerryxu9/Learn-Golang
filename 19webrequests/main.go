package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Cat Facts API: https://catfact.ninja/#/Facts
const url = "https://catfact.ninja/fact"

func main() {
	fmt.Println("Welcome to web requests")
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	} 

	// The caller must close the response body when finished with it
	defer resp.Body.Close()
	fmt.Printf("Response is of type: %T\n", resp)

	dataBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println(content)
	
	// Parsing JSON response
	// https://rakaar.github.io/posts/2021-04-23-go-json-res-parse/
	var jsonRes map[string]interface{} // declaring a map for key names as string and values as interface 
	_ = json.Unmarshal(dataBytes, &jsonRes) // Unmarshalling

	fmt.Println(jsonRes)
	fmt.Println(jsonRes["fact"])
}