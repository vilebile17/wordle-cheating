package main

import "strings"

func (tl TestedLetters) GetValidWords(words []string) []string {
	validWords := []string{}

	for _, word := range words {
		isValid := true
		for letter := range tl {
			switch tl[letter].result {
			case correct:
				for _, i := range tl[letter].isFoundAt {
					if string(word[i]) != letter {
						isValid = false
					}
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
