package main

import (
	"API"
	"log"
	"net/http"
)

func main() {
	s := &API.Server{}
	API.Hangman.Init()
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
