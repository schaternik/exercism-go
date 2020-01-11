package isogram

import "strings"

// IsIsogram check is string contains only uniq chars or not
func IsIsogram(input string) bool {
	var checkedLettes = make(map[rune]int)
	var str = strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(input), "-", ""), " ", "")

	for _, char := range str {
		if val, ok := checkedLettes[char]; ok && val >= 1 {
			return false
		}
		checkedLettes[char]++
	}

	return true
}
