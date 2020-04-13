package scale

import "strings"

var sharpScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flatScale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
var intervals = map[rune]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

// Scale generates scale based on tonic and interval
func Scale(tonic, interval string) []string {
	usedScale := []string{}

	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		usedScale = append(usedScale, sharpScale...)
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		usedScale = append(usedScale, flatScale...)
	default:
		return []string{}
	}

	i := indexOf(strings.Title(tonic), usedScale)
	usedScale = append(usedScale[i:], usedScale[:i]...)

	return usedScale
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}
