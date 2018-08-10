// Package bob implements functions to answer for bob
package bob

import (
	"regexp"
	"strings"
)

// Hey gets a sentence as a param and returns an answer
func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")

	if remark == "" {
		return "Fine. Be that way!"
	}

	hasUpperCaseMatcher, _ := regexp.Compile("[A-Z]")
	hasLowerCaseMatcher, _ := regexp.Compile("[a-z]")

	isYeldMatcher := func(text string) bool {
		if hasUpperCaseMatcher.MatchString(text) && !hasLowerCaseMatcher.MatchString(text) {
			return true
		}
		return false
	}

	endsWithQuestionMarkMatcher, _ := regexp.Compile("\\?$")
	isYeld := isYeldMatcher(remark)
	endsWithQuestionMark := endsWithQuestionMarkMatcher.MatchString(remark)

	if isYeld && endsWithQuestionMark {
		return "Calm down, I know what I'm doing!"
	} else if isYeld {
		return "Whoa, chill out!"
	} else if endsWithQuestionMark {
		return "Sure."
	}

	return "Whatever."
}
