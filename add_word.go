package main

import "strings"

func (tl TestedLetters) AddWord(realWord, guessedWord string) {
	for i, letterRune := range guessedWord {
		letter := string(letterRune)
		if _, ok := tl[letter]; !ok {
			tl[letter] = LetterState{}
		}

		if letter == string(realWord[i]) {
			temp := tl[letter]
			temp.result = correct
			temp.isFoundAt = append(temp.isFoundAt, i)
			tl[letter] = temp
		} else if strings.Contains(realWord, letter) {
			if tl[letter].result != correct {
				temp := tl[letter]
				temp.result = inWrongSpot
				temp.isNotFoundAt = append(temp.isNotFoundAt, i)
				tl[letter] = temp
			}
		} else {
			temp := tl[letter]
			temp.result = wrong
			temp.isNotFoundAt = []int{0, 1, 2, 3, 4}
			tl[letter] = temp
		}
	}
}
