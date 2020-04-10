package isogram

import "strings"

// IsIsogram check is string contains only uniq chars or not
func IsIsogram(input string) bool {
	checkedLettes := map[rune]int{}
	lowercaseStr := strings.ToLower(input)

	for _, char := range lowercaseStr {
		if char == ' ' || char == '-' {
			continue
		}
		if val, ok := checkedLettes[char]; ok && val >= 1 {
			return false
		}
		checkedLettes[char]++
	}

	return true
}
