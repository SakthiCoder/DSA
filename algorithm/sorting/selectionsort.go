package sorting

func SelectionSort(inputArr []int) {

	for i := 0; i < len(inputArr); i++ {
		minIndex := i
		for j := i + 1; j < len(inputArr); j++ {
			if inputArr[minIndex] > inputArr[j] {
				minIndex = j
			}
		}
		swap(inputArr, minIndex, i)
	}
}
