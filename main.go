package main

import (
	"fmt"
	"net/http"
)

var taskList = []string{}

func main() {
	http.HandleFunc("/add", addToList)
	http.HandleFunc("/", showList)
	http.ListenAndServe(":8080", nil)
}

func addToList(writer http.ResponseWriter, request *http.Request) {
	task := request.URL.Query().Get("task")
	taskList = append(taskList, task)

	http.Redirect(writer, request, "/", http.StatusSeeOther)
}

func showList(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")

	listItems := ""
	for _, task := range taskList {
		listItems += "<li>" + task + "</li>"
	}

	html := `
		<form action="/add" method="GET">
			<input type="text" name="task" placeholder="Enter a task">
			<button type="submit">Add task</button>
		</form>
		<ul>` + listItems + `</ul>
	`
	fmt.Fprintln(writer, html)
}
