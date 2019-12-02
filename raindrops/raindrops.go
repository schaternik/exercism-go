package raindrops

import (
	"sort"
	"strconv"
)

var convertConfig = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert returns string depending on input's factor
func Convert(input int) string {
	var result string

	for _, factor := range factors() {
		if input%factor == 0 {
			result += convertConfig[factor]
		}
	}

	if result == "" {
		result = strconv.Itoa(input)
	}

	return result
}

func factors() []int {
	var keys []int

	for k := range convertConfig {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys
}
