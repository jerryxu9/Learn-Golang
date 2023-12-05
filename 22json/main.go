package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"coursename"` // change the field name from Name to coursename when encoding to json
	Price int
	Platform string	`json:"website"`
	Password string `json:"-"` // - means the field will never be encoded
	Task []string `json:"tags,omitempty"` // omitempty means if the field is empty, we don't encode it to json
}


func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	courses := []course{
		{"React", 299, "Coursera", "abc123", []string{"web-dev", "js"}},
		{"MERN", 199, "Coursera", "bcd123", []string{"full-stack", "js"}},
		{"Angular", 299, "Coursera", "wsd123", nil},
	}

	// Enocde courses as JSON data

	// 1. this will work, but the printed finalJson is hard to read in terminal
	// finalJson, err := json.Marshal(courses)

	// 2. Formatted version
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonData := []byte(`
		{
			"coursename": "MERN",
			"Price": 199,
			"website": "Coursera",
			"tags": ["full-stack","js"]
        }
	`)

	var myCourse course

	// Checks whether data is a valid JSON encoding.
	checkValid := json.Valid(jsonData) 
	
	// 1. Decode into struct
	if checkValid {
		fmt.Println("JSON was valid")
		// Pass the reference b/c we want to store the info in myCourse instead of a copy
		json.Unmarshal(jsonData, &myCourse)
		fmt.Printf("%#v\n", myCourse)
	} else {
		fmt.Println("JSON was NOT valid")
	}

	// 2. Decode into key-value pair
	// use interface as value type since different data types can come in (e.g. int, string, slice)
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v, Value is %v, Value Type is %T\n", key, value, value)
	}
}