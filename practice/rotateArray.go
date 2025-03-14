package practice

func RotateArray(inputArray []int, n int) {
	// Get the length of the array
	length := len(inputArray)

	// If the array is empty or the rotation count is 0, no need to do anything
	if length == 0 || n == 0 {
		return
	}

	// Normalize n to be within the bounds of the array length
	k := n % length

	if k < 0 {
		k += length
	}

	// Reverse the entire array
	reverse(inputArray, 0, length-1)

	// Reverse the first part (0 to k-1)
	reverse(inputArray, 0, k-1)

	// Reverse the second part (k to length-1)
	reverse(inputArray, k, length-1)
}

func reverse(arr []int, start int, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
