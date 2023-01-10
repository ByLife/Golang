package Groupie

import (
	"encoding/json"
	"net/http"
	"reflect"
)

type Artist struct {
	Page        int
	HasSearched bool
	DataArtist  struct {
		All ArtistStruct
	}
	DataArtists struct {
		All []ArtistStruct
	}
	DataDatesLocations struct {
		All DatesLocations
	}
	Dates struct {
		All Dates
	}
}

type ArtistStruct struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	ImgURI       string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
	Page         int
}

type DatesLocations struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Dates struct {
	AlbumDates []string `json:"dates"`
}

func GetFromApi(_URI string, target interface{}) error {
	res, err := http.Get(_URI)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

func HasBeenChecked(a []ArtistStruct, b ArtistStruct) bool {
	for _, v := range a {
		if reflect.DeepEqual(v, a) {
			return true
		}
	}
	return false
}
