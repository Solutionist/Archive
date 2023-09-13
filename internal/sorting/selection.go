package sorting

func SelectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		smallest := array[i]
		for j := 0; j < len(array)-1; j++ {
			if smallest < array[j] {
				smallest, array[j] = array[j], smallest
			}
		}
		array[i] = smallest
	}

	return array
}
