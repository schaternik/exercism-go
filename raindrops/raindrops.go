package raindrops

import "strconv"

var convertConfig = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert returns string depending on input's factor
func Convert(input int) string {
	var result string

	for k, v := range convertConfig {
		if input%k == 0 {
			result += v
		}
	}

	if result == "" {
		result = strconv.Itoa(input)
	}

	return result
}
