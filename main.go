package main

import (
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

var templates *template.Template

func main() {
	// Parse all templates
	templates = template.Must(template.ParseFiles(
		filepath.Join("templates", "base.html"),
		filepath.Join("templates", "index.html"),
		filepath.Join("templates", "time.html"),
		filepath.Join("templates", "greet.html"),
	))

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/time", timeHandler)
	mux.HandleFunc("/greet", greetHandler)

	println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "time.html", map[string]string{
		"Now": time.Now().Format("15:04:05"),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	data := map[string]string{"Name": name}

	err := templates.ExecuteTemplate(w, "greet.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
