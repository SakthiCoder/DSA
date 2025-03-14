package searching

func BinarySearch(inputArray []int, TargetNumber int) int {
	left := 0
	right := len(inputArray) - 1
	for left <= right {
		midIndex := (left + right) / 2
		if inputArray[midIndex] == TargetNumber {
			return midIndex
		} else if inputArray[midIndex] > TargetNumber {
			right = midIndex - 1
		} else if inputArray[midIndex] < TargetNumber {
			left = midIndex + 1
		}
	}
	return -1
}
