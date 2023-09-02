package main

import (
	"fmt"

	"github.com/Solutionist/Archive/internal/recursive"
)

func main() {
	result := recursive.FindFactorialRecursive(5)
	fmt.Println(result)

	result2 := recursive.FindFactorialRecursive(5)
	fmt.Println(result2)

	result3 := recursive.FibonacciRecursive(8)
	fmt.Println(result3)

	result4 := recursive.FibonacciIterative(8)
	fmt.Println(result4)
}
