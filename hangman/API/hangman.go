package API

/*
reset
input

*/

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"
)

var Hangman = &HangManData{}

/*
Partie POO
Par Thomas
*/
type HangManData struct {
	Word             string
	ToFind           string
	Attempts         int
	HangmanPositions [10][]string
	Error            bool
	LetterUtils      []rune
}

func (h *HangManData) CheckLetter(input rune) bool {
	/*
		Permet de Check les lettres deja utilises et return
	*/
	for _, v := range h.LetterUtils {
		if input == v {
			return true
		}
	}
	h.LetterUtils = append(h.LetterUtils, input)
	return false
}

func (h *HangManData) GetWord() string {
	return `{"msg":"` + h.Word + `"}`
}

func (h *HangManData) Init() {
	h.GenerateHangMan()
	//init ToFind
	h.ToFind = GetRandWord(SplitByReturns(ReadFile("Pendu.txt")))
	h.GenerateWord()
	fmt.Println(h.ToFind)
}

func (h *HangManData) Reset() {
	h.Word = ""
	h.ToFind = GetRandWord(SplitByReturns(ReadFile("Pendu.txt")))
	h.GenerateWord()
	h.Attempts = 0
	h.Error = false
	h.LetterUtils = nil
	fmt.Println(h.ToFind)
}

func (h *HangManData) GoodWord(s string) {
	/*
		Permet de Genére le Mot à troue de Base
	*/
	if s+"\r" == h.ToFind {
		h.Word = s
		h.Error = false
	} else {
		h.Attempts += 2
		h.Error = true
	}
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

func (h *HangManData) Print() string {
	/*
		permet d'afficher le text du jeu et le hangman en fonction de la dernière entrer
	*/
	var res string
	if h.Error {
		res += `{"msg":"Not present in the word ` + strconv.Itoa(10-h.Attempts) + ` attempts remaining",`
		res += `"Hangman":"`
		for _, i := range h.HangmanPositions[h.Attempts-1] {
			res += i + `\n`
		}
		res += `"}`
	} else {
		res += `{"msg":"`
		for _, i := range h.Word {
			res += string(i) + " "
		}
		res += `"}`
	}
	return res
}

func (h *HangManData) GetHangMan() string {
	var res string
	res += `{"Hangman":"`
	if h.Attempts > 0 {
		for _, i := range h.HangmanPositions[h.Attempts-1] {
			res += i + `\n`
		}
		res += `"}`
	} else {
		res = `{"Hangman":""}`
	}
	return res
}

func (h *HangManData) Win() (string, bool) {
	/*
		Check les conditions de fin du jeu
	*/
	var res string
	for _, i := range h.Word {
		if i == '_' && h.Attempts >= 10 {
			res += `{"msg":"You Lost ^^"}`
			return res, true
		} else if i == '_' {
			return "", false
		}
	}
	res += `{"msg":"Congrats !"}`
	return res, true
}

func (h *HangManData) GetAttempt() string {
	return `{"msg":"` + string(h.Attempts+48) + `"}`
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

func hangPos(cont string, nbError int) []string {
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
	var final []string
	var total int
	for _, v := range cont {
		if v == '\n' {
			total++
		}
		if total >= start && total <= end {
			if v != '\n' && v != '\r' {
				result += string(v)
			} else if v == '\n' {
				final = append(final, result)
				result = ""
			}
		}
	}
	if nbError == 0 {
		result += "\n"
	}
	return final
}

func (h *HangManData) Input(str string) string {
	out := []rune(str)
	if len(str) <= 1 {
		if !h.CheckLetter(out[0]) {
			h.GoodLetter(out[0])
			win_msg, is_win := h.Win()
			if is_win {
				h.Reset()
				return win_msg
			} else {
				return h.Print()
			}
		} else {
			return `{"msg":"The letter (` + string(out[0]) + `) entered has been already used, you retard ^^ !"}`
		}
	} else {
		h.GoodWord(str)
		win_msg, is_win := h.Win()
		if is_win {
			h.Reset()
			return win_msg
		} else {
			return h.Print()
		}
	}
}
