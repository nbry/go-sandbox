package leetcode

import (
	"math"
)

func getColumnNumber(letter string) int {
	asciiValue := letter[0]

	return int(asciiValue) - 64
}

func valueForLetterAndPosition(index int, length int, letter string) int {
	power := float64(length - index - 1)
	c := getColumnNumber(letter)

	return int(math.Pow(26, power)) * c
}

func TitleToNumber(columnTitle string) int {
	length := len(columnTitle)
	result := 0

	for i, c := range columnTitle {
		result += valueForLetterAndPosition(i, length, string(c))
	}

	return result
}
