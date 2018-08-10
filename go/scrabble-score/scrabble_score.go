// Package scrabble provices functions to calculate a string value
package scrabble

import (
	"strings"
)

var values = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

// Score gets a string as a param and returns
// its value according to a given value table
func Score(input string) int {
	total := 0
	for _, char := range input {
		for key, value := range values {
			convertedChar := string(char)
			if strings.Contains(key, strings.ToUpper(convertedChar)) {
				total += value
			}
		}
	}

	return total
}
