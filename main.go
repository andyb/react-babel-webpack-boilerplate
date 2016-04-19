package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", fs))
	http.HandleFunc("/", RootHandler)
	log.Printf("### Server started on %s ###", port)
	http.ListenAndServe(port, nil)
}

func RootHandler(rw http.ResponseWriter, req *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "Error: %s", err)
	}

	t.Execute(rw, nil)
}
