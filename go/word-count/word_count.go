// Package wordcount implements functions to count words
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency maps words as keys and quantities as values
type Frequency map[string]int

// WordCount returns the count of the words in a sentence
func WordCount(sentence string) Frequency {
	sentence = strings.ToLower(sentence)
	regexp, _ := regexp.Compile("[^ a-z0-9]+,-;")
	sentence = regexp.ReplaceAllString(sentence, "")
	words := strings.FieldsFunc(sentence, Split)

	var result Frequency
	result = make(map[string]int)
	for _, word := range words {
		result[word]++
	}

	return result
}

func Split(r rune) bool {
	return r == ' ' || r == ',' || r == '-' || r == ';' || r == '\n'
}
