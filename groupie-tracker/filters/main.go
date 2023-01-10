package main

import (
	Groupie "Groupie/go-modules"
	"fmt"
	"strconv"
)

func main() {
	s := Groupie.Server{}
	for i := 1; i < 7; i++ {
		Groupie.ApiData.NumberOfMembers = append(Groupie.ApiData.NumberOfMembers, strconv.Itoa(i))
	}
	fmt.Printf("Running on %s, API connected to %s. %c", Groupie.Web_PORT, Groupie.API_BaseLink, '\n')
	//go Groupie.URIResolve(Groupie.API_BaseLink, Groupie.AllErrors[Groupie.URINotResolving])
	s.Init()
}
