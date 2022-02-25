package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"Indonesia",
	"Apple",
	"Programming",
	"Thunder",
	"Hangman",
	"Tourist",
	"Leonardo",
	"Semaphore",
	"Sunday",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := getRandomWord()
	fmt.Println(targetWord)
}

func getRandomWord() string {
	return dictionary[rand.Intn(len(dictionary))]
}
