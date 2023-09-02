package recursive

func FindFactorialRecursive(number int) int {
	if number <= 2 {
		return factorialChecker(number)
	}

	return number * FindFactorialRecursive(number-1)
}

func FindFactorialIterative(number int) int {
	if number <= 2 {
		return factorialChecker(number)
	}

	for i := (number - 1); i > 1; i-- {
		number *= i
	}

	return number
}

func factorialChecker(number int) int {
	switch number {
	case 2:
		return 2
	default:
		return 1
	}
}
