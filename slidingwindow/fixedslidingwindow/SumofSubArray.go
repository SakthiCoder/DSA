/*
Sum of minimum and maximum elements of all subarrays of given size K

Solve
Count Distinct elements in each window of size K

Solve

Find first negative integer in every k size window

Solve

Maximum of all subarray of size K

Solve

Maximum MEX of all subarray of size K */

package fixedslidingwindow

import "math"

/*
Maximise the number of toys that can be purchased with amount K
Last Updated : 27 Mar, 2024
Given an array consisting of the cost of toys. Given an integer K depicting the amount of money available to purchase toys. Write a program to find the maximum number of toys one can buy with the amount K.

Note: One can buy only 1 quantity of a particular toy.

Examples:

Input:  N = 10, K =  50,  cost = { 1, 12, 5, 111, 200, 1000, 10, 9, 12, 15 }
Output: 6
Explanation: Toys with amount 1, 5, 9, 10, 12, and 12  can be purchased resulting in a total amount of 49. Hence, maximum number of toys is 6.




Input: N = 7, K = 50,  cost = { 1, 12, 5, 111, 200, 1000, 10 }
Output: 4

*/

func SumofSubArray(arr []int, k int, str string) int {

	maxSum := 0
	currentSum := 0

	if str == "max" {
		maxSum = math.MinInt
	} else if str == "min" {
		maxSum = math.MaxInt
	} else {
		return -1
	}

	i := 0
	for i < len(arr) {
		currentSum += arr[i]

		if i >= k-1 {
			if currentSum > maxSum && str == "max" {
				maxSum = currentSum
			} else if currentSum < maxSum && str == "min" {
				maxSum = currentSum
			}
			currentSum -= arr[i-(k-1)]
		}
		i++
	}

	return maxSum

}
