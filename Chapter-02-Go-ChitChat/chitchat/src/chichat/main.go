package main

import (
	"html/template"
	"net/http"
)


func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data
}


func main() {
	p("ChiChat")
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	//mux.HandleFunc("/", index)
	server := &http.Server{
		Addr: "0.0.0.0:8080",
		Handler:mux,
	}
	server.ListenAndServe()
}
