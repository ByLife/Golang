package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func send_post_request(msg map[string]string) string {
	postBody, _ := json.Marshal(msg)
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:8080", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	return sb

}

func GetParam() []rune {
	/*
		Permet de récupérer la lettre ou le mot
		passer dans le terminal
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("input: ")
	input, _ := reader.ReadString('\n')
	if len(input) <= 1 {
		temp := []rune(input[:1])
		return temp
	} else {
		temp := []rune(input[:len(input)-2])
		return temp
	}
}

func GetInput() []rune {
	/*
		Permet de récupérer la lettre ou le mot
		passer dans le terminal
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Guess: ")
	input, _ := reader.ReadString('\n')
	if len(input) <= 1 {
		temp := []rune(input[:1])
		return temp
	} else {
		temp := []rune(input[:len(input)-2])
		return temp
	}
}

func main() {
	fmt.Println("Welcome to HANGMAN API client")
	fmt.Println()
	run := true

	i := (map[string]string{
		"cmd": "reset",
	})
	send_post_request(i)
	i = (map[string]string{
		"cmd": "getword",
	})
	s := send_post_request(i)[8:]
	s = s[:len(s)-2]
	fmt.Println(s)
	fmt.Println()
	for run {
		fmt.Println("possible command :")
		fmt.Println("1 - restart the game")
		fmt.Println("2 - make a guess")
		fmt.Println("3 - get the current word status")
		fmt.Println("4 - get the current hangman status")
		fmt.Println("5 - get the help pannel")
		fmt.Println("6 - exit the game")
		fmt.Println()
		s := string(GetParam())
		switch s {
		case "1":
			i := (map[string]string{
				"cmd": "reset",
			})
			send_post_request(i)
			i = (map[string]string{
				"cmd": "getword",
			})
			s := send_post_request(i)[8:]
			s = s[:len(s)-2]
			fmt.Println(s)
			fmt.Println()
		case "2":
			s := string(GetInput())
			i := (map[string]string{
				"cmd": "input", "ent": s,
			})
			fmt.Println(send_post_request(i))
			fmt.Println()
		case "3":
			i := (map[string]string{
				"cmd": "getword",
			})
			s := send_post_request(i)[8:]
			s = s[:len(s)-2]
			fmt.Println(s)
			fmt.Println()
		case "4":
			i := (map[string]string{
				"cmd": "gethangman",
			})
			fmt.Println(send_post_request(i))
			fmt.Println()
		case "5":
			i := (map[string]string{
				"cmd": "help",
			})
			fmt.Println(send_post_request(i))
			fmt.Println()
		case "6":
			run = false
		}
	}
}
