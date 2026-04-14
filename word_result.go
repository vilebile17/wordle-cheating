package main

import (
	"strings"
)

type TestedLetters map[string]LetterState

type LetterState struct {
	result       LetterResult
	isFoundAt    int
	isNotFoundAt []int
}

type LetterResult int

const (
	wrong LetterResult = iota
	inWrongSpot
	correct
)

func (tl TestedLetters) addWord(realWord, guessedWord string) {
	for i, letterRune := range guessedWord {
		letter := string(letterRune)
		if _, ok := tl[letter]; !ok {
			tl[letter] = LetterState{}
		}

		if letter == string(realWord[i]) {
			temp := tl[letter]
			temp.result = correct
			temp.isFoundAt = i
			tl[letter] = temp
		} else if strings.Contains(realWord, letter) && tl[letter].result != correct {
			temp := tl[letter]
			temp.result = inWrongSpot
			temp.isNotFoundAt = append(temp.isNotFoundAt, i)
			tl[letter] = temp
		} else {
			temp := tl[letter]
			temp.result = wrong
			temp.isNotFoundAt = []int{0, 1, 2, 3, 4}
			tl[letter] = temp
		}
	}
}

func (tl TestedLetters) getValidWords(words []string) []string {
	validWords := []string{}

	for _, word := range words {
		isValid := true
		for letter := range tl {
			switch tl[letter].result {
			case correct:
				if string(word[tl[letter].isFoundAt]) != letter {
					isValid = false
				}
			case inWrongSpot:
				if !strings.Contains(word, letter) {
					isValid = false
				}
				for _, i := range tl[letter].isNotFoundAt {
					if string(word[i]) == letter {
						isValid = false
					}
				}
			case wrong:
				if strings.Contains(word, letter) {
					isValid = false
				}
			}
		}
		if isValid {
			validWords = append(validWords, word)
		}
	}

	return validWords
}
