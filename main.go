package main

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed static/*
var staticFS embed.FS

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//tpl, err := template.ParseFiles("templates/home.html")
	tpl, err := template.ParseFS(templatesFS, "templates/home.html")

	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Could not execute template", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
		return
	}
}

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := path.Join("static", r.URL.Path[len("/static/"):])

	// Read the static file from the embedded filesystem
	data, err := staticFS.ReadFile(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Determine the content type and serve the file
	http.ServeContent(w, r, filePath, time.Now(), bytes.NewReader(data))
}

func main() {
	// Serve static files from the "static" directory
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/static/", staticFileHandler)

	http.HandleFunc("/", homeHandler)

	fmt.Println("Starting server on http://127.0.0.1:80")

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
