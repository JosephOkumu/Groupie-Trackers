package main

import (
    "log"
    "net/http"
    "html/template"
)

var tmpl *template.Template

func main() {
    tmpl = template.Must(template.ParseGlob("templates/*.html"))

    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/artist", artistHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}