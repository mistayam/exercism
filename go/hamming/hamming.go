// Package hamming implements a simple functon to calculate
// the hamming distance between two strands of DNA.
package hamming

import (
	"errors"
)

// Distance takes two strings, containing strands of DNA.
// The function returns the hamming distance between the two.
func Distance(a, b string) (int, error) {
	hamDist := 0

	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				hamDist++
			}
		}
	} else {
		return -1, errors.New("error: strands are not equal in length")
	}
	return hamDist, nil
}
