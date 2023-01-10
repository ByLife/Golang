package API

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Server struct{}

var CommandsToGet []string = []string{
	"help",
	"reset",
	"input",
	"gethangman",
	"save",
	"getword",
	"getattempt",
}

var HelpMessage string = `{"help":"Hello and Welcome to the Ynov Hangman ! \nTo interact with our API, send POST Request with the methods:\nhelp: What you did right now ... \nreset: To reset the actual hangman game \ninput: Enter your answer (Can be a word or a letter) \ngetHangMan: Get the actual ASCII Art of the hangman \nsave: To save the session \n"}`

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write(GetResponse(r))
	case "POST":
		reqBody, _ := ioutil.ReadAll(r.Body)
		w.Write(PostResponse(string(reqBody)))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func GetResponse(r *http.Request) []byte {
	var js string
	js += `{"cmd":"` + r.FormValue("cmd") + `",`
	js += `"ent":"` + r.FormValue("ent") + `"}`
	return PostResponse(js)
}

func Response(cmd string) []byte {
	var res []byte
	for _, v := range cmd {
		res = append(res, byte(v))
	}
	return res
}

func PostResponse(cmd string) []byte {
	dict, err := GetMapPOST(Response(cmd))
	if err != nil {
		cmd = `{"err":"` + err.Error() + `"}`
	} else {
		if dict["cmd"] != nil {
			cmd = SwitchCommand(dict)
		}
	}
	return []byte(cmd)
}

func GetMapPOST(s []byte) (map[string]interface{}, error) {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(s, &jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}

func valideCmd(cmd string, other []string) bool {
	for _, v := range other {
		if cmd == v {
			return true
		}
	}
	return false
}

func ToLower(s string) string {
	a := []rune(s)

	for i := range a {
		if a[i] >= 65 && a[i] <= 90 {
			a[i] = rune(a[i] + 32)
		}
	}
	return string(a)
}

func SwitchCommand(cmd map[string]interface{}) string {
	var final string
	if valideCmd(ToLower(cmd["cmd"].(string)), CommandsToGet) {
		switch ToLower(cmd["cmd"].(string)) {
		case "reset":
			final = `{"msg":"ok"}`
			Hangman.Reset()
		case "input":
			if cmd["ent"] != nil {
				final = Hangman.Input(cmd["ent"].(string))
			} else {
				final = `{"err":"ent is nil"}`
			}
		case "getattempt":
			return Hangman.GetAttempt()
		case "help":
			final = HelpMessage
		case "getword":
			final = Hangman.GetWord()
		case "gethangman":
			final = Hangman.GetHangMan()
		default:
			final = `{"msg":"Commande pas encore implemente"}`
		}
	} else {
		final = `{"msg":"Commande non existante"}`
	}
	return final
}
