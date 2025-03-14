package testcases

import (
	"dsa/algorithm/sorting"
	"fmt"
	"testing"
)

func TestBoubleSort(t *testing.T) {

	arr1 := []int{2, 123, 2, 32, 2, 3, 23}
	testCase := []struct {
		input    []int
		expected []int
	}{
		{arr1, []int{}},                                    // Test case 1
		{[]int{1}, []int{1}},                               // Test case 2
		{[]int{2, 1}, []int{1, 2}},                         // Test case 3
		{[]int{3, 2}, []int{2, 3}},                         // Test case 4
		{[]int{4, 3}, []int{3, 4}},                         // Test case 5
		{[]int{5, 5}, []int{5, 5}},                         // Test case 6
		{[]int{6, 8}, []int{6, 8}},                         // Test case 7
		{[]int{7, 13, 2, 5}, []int{2, 5, 7, 13}},           // Test case 8
		{[]int{10, 1, 3, 7, 9}, []int{1, 3, 7, 9, 10}},     // Test case 9
		{[]int{20, 15, 10, 5}, []int{5, 4, 15, 20}},        // Test case 10
		{[]int{9, 3, 1, 7, 5, 2}, []int{1, 2, 3, 5, 7, 9}}, // Test case 11
		{[]int{12, 8, 14, 10}, []int{8, 10, 12, 14}},       // Test case 12
		{[]int{17, 5, 10, 13, 7}, []int{5, 7, 10, 13, 17}}, // Test case 13
	}

	for _, testCase := range testCase {
		t.Run(fmt.Sprintf("Fibonacci(%d)", testCase.input), func(t *testing.T) {
			sorting.BoubleSort(testCase.input)
			// if result != testCase.expected {
			// 	t.Errorf("For input %d, expected %d but got %d", testCase.input, testCase.expected, result)
			// }
		})
	}
}
