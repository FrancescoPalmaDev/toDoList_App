package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", greetUser)
	http.HandleFunc("/", showForm)
	http.ListenAndServe(":8080", nil)
}

func showForm(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	html := `
		<form action="/greet" method="GET">
			<input type="text" name="name" placeholder="Enter your name">
			<button type="submit">Greet me</button>
		</form>
	`
	fmt.Fprintln(writer, html)
}
func greetUser(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		name = "User"
	}
	var greeting = "Hello " + name + "!"
	fmt.Fprintln(writer, greeting)
}
