// Package diffsquares implements functions
// to return the differences of squares
package diffsquares

// SquareOfSums returns the square of the sum of the first n natural numbers
func SquareOfSums(n int) int {
	sum, _ := returnBothSums(n)

	return sum
}

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	_, sum := returnBothSums(n)

	return sum
}

// Difference returns the differente between the square of sums and the sum of the squares
func Difference(n int) int {
	squareOfSums := SquareOfSums(n)
	sumOfSquares := SumOfSquares(n)

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
