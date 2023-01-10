package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
Partie POO
Par Thomas
*/
type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10]string
	Error            bool
	Run              bool
}

func (h *HangManData) GoodLetter(letter rune) {
	/*
		Permet de Verifier si la lettre est dans le mot h.ToFind
		Et de regenéré le h.Word si elle y est
	*/
	t := false
	for k, i := range h.ToFind {
		if i == letter {
			t = true
			var temp string
			for g, y := range h.Word {
				if g != k {
					temp += string(y)
				} else {
					temp += string(i)
				}
			}
			h.Word = temp
			h.Error = false
		}
	}
	if !t {
		h.Attempts++
		h.Error = true
	}
}

func (h *HangManData) GenerateWord() {
	/*
		Permet de Genére le Mot à troue de Base
	*/
	n := len(h.ToFind)/2 - 1
	var t []int
	var t2 []int
	for i := 0; i < n; i++ {
		t3 := rand.Intn(len(h.ToFind) - 1)
		for In(t3, t2) {
			t3 = rand.Intn(len(h.ToFind) - 1)
		}
		t = append(t, t3)
		t2 = append(t2, t3)
	}
	for k, i := range h.ToFind {
		if k+1 < len(h.ToFind) {
			if In(k, t) {
				h.Word += string(i)
			} else {
				h.Word += "_"
			}
		}
	}
}

func In(k int, t []int) bool {
	/*
		check if k is in t
	*/
	for _, i := range t {
		if k == i {
			return true
		}
	}
	return false
}

func (h *HangManData) GenerateHangMan() {
	/*
		Assigne les differents ascii art des hangman au tableau h.HangmanPositions
	*/
	t := ReadFile("hangman.txt")
	for i := 0; i < 10; i++ {
		h.HangmanPositions[i] = hangPos(t, i)
	}
}

func (h *HangManData) Print() {
	/*
		permet d'afficher le text du jeu et le hangman en fonction de la dernière entrer
	*/
	if h.Error {
		fmt.Print("Not present in the word, " + strconv.Itoa(10-h.Attempts) + " attempts remaining")
		fmt.Println(h.HangmanPositions[h.Attempts-1])
		if h.Attempts == 1 {
			fmt.Println()
		}
	} else {
		for _, i := range h.Word {
			fmt.Print(string(i) + " ")
		}
		fmt.Println()
		fmt.Println()
	}
}

func (h *HangManData) Win() {
	/*
		Check les conditions de fin du jeu
	*/
	for _, i := range h.Word {
		if i == '_' && h.Attempts >= 10 {
			h.Run = false
			fmt.Println("You Lost ^^")
			return
		} else if i == '_' {
			return
		}
	}
	fmt.Println("Congrats !")
	h.Run = false
}

/*
   Partie terminal / Input / CIN
   Par Luca
*/

func GetRandWord(word_array []string) string {
	/*
		Permet de Genére un index aléatoire
		dans la liste de mot possible
	*/
	rand.Seed(time.Now().UTC().UnixNano())
	return word_array[rand.Intn(len(word_array)-1)]
}

func GetInput() rune {
	/*
		Permet de récupérer la lettre
		passer dans le terminal
	*/
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose: ")
	input, _ := reader.ReadString('\n')
	temp := []rune(input)
	return temp[0]
}

/*
   Partie lecture fichier
   Par Léo
*/

func ReadFile(name string) string {
	/*
		Permet d'obtenir le contenu du fichier
		 et de savoir s'il existe ou pas sinon il renvoie une erreur
	*/
	cont, err := ioutil.ReadFile(name)
	if err != nil {
		panic("Error opening file, missing permissions ?")
	}
	return string(cont)
}

func SplitByReturns(cont string) []string {
	/*
		Permet d'interprêter et d'ajouter à un string
		chaques lignes voulues
	*/
	var str string
	var final []string
	for _, v := range cont {
		if v == 10 {
			final = append(final, str)
			str = ""
		} else {
			str += string(v)
		}
	}
	return final
}

func hangPos(cont string, nbError int) string {
	/*
		Permet d'obtenir dans un string le contenu du dessin du pendu
		selon le nombre d'erreurs
	*/
	nbr := 8
	if nbError == 0 {
		nbr = 7
	}
	start := nbError * nbr
	end := start + nbr
	var result string
	var total int
	for _, v := range cont {
		if v == '\n' {
			total++
		}
		if total >= start && total <= end-1 {
			result += string(v)
		}
	}
	if nbError == 0 {
		result += "\n"
	}
	return result
}

// =-=-=-=-=-=-= MAIN =-=-=-=-=-=-=

func main() {
	h := &HangManData{}
	h.ToFind = GetRandWord(SplitByReturns(ReadFile("Pendu.txt"))) // insert le mot générer aléatoirement dans la struct
	h.Run = true                                                  // indique au programme que la partie commance
	h.GenerateWord()
	h.GenerateHangMan()
	fmt.Println("Good Luck, you have 10 attempts.")
	for _, i := range h.Word { // affiche le mot a trou
		fmt.Print(string(i) + " ")
	}
	fmt.Println()
	fmt.Println()
	for h.Run {
		h.GoodLetter(GetInput())
		h.Print()
		h.Win()
	}
}
