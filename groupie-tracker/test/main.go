package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//a17c554dedf7c0717a171299d554f609
	//http://api.positionstack.com/v1/
	baseURL, _ := url.Parse("http://api.positionstack.com")

	baseURL.Path += "v1/forward"

	params := url.Values{}
	params.Add("access_key", "a17c554dedf7c0717a171299d554f609")
	params.Add("query", "california usa")
	params.Add("output", "json")
	params.Add("limit", "1")

	baseURL.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", baseURL.String(), nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
