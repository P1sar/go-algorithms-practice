package diffsquares



func SquareOfSum(number int) int {
	squareOfSum := 0
	for i := 0; i <= number; i++ {
		squareOfSum += i
	}
	return squareOfSum * squareOfSum
}

func SumOfSquares(number int) int {
	sumOfSquares := 0
	for i := 0; i <= number; i++ {
		sumOfSquares += i * i
	}
	return sumOfSquares
}

func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}

