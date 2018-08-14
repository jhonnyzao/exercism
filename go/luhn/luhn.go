// Package luhn implements functions to validate credit card inputs
package luhn

import (
	"strconv"
	"strings"
)

// Valid returns if a credit card string is valid
func Valid(number string) bool {
	number = strings.Replace(number, " ", "", -1)

	if len(number) <= 1 {
		return false
	}

	doubleDigitFlag := false
	sum := 0

	for i := len(number) - 1; i >= 0; i-- {
		curNumber, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}

		if doubleDigitFlag {
			sum += doubleDigit(curNumber)
		} else {
			sum += curNumber
		}

		doubleDigitFlag = !doubleDigitFlag
	}

	return sum%10 == 0
}

func doubleDigit(digit int) int {
	total := digit * 2
	if total > 9 {
		return total - 9
	}

	return total
}
