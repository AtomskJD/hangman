package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	"unicode"
)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"Indonesia",
	"Apple",
	"Programming",
	"Thunder",
	"Hangman",
	"Mountain",
	"Tourist",
	"Leonardo",
	"Semaphore",
	"Sunday",
	"Golang Is Awesome going",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWord()

	guessedLetters := initializeGuessedWords(targetWord)
	hangmanState := 0
	printGameState(targetWord, guessedLetters, hangmanState)
}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}

func printGameState(
	targetWord string,
	guessedLetters map[rune]bool,
	hangmangState int,
) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(getDrawing(hangmangState))
}

func getWordGuessingProgress(
	targetWord string,
	guessedLetters map[rune]bool,
) string {
	var result string = ""
	for _, ch := range targetWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			result += string(ch)
		} else {
			result += "_"
		}
		result += " "
	}

	return result
}

func getDrawing(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("slides/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}
	return string(data)
}
