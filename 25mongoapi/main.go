package main

import (
	"fmt"
	"mongoapi/router"
	"net/http"
)

func main() {
	r := router.Router()
	http.ListenAndServe(":4000", r)
	fmt.Println("Listening at port 4000")
}
