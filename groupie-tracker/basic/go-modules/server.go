package Groupie

import (
	"html/template"
	"net/http"
)

type Server struct {
}

func init() {
	WebHtml = template.Must(template.ParseGlob(Web_TemplatesDirectory))
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	key, err := r.URL.Query()["id"]
	if !err {
		ApiData.Page = 0
		GetFromApi(API_LinkToArtist, &ApiData.DataArtists.All)
	} else {
		ApiData.Page = 1
		GetFromApi(API_LinkToArtist+"/"+key[0], &ApiData.DataArtist.All)
		GetFromApi(ApiData.DataArtist.All.Relations, &ApiData.DataDatesLocations.All)
	}
	WebHtml.ExecuteTemplate(w, "index.html", ApiData)
}

func (s *Server) Init() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(Web_CssDirectory))))
	http.HandleFunc("/", s.index)
	http.ListenAndServe(Web_PORT, nil)
}
