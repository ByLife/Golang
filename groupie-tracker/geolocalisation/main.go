package main

import (
	Groupie "Groupie/go-modules"
	"fmt"
)

func main() {
	s := Groupie.Server{}
	fmt.Printf("Running on %s, API connected to %s. %c", Groupie.Web_PORT, Groupie.API_BaseLink, '\n')
	Groupie.GPSBaseURL.Path += Groupie.GPSUrlPath
	//go Groupie.URIResolve(Groupie.API_BaseLink, Groupie.AllErrors[Groupie.URINotResolving])
	s.Init()
}
