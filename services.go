package main

import (
	"encoding/json"
	"net/http"
)

// FetchAPI function
func fetchAPI(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(target)

}

// Fetch and organize data
func FetchAndOrganizeData() ([]ArtistInfo, error) {
	var artists []Artist
	var locations []Location
	var dates []Date
	var relations []Relation

	// Fetch data from APIs
	if err := fetchAPI("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		return nil, err
	}
	if err := fetchAPI("https://groupietrackers.herokuapp.com/api/locations", &locations); err != nil {
		return nil, err
	}
	if err := fetchAPI("https://groupietrackers.herokuapp.com/api/dates", &dates); err != nil {
		return nil, err
	}
	if err := fetchAPI("https://groupietrackers.herokuapp.com/api/relation", &relations); err != nil {
		return nil, err
	}

	// Organize data
	var artistInfos []ArtistInfo
	for _, relation := range relations {
		artistInfo := ArtistInfo{
			Artist:   artists[relation.ID],
			Location: locations[relation.ID].Locations,
			Dates:    dates[relation.ID].Dates,
		}
		artistInfos = append(artistInfos, artistInfo)
	}

	return artistInfos, nil
}
