package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", greetUser)
	http.ListenAndServe(":8080", nil)
}

func greetUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello User!"
	fmt.Fprintln(writer, greeting)
}
