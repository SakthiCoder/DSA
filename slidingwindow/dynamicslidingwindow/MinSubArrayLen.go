package dynamicslidingwindow

import "fmt"

func MinSubArrayLen(target int, inputArr []int) int {

	left, sum, minLength := 0, 0, len(inputArr)+1

	for right := 0; right < len(inputArr); right++ {
		sum += inputArr[right]

		for sum >= target {
			minLength = min(minLength, right-left+1)

			sum -= inputArr[left]

			left++
		}

	}

	if minLength == len(inputArr)+1 {
		return 0
	}

	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func LongestSubarrayWithKDistinct(nums []int, k int) int {
	// Map to store the frequency of each element in the window
	count := make(map[int]int)
	left, maxLength := 0, 0

	for right := 0; right < len(nums); right++ {
		// Add the current element to the map
		count[nums[right]]++
		fmt.Println("RIGHT: ", count)

		// Shrink the window until there are no more than K distinct elements
		for len(count) > k {
			fmt.Println("LEFT: ", count)
			count[nums[left]]--
			if count[nums[left]] == 0 {
				delete(count, nums[left])
			}
			left++
		}

		fmt.Println("After Shrink: ", count)
		// Update the maximum length of the subarray
		maxLength = max(maxLength, right-left+1)
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
