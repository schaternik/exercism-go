package scrabble

import (
	"strings"
)

// Score2 return scrabble score of the word
func Score2(word string) int {
	var result int
	for _, letter := range strings.ToLower(word) {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			result++
		case 'd', 'g':
			result += 2
		case 'b', 'c', 'm', 'p':
			result += 3
		case 'f', 'h', 'v', 'w', 'y':
			result += 4
		case 'k':
			result += 5
		case 'j', 'x':
			result += 8
		case 'q', 'z':
			result += 10
		}
	}

	return result
}
