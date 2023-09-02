package recursive

// O(2^n)
func FibonacciRecursive(number int) int {
	if number <= 2 {
		return fibonacciChecker(number)
	}

	return FibonacciRecursive(number-1) + FibonacciRecursive(number-2)
}

// O(n)
func FibonacciIterative(number int) int {
	if number <= 2 {
		return fibonacciChecker(number)
	}

	sequence := []int{0, 1, 1}
	for i := 3; i <= number; i++ {
		result := sequence[i-1] + sequence[i-2]
		sequence = append(sequence, result)
	}

	return sequence[number]
}

func fibonacciChecker(number int) int {
	switch number {
	case 0:
		return 0
	default:
		return 1
	}
}
