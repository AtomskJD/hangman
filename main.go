package main

import (
	"fmt"
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

	targetWord = "Awesome GO Going"
	guessedLetters := initializeGuessedWords(targetWord)
	printGameState(targetWord, guessedLetters)
	// guessedLetters['o'] = true
	// printGameState(targetWord, guessedLetters)

	fmt.Println(targetWord)
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

func printGameState(targetWord string, guessedLetters map[rune]bool) {
	for _, ch := range targetWord {
		if ch == ' ' {
			fmt.Print(" ")
		} else if guessedLetters[unicode.ToLower(ch)] == true {
			fmt.Print(string(ch))
		} else {
			fmt.Print("_")
		}
		fmt.Print(" ")

	}
	fmt.Println()
}
