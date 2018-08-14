// Package strain implements functions to keep or
// discard elements given a list of numbers as a param
package strain

// Ints is a slice of ints
type Ints []int

// Lists is a slice of slice of ints
type Lists [][]int

// Strings is a slice of strings
type Strings []string

// Keep returns a list of ints to be kept
// according to the given predicate
func (list Ints) Keep(predicate func(int) bool) Ints {
	var result Ints
	for _, value := range list {
		if predicate(value) {
			result = append(result, value)
		}
	}

	return result
}

// Discard returns a list of ints to be discarded
// according to the given predicate
func (list Ints) Discard(predicate func(int) bool) Ints {
	var result Ints
	for _, value := range list {
		if !predicate(value) {
			result = append(result, value)
		}
	}

	return result
}

// Keep returns a list of ints to be kept
// according to the given predicate
func (list Strings) Keep(predicate func(string) bool) Strings {
	var result Strings
	for _, value := range list {
		if predicate(value) {
			result = append(result, value)
		}
	}

	return result
}

// Discard returns a list of ints to be discarded
// according to the given predicate
func (list Strings) Discard(predicate func(string) bool) Strings {
	var result Strings
	for _, value := range list {
		if !predicate(value) {
			result = append(result, value)
		}
	}

	return result
}

// Keep returns a list of lists to be kept
// according to the given predicate
func (list Lists) Keep(predicate func([]int) bool) Lists {
	var result Lists
	for _, value := range list {
		if predicate(value) {
			result = append(result, value)
		}
	}

	return result
}

// Discard returns a list of lists to be discarted
// according to the given predicate
func (list Lists) Discard(predicate func([]int) bool) Lists {
	var result Lists
	for _, value := range list {
		if !predicate(value) {
			result = append(result, value)
		}
	}

	return result
}
