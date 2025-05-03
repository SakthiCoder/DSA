package sorting

func MergeSort(inputArr []int) {

	// Base Condition
	if len(inputArr) <= 1 {
		return
	}

	inputLength := len(inputArr)

	midIndex := inputLength / 2

	leftHalf := make([]int, midIndex)

	rightHalf := make([]int, inputLength-midIndex)

	for i := 0; i < midIndex; i++ {
		leftHalf[i] = inputArr[i]
	}

	for i := midIndex; i < inputLength; i++ {
		rightHalf[i-midIndex] = inputArr[i]
	}

	MergeSort(leftHalf)
	MergeSort(rightHalf)

	merge(inputArr, leftHalf, rightHalf)

}

func merge(inputArr []int, leftHalf []int, rightHalf []int) {

	leftSize := len(leftHalf)
	rightSize := len(rightHalf)

	i, j, k := 0, 0, 0

	for i < leftSize && j < rightSize {
		if leftHalf[i] <= rightHalf[j] {
			inputArr[k] = leftHalf[i]
			i++
		} else {
			inputArr[k] = rightHalf[j]
			j++
		}
		k++
	}

	for i < leftSize {
		inputArr[k] = leftHalf[i]
		i++
		k++
	}

	for j < rightSize {
		inputArr[k] = rightHalf[j]
		j++
		k++
	}

}
