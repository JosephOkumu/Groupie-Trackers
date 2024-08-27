package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var artists []Artist
	err := fetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		http.Error(w, "Failed to load artists", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("Artist ID is required but not provided")
		http.Error(w, "Artist ID is required", http.StatusBadRequest)
		return
	}

	var artist Artist
	err := fetchAPI("https://groupietrackers.herokuapp.com/api/artists/"+id, &artist)
	if err != nil {
		log.Printf("Error fetching artist with ID %s: %v", id, err)
		http.Error(w, "Failed to load artist", http.StatusInternalServerError)
		return
	}

	var relation Relation
	err = fetchAPI("https://groupietrackers.herokuapp.com/api/relation/"+id, &relation)
	if err != nil {
		log.Printf("Error fetching relation data for artist with ID %s: %v", id, err)
		http.Error(w, "Failed to load relation data", http.StatusInternalServerError)
		return
	}

	// Combine artist and relation data into a struct
	data := struct {
		Artist   Artist
		Relation Relation
	}{
		Artist:   artist,
		Relation: relation,
	}

	// Pass combined data to the template
	err = tmpl.ExecuteTemplate(w, "artist.html", data)
	if err != nil {
		log.Printf("Error executing template for artist with ID %s: %v", id, err)
		http.Error(w, "Failed to render page", http.StatusInternalServerError)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Search query is required", http.StatusBadRequest)
		return
	}

	var artists []Artist
	err := fetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		http.Error(w, "Failed to load artists", http.StatusInternalServerError)
		return
	}

	// Filter artists based on the query
	var results []Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			results = append(results, artist)
		}
	}

	if len(results) == 0 {
		http.Error(w, "No artists found", http.StatusNotFound)
		return
	}

	// Load and parse the template file
	tmpl, err := template.ParseFiles("templates/results.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Failed to render results", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, results); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Failed to render results", http.StatusInternalServerError)
	}
}
