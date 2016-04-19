package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
  port := ":8080"
  http.HandleFunc("/", RootHandler)
  fmt.Printf("### Server started on %s ###", port)
  http.ListenAndServe(port, nil)
}

func RootHandler(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(rw, "Error: %s", err)
	}

	t.Execute(rw, nil)
}
