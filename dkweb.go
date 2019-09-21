package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const (
	PORT = ":3000"
	HOST = "localhost"
)

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, nil)
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("books.html"))
	tmpl.Execute(w, nil)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/books", handleBooks)

	fmt.Println("Listening on " + HOST + PORT)
	http.ListenAndServe(":3000", nil)
}
