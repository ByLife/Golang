package Groupie

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artist struct {
	Page       int
	DataArtist struct {
		All ArtistStruct
	}
	DataArtists struct {
		All []ArtistStruct
	}
	DataDatesLocations struct {
		All DatesLocations
	}
	DataGPS struct {
		Long float64
		Lat  float64
		Addr string
	}
}

type GPSData struct {
	Data map[string][]map[string]interface{}
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

func GetFromApi(_URI string, target interface{}) error {
	res, err := http.Get(_URI)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	err2 := json.NewDecoder(res.Body).Decode(target)
	if err2 == nil {
		return json.NewDecoder(res.Body).Decode(target)
	}
	fmt.Println(err2)
	return nil
}

func ClearSpecialChars(str string) string {
	var final string
	for _, v := range str {
		if v != '-' && v != '_' {
			final += string(v)
		} else {
			final += string(" ")
		}
	}
	return final
}
