package main

import (
	"fmt"

	"github.com/Solutionist/Archive/internal/recursive"
	"github.com/Solutionist/Archive/internal/sorting"
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

	result5 := sorting.BubbleSort([]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0})
	fmt.Println(result5)

	result6 := sorting.SelectionSort([]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0})
	fmt.Println(result6)

	result7 := sorting.InsertionSort([]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0})
	fmt.Println(result7)

	result8 := sorting.MergeSort([]int{99, 44, 6, 2, 1, 5, 63, 87, 283, 4, 0})
	fmt.Println(result8)
}
