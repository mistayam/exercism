// Package hamming implements a simple functon to calculate
// the hamming distance between two strands of DNA.
package hamming

import (
	"errors"
)

// Distance takes two strings, containing strands of DNA.
// The function returns the hamming distance between the two.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("error: strands are not equal in length")
	}

	hamDist := 0

	for i := range a {
		if a[i] != b[i] {
			hamDist++
		}
	}

	return hamDist, nil
}
