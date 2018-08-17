// Package grains implements functions to count grains
package grains

import "errors"

// Square returns the square of a number in a chess table
func Square(amount int) (uint64, error) {
	if amount <= 0 || amount > 64 {
		return 0, errors.New("invalid amount")
	}

	return 1 << uint64(amount-1), nil
}

// Total returns the total sum of the squares
func Total() uint64 {
	var total uint64

	for i := 0; i <= 64; i++ {
		aux, err := Square(i)
		if err != nil {
			panic("you shouldnt change the code, babe")
		}

		total += aux
	}

	return total
}
