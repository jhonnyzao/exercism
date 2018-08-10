// Package isogram implements functions to
// determine if a sentence is an isogram
package isogram

import (
	"sort"
	"strings"
)

// IsIsogram returns a bool saying if a
// given sentence is an isogram
func IsIsogram(sentence string) bool {
	result := true

	sentence = strings.ToLower(sentence)
	chars := strings.Split(sentence, "")
	sort.Strings(chars)

	for i := 0; i < len(chars)-1; i++ {
		currentChar := chars[i]
		nextChar := chars[i+1]
		if currentChar == nextChar &&
			!isRepeatableChar(currentChar) {
			result = false
		}
	}

	return result
}

func isRepeatableChar(char string) bool {
	return char == " " || char == "-"
}
