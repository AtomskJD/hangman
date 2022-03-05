// Extra Challenge:
//  * Build the game & distribute it to your friend
//    * Build a binary for another distro as well - e.g. Windows
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)

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
var hintAttempts int = 2

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWord()

	guessedLetters := initializeGuessedWords(targetWord)
	hangmanState := 0
	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)
		input := readInput()
		if input == "hint" {
			if hintAttempts > 0 {
				hintAttempts--
				hintLetter := runHint(targetWord, &guessedLetters)
				fmt.Printf("your hint is %s, you have %d hint attenpt\n", hintLetter, hintAttempts)
			} else {
				fmt.Println("You out of hints!")
			}
			continue
		} else if len(input) != 1 {
			fmt.Println("Err: invalid input. Please use letters only.")
			continue
		}
		letter := rune(input[0])

		if isLetterGuessed(guessedLetters, letter) {
			fmt.Printf("the letter %s already guessed try another\n", letter)
		} else if isCorrectGuess(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}

	}

	printGameState(targetWord, guessedLetters, hangmanState)
	fmt.Println("GAME OVER")
	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You lose.")
	} else {
		panic("invalid state.")
	}
}

func initializeGuessedWords(targetWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(targetWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(targetWord[len(targetWord)-1]))] = true

	return guessedLetters
}

func runHint(targetWord string, guessedLetters *map[rune]bool) string {
	runeWord := []rune(targetWord)
	var otherRunes []rune
	for _, l := range runeWord {
		if (*guessedLetters)[unicode.ToLower(l)] != true {
			otherRunes = append(otherRunes, l)
		}
	}
	hintRune := otherRunes[rand.Intn(len(otherRunes))]
	(*guessedLetters)[hintRune] = true

	return string(hintRune)

}

func isGameOver(targetWord string,
	guessedLetters map[rune]bool,
	hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, ch := range targetWord {
		if !guessedLetters[unicode.ToLower(ch)] {
			return false
		}
	}

	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
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

func readInput() string {
	fmt.Print("> ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}
func isLetterGuessed(guessedLetters map[rune]bool, letter rune) bool {
	if guessedLetters[letter] == true {
		return true
	}
	return false
}
func isCorrectGuess(targetWord string, letter rune) bool {
	return strings.ContainsRune(targetWord, letter)
}
