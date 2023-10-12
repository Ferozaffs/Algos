package MergeSortIP

func merge(list *[]int, l int, m int, r int) {
	if (*list)[m] < (*list)[m+1] {
		return
	}

	leftIndex := l
	midIndex := m + 1
	rightIndex := midIndex
	for leftIndex < midIndex && rightIndex < r+1 {
		if (*list)[leftIndex] < (*list)[rightIndex] {
			leftIndex++
		} else {
			tmp := (*list)[rightIndex]
			for i := rightIndex; i > leftIndex; i-- {
				(*list)[i] = (*list)[i-1]
			}
			(*list)[leftIndex] = tmp

			leftIndex++
			midIndex++
			rightIndex++
		}
	}
}

func mergeSort(list *[]int, l int, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(list, l, m)
		mergeSort(list, m+1, r)
		merge(list, l, m, r)
	}
}

func Sort(list *[]int) {
	mergeSort(list, 0, len(*list)-1)
}
