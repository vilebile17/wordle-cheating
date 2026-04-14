package main

type TestedLetters map[string]LetterState

type LetterState struct {
	result       LetterResult
	isFoundAt    []int
	isNotFoundAt []int
}

type LetterResult int

const (
	wrong LetterResult = iota
	inWrongSpot
	correct
)
