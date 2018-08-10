// Package hamming implements routines for 
// measuring differences between given DNAs
package hamming

import (
	"errors"
)

// Distance returns the total of different elements in DNAs
func Distance(a, b string) (int, error) {
	total := 0

	if len(a) != len(b) {
		return total, errors.New("DNAs have different sizes")
	}

	for i := range a {
		if a[i] != b[i] {
			total++
		}
	}

	return total, nil
}
