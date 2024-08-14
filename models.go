package main

import (
    "encoding/json"
    "net/http"
)

type Artist struct {
    ID          int      `json:"id"`
    Name        string   `json:"name"`
    Image       string   `json:"image"`
    Year        int      `json:"year"`
    FirstAlbum  string   `json:"first_album"`
    Members     []string `json:"members"`
}

type Location struct {
    Locations []string `json:"locations"`
}

type Date struct {
    Dates []string `json:"dates"`
}

type Relation struct {
    Index map[string]struct {
        Dates    []string `json:"dates"`
        Locations []string `json:"locations"`
    } `json:"index"`
}

func fetchAPI(url string, target interface{}) error {
    resp, err := http.Get(url)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    return json.NewDecoder(resp.Body).Decode(target)
}
