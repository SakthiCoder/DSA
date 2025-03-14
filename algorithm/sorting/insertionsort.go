package sorting

func InsertionSort(inputArr []int) {
	for i := 1; i < len(inputArr); i++ {
		currentValue := inputArr[i]

		j := i - 1
		for j >= 0 && inputArr[j] > currentValue {
			inputArr[j+1] = inputArr[j]
			j--
		}
		inputArr[j+1] = currentValue
	}
}
