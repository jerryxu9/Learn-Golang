package main

import "fmt"

func main()  {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	// Delete a key
	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// Loop through the map
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}