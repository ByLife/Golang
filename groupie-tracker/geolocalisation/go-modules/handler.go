package Groupie

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// API URI Handler

func URIResolve(URI string, errToSend error) {
	for {
		out, _ := exec.Command("ping", API_BaseLink).Output()
		fmt.Println(string(out))
		if !strings.Contains(string(out), "TTL=") {
			log.Fatal(errToSend)
		}
		fmt.Println("Successfull ping to " + API_BaseLink + " !")
		time.Sleep(time.Second * 30)
	}
}
