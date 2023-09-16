package sorting

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	middle := len(array) / 2

	return merge(MergeSort(array[:middle]), MergeSort(array[middle:]))
}

func merge(i1, i2 []int) []int {
	result := make([]int, len(i1)+len(i2))

	for i, j, k := 0, 0, 0; k < len(result); k++ {
		if i >= len(i1) {
			result[k] = i2[j]
			j++
			continue
		} else if j >= len(i2) {
			result[k] = i1[i]
			i++
			continue
		}

		if i1[i] < i2[j] {
			result[k] = i1[i]
			i++
		} else {
			result[k] = i2[j]
			j++
		}
	}

	return result
}
