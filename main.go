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
	name := request.URL.Query().Get("name")

	if name == "" {
		name = "User"
	}
	var greeting = "Hello " + name + "!"
	fmt.Fprintln(writer, greeting)
}
