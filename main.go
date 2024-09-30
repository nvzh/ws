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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//tpl, err := template.ParseFS(templatesFS, "templates/home.html")
	tpl, err := template.ParseFS(templatesFS, "templates/index.html")

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

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFS(templatesFS, "templates/about.html")

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

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFS(templatesFS, "templates/contact.html")

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

func priceHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFS(templatesFS, "templates/price.html")

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

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFS(templatesFS, "templates/projects.html")

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

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFS(templatesFS, "templates/services.html")

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

func sidebarHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFS(templatesFS, "templates/sidebar-right.html")

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

func homeHandler(w http.ResponseWriter, r *http.Request) {
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

	http.HandleFunc("/static/", staticFileHandler)

	//http.HandleFunc("/", homeHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/":
			indexHandler(w, r)
		case "/index.html":
			indexHandler(w, r)
		case "/about.html":
			aboutHandler(w, r)
		case "/contact.html":
			contactHandler(w, r)
		case "/price.html":
			priceHandler(w, r)
		case "/projects.html":
			projectsHandler(w, r)
		case "/services.html":
			servicesHandler(w, r)
		case "/sidebar-right.html":
			sidebarHandler(w, r)
		case "/home.html":
			homeHandler(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	fmt.Println("Starting server on http://127.0.0.1:80")

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
