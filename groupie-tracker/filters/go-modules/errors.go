package Groupie

import "errors"

// Negligable but just to call easily the Error Codes with VSCode

var (
	PortAuthorisationProblem = "Port authorisation problem, allow firewall ?"
	URINotResolving          = "URI is not resolving, conneciton problem ?"
	APINotResponding         = "Api is not responding, connection or API down ?"
)

// Map, key => string and values => error to handle errors

var AllErrors map[string]error = map[string]error{
	"PortAuthorisationProblem": errors.New("Can't open port on " + Web_PORT + ", authorisation problem ?"),
	"URINotResolving":          errors.New("URI isn't responding, lookup at: " + API_BaseLink + ", maybe connection problem ?"),
}
