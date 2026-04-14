package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	wordsBytes, err := os.ReadFile("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	wordsString := string(wordsBytes)
	words := strings.Split(wordsString, ", ")
	words[len(words)-1], _ = strings.CutSuffix(words[len(words)-1], " ")

	realWord := "cycle"
	testedLetters := make(TestedLetters)

	testedLetters.AddWord(realWord, "crazy")
	validWords := testedLetters.GetValidWords(words)

	fmt.Println(validWords)
	fmt.Println(testedLetters)
}
