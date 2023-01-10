package Groupie

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type Server struct {
}

func init() {
	WebHtml = template.Must(template.ParseGlob(Web_TemplatesDirectory))
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	key, err := r.URL.Query()["id"]
	searchGetVal, err2 := r.URL.Query()["s"]

	if !err {
		ApiData.Page = 0
		ApiDataTemp = Artist{}
		GetFromApi(API_LinkToArtist, &ApiDataTemp.DataArtists.All)
		if err2 && len(searchGetVal) > 0 {
			SearchValue = searchGetVal[0]
			if SearchValue != "" {
				ApiData = Artist{}
				ApiData.HasSearched = true
				for _, v := range ApiDataTemp.DataArtists.All {
					if !HasBeenChecked(ApiData.DataArtists.All, v) {
						if strings.Contains(strings.ToLower(v.Name+" - Artist/Band"), strings.ToLower(SearchValue)) {
							ApiData.DataArtists.All = append(ApiData.DataArtists.All, v)
						}
						for i := range v.Members {
							if strings.Contains(strings.ToLower(v.Members[i]+" - Member"), strings.ToLower(SearchValue)) {
								ApiData.DataArtists.All = append(ApiData.DataArtists.All, v)
							}
						}
						if strings.Contains(strings.ToLower(strconv.Itoa(v.CreationDate)+" - Creation Date"), strings.ToLower(SearchValue)) {
							ApiData.DataArtists.All = append(ApiData.DataArtists.All, v)
						}
					}
				}
			}
		} else {
			ApiData.HasSearched = false
			ApiData.DataArtists.All = ApiDataTemp.DataArtists.All
		}
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
