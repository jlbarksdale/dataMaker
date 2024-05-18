package generator

import (
	"math/rand"
	"strconv"
)

func GenerateString(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		s += string(rune(GenerateChar()))
	}
	return s
}
func GenerateChar() int {
	return rand.Intn(93) + 33
}

func GenerateInt(numberofDigits int) (int, error) {
	s := ""
	for i := 0; i < numberofDigits; i++ {
		s += strconv.Itoa(GenerateDigit())
	}
	return strconv.Atoi(s)
}

func GenerateDigit() int {
	return rand.Intn(10)
}

func GenerateBoolean() bool {
	val := rand.Intn(2)
	return val == 1
}
