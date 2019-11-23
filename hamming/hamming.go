package hamming

import "errors"

// Distance founds Hamming distance of two DNA codes
func Distance(a, b string) (distance int, err error) {
	sampleLen := len(a)

	if sampleLen != len(b) {
		err = errors.New("DNAs should be the same length")
		return
	}

	for i := 0; i < sampleLen; i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return
}
