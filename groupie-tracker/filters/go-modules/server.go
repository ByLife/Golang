package Groupie

import (
	"html/template"
	"net/http"
	"strconv"
)

type Server struct {
}

func init() {
	WebHtml = template.Must(template.ParseGlob(Web_TemplatesDirectory))
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	key, err := r.URL.Query()["id"]
	searchGetVal, err2 := r.URL.Query()["date"]
	searchGetValMembers, numMemb := r.URL.Query()["nummembers"]
	if !err {
		ApiDataTemp := Artist{}
		ApiData.DataArtists.All = ApiDataTemp.DataArtists.All
		ApiData.Page = 0
		GetFromApi(API_LinkToArtist, &ApiDataTemp.DataArtists.All)
		if err2 && len(searchGetVal) > 0 {
			dateValue, atoiErr := strconv.Atoi(searchGetVal[0])
			if atoiErr == nil {
				for _, v := range ApiDataTemp.DataArtists.All {
					if !numMemb {
						if v.CreationDate >= dateValue {
							ApiData.DataArtists.All = append(ApiData.DataArtists.All, v)
						}
					} else if numMemb && len(searchGetValMembers) > 0 {
						numMembers, errNumMemb := strconv.Atoi(searchGetValMembers[0])
						if errNumMemb == nil {
							if v.CreationDate >= dateValue && len(v.Members) >= numMembers {
								ApiData.DataArtists.All = append(ApiData.DataArtists.All, v)
							}
						}
					}
				}
			}
			ApiData.HasSearched = true

		} else {
			ApiData.DataArtists.All = ApiDataTemp.DataArtists.All
			ApiData.HasSearched = false
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
