// Package reverse implements functions to reverse strings
package reverse

import "strings"

// String returns the inverse of a given sentence
func String(sentence string) string {
	chars := strings.Split(sentence, "")

	result := ""
	for i := len(chars) - 1; i >= 0; i-- {
		result = result + chars[i]
	}

	return result
}
