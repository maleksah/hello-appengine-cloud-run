package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	log.Printf("The server started successfully and is listening for HTTP requests on port %s...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
