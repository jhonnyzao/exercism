// Package acronym provides function to abbreviate sentences
package acronym

import "strings"

// Abbreviate expects a sentence and returns
// a string representing its acronym
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, Split)
	result := ""
	for _, word := range words {
		result += string(string(word[0]))
	}
	return strings.ToUpper(result)
}

// Split creates multiple delimiters for string splitting
func Split(r rune) bool {
	return r == ' ' || r == '-'
}
