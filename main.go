package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		panic(err)
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", homeHandler)
	fmt.Println("Starting server on http://127.0.0.1:80")
	http.ListenAndServe(":80", nil)
}
