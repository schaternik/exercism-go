package scrabble

import "strings"

// Score calculates score of the word depending on the letters score.
func Score(word string) int {
	if word == "" {
		return 0
	}

	var result int

	for _, letter := range strings.ToLower(word) {
		result += scores[letter]
	}

	return result
}
