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
	words[len(words)-1] = words[len(words)-1][:len(words[len(words)-1])-1] // umm, this just removes the whitespace at the end ¯\_(ツ)_/¯

	realWord := "great"
	testedLetters := make(TestedLetters)

	testedLetters.addWord(realWord, "crane")
	testedLetters.addWord(realWord, "treat")
	validWords := testedLetters.getValidWords(words)

	fmt.Println(validWords)
	fmt.Println(testedLetters)
}
