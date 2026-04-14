package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
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

	testedLetters := make(TestedLetters)
	reader := bufio.NewReader(os.Stdin)
	i := rand.Intn(len(words))
	wordToTry := words[i]
	fmt.Printf("Try the word '%v'\n", wordToTry)

	for {
		fmt.Print("Enter the result: ")
		result, _ := reader.ReadString('\n')

		if result == "22222" {
			fmt.Println("good job! you are a good cheater")
		}

		testedLetters.ManuallyAddWord(wordToTry, result)
		validWords := testedLetters.GetValidWords(words)

		if len(validWords) == 0 {
			fmt.Println("Err, some error occured. I have ran out of words...")
			break
		}
		if len(validWords) == 1 {
			fmt.Printf("Try the word '%v' and that should be the final answer!\n", validWords[0])
			break
		}

		i = rand.Intn(len(validWords))
		wordToTry = validWords[i]
		fmt.Printf("Try the word '%v'\n", wordToTry)
	}
}
