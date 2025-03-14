package sorting

func QuickSort(array []int, lowIndex int, highIndex int) {

	if lowIndex >= highIndex {
		return
	}

	// pick the pivot
	pivot := array[highIndex]

	leftPointer := lowIndex
	rightPointer := highIndex

	for leftPointer < rightPointer {

		for array[leftPointer] <= pivot && leftPointer < rightPointer {
			leftPointer++
		}

		for array[rightPointer] >= pivot && leftPointer < rightPointer {
			rightPointer--
		}

		swap(array, leftPointer, rightPointer)

	}

	swap(array, leftPointer, highIndex)

	//sort the left partition
	QuickSort(array, lowIndex, leftPointer-1)

	//sort the right partition
	QuickSort(array, leftPointer+1, highIndex)

}

func swap(arr []int, index1 int, index2 int) {
	arr[index1], arr[index2] = arr[index2], arr[index1]
}
