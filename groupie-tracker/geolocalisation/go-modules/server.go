package Groupie

import (
	"html/template"
	"net/http"
	"net/url"
)

type Server struct {
}

func init() {
	WebHtml = template.Must(template.ParseGlob(Web_TemplatesDirectory))
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	key, err := r.URL.Query()["id"]
	_, err2 := r.URL.Query()["location"]
	ApiData = Artist{}
	if !err && !err2 {
		ApiData.Page = 0
		GetFromApi(API_LinkToArtist, &ApiData.DataArtists.All)
	} else if err && !err2 {
		ApiData.Page = 1
		GetFromApi(API_LinkToArtist+"/"+key[0], &ApiData.DataArtist.All)
		GetFromApi(ApiData.DataArtist.All.Relations, &ApiData.DataDatesLocations.All)
	} else if !err && err2 {
		ApiData.Page = 2
		params := url.Values{}
		GPSLocation, addrBool := r.URL.Query()["address"]
		if addrBool {
			if len(GPSLocation) > 0 {

				params.Add("access_key", GPSAccessKey)
				params.Add("query", ClearSpecialChars(GPSLocation[0]))
				params.Add("output", "json")
				params.Add("limit", "1")

				GPSBaseURL.RawQuery = params.Encode()
				GPSDataReq := GPSData{}
				GetFromApi(GPSBaseURL.String(), &GPSDataReq.Data)
				ApiData.DataGPS.Lat, ApiData.DataGPS.Long, ApiData.DataGPS.Addr = GPSDataReq.Data["data"][0]["latitude"].(float64), GPSDataReq.Data["data"][0]["longitude"].(float64), ClearSpecialChars(GPSLocation[0])
			}
		}
	}
	WebHtml.ExecuteTemplate(w, "index.html", ApiData)
}

func (s *Server) Init() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(Web_CssDirectory))))
	http.HandleFunc("/", s.index)
	http.ListenAndServe(Web_PORT, nil)
}
