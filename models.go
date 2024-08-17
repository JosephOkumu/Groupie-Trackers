package main

import (
    "encoding/json"
    "net/http"
)

type Artist struct {
    ID          int      `json:"id"`
    Image       string   `json:"image"`
    Name        string   `json:"name"`
    Members     []string `json:"members"`
    CreationDate int      `json:"date"`
    FirstAlbum  string   `json:"first_album"`
    
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
    Dates []string `json:"dates"`
}

type Relation struct {
	ID            int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func fetchAPI(url string, target interface{}) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    return json.NewDecoder(resp.Body).Decode(target)
}
