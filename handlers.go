package main

import (
    "net/http"
    "log"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    var artists []Artist
    err := fetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
    if err != nil {
        http.Error(w, "Failed to load artists", http.StatusInternalServerError)
        return
    }
    err = tmpl.ExecuteTemplate(w, "index.html", artists)
    if err != nil {
        log.Println(err)
    }
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Artist ID is required", http.StatusBadRequest)
        return
    }

    var artist Artist
    err := fetchAPI("https://groupietrackers.herokuapp.com/api/artists"+id, &artist)
    if err != nil {
        http.Error(w, "Failed to load artist", http.StatusInternalServerError)
        return
    }

    err = tmpl.ExecuteTemplate(w, "artist.html", artist)
    if err != nil {
        log.Println(err)
    }
}
