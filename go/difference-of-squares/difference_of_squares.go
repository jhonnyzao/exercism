// Package diffsquares implements functions
// to return the differences of squares
package diffsquares

// SquareOfSums returns the square of the sum of the first n natural numbers
func SquareOfSums(n int) int {
	total := 0

	for i := n; i > 0; i-- {
		total += i
	}

	total *= total

	return total
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	total := 0

	for i := n; i > 0; i-- {
		total += i * i
	}

	return total
}

// Difference returns the differente between the square of sums and the sum of the squares
func Difference(n int) int {
	squareOfSums, sumOfSquares := returnBothSums(n)

	return squareOfSums - sumOfSquares
}

func returnBothSums(n int) (int, int) {
	squareOfSums := 0
	sumOfSquares := 0

	for i := n; i > 0; i-- {
		squareOfSums += i
		sumOfSquares += i * i
	}

	squareOfSums *= squareOfSums

	return squareOfSums, sumOfSquares
}
