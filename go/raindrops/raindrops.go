// Package raindrops implements functions
// to convert numbers to strings
package raindrops

import "strconv"

// Convert recieves an integer as a param
// and returns different strings according
// to its factoring
func Convert(number int) string {
	result := ""
	if factored := number % 3; factored == 0 {
		result += "Pling"
	}

	if factored := number % 5; factored == 0 {
		result += "Plang"
	}

	if factored := number % 7; factored == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}
