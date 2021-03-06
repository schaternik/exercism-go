package hamming

import "errors"

// Distance founds Hamming distance of two DNA codes
func Distance(a, b string) (int, error) {
	aRunes, bRunes := []rune(a), []rune(b)
	var distance int

	if len(aRunes) != len(bRunes) {
		err := errors.New("DNAs should be the same length")
		return 0, err
	}

	for i := range aRunes {
		if aRunes[i] != bRunes[i] {
			distance++
		}
	}

	return distance, nil
}
