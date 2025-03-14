package testcases

import (
	"dsa/recursion"
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{0, 0},  // Test case 1
		{1, 1},  // Test case 2
		{2, 1},  // Test case 3
		{3, 2},  // Test case 4
		{4, 3},  // Test case 5
		{5, 5},  // Test case 6
		{6, 8},  // Test case 7
		{7, 13}, // Test case 8
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Fibonacci(%d)", testCase.input), func(t *testing.T) {
			result := recursion.Fibonacci(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %d, expected %d but got %d", testCase.input, testCase.expected, result)
			}
		})
	}
}
