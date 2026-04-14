package main

func (tl TestedLetters) ManuallyAddWord(word, letterStates string) {
	for i, letterRune := range word {
		letter := string(letterRune)
		if _, ok := tl[letter]; !ok {
			tl[letter] = LetterState{}
		}

		switch string(letterStates[i]) {
		case "2":
			temp := tl[letter]
			temp.result = correct
			temp.isFoundAt = append(temp.isFoundAt, i)
			tl[letter] = temp
		case "1":
			if tl[letter].result != correct {
				temp := tl[letter]
				temp.result = inWrongSpot
				temp.isNotFoundAt = append(temp.isNotFoundAt, i)
				tl[letter] = temp
			}
		case "0":
			temp := tl[letter]
			temp.result = wrong
			temp.isNotFoundAt = []int{0, 1, 2, 3, 4}
			tl[letter] = temp
		}
	}
}
