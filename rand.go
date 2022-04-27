package utils

import (
	"math/rand"
)

var (
	LowerLetterRunes = []rune("abcdefghijklmnopqrstuvwxyz")
	UpperLetterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	NumberRunes      = []rune("0123456789")

	LetterRunes []rune
)

func init() {
	LetterRunes = append(LetterRunes, LowerLetterRunes...)
	LetterRunes = append(LetterRunes, UpperLetterRunes...)
}

func RandStringLetter(n int) string {
	return RandStringWithSeed(n, LetterRunes)
}

func RandStringWithSeed(n int, seed []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = seed[rand.Intn(len(seed))]
	}
	return string(b)
}
