package main

import (
    "log"
    "net/http"
    "html/template"
)

var tmpl *template.Template

func main() {
    var err error
    tmpl, err = template.ParseGlob("templates/*.html")
    if err != nil {
        log.Fatalf("Error parsing templates: %v", err)
    }

    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/artist", artistHandler)
    http.HandleFunc("/search", searchHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}