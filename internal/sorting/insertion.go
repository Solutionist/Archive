package sorting

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		currentValue := array[i]
		j := i - 1
		for j >= 0 && array[j] > currentValue {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = currentValue
	}

	return array
}
