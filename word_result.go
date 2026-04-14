package main

import (
	"fmt"
	"strings"
)

type WordResult struct {
	word          string
	letterResults []LetterResult
}

type LetterResult int

const (
	wrong LetterResult = iota
	inWrongSpot
	correct
)

func getWordResult(actualWord, guessedWord string) WordResult {
	wr := WordResult{
		word: guessedWord,
	}

	for i, letter := range guessedWord {
		if byte(letter) == actualWord[i] {
			wr.letterResults = append(wr.letterResults, correct)
		} else if strings.Contains(actualWord, string(letter)) {
			wr.letterResults = append(wr.letterResults, inWrongSpot)
		} else {
			wr.letterResults = append(wr.letterResults, wrong)
		}
	}

	return wr
}

func getValidWords(wr WordResult, words []string) {
	incorrectIndices := []int{}
	correctIndices := []int{}
	inWrongSpotIndices := []int{}
	for i, res := range wr.letterResults {
		switch res {
		case wrong:
			incorrectIndices = append(incorrectIndices, i)
		case inWrongSpot:
			inWrongSpotIndices = append(inWrongSpotIndices, i)
		default:
			correctIndices = append(correctIndices, i)
		}
	}

	for _, word := range words {
		isValid := true
		for _, incorrectIndex := range incorrectIndices {
			if strings.Contains(word, string(wr.word[incorrectIndex])) {
				isValid = false
			}
		}
		for _, inWrongSpotIndex := range inWrongSpotIndices {
			if !strings.Contains(word, string(wr.word[inWrongSpotIndex])) || word[inWrongSpotIndex] == wr.word[inWrongSpotIndex] {
				isValid = false
			}
		}
		for _, correctIndex := range correctIndices {
			if wr.word[correctIndex] != word[correctIndex] {
				isValid = false
			}
		}

		if isValid {
			fmt.Println(word)
		}
	}

}
