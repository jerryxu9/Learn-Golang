package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup // this is normally a pointer
var mut sync.Mutex    // this is normally a pointer

func main() {
	// go is a keyword used to create goroutine
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		log.Fatal(err)
	} else {
		// Allow only one goroutine to modify signals at a time
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for website %s\n", res.StatusCode, endpoint)
	}
}
