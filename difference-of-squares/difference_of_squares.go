package diffsquares

func SquareOfSum(n int) int {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	sqOfSum := 0
	sqOfSum = sum * sum

	return sqOfSum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}

	return sum
}

func Difference(n int) int {
	total := SquareOfSum(n) - SumOfSquares(n)

	return total
}
