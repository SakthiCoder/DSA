package testcases

import (
	"dsa/practice"
	"fmt"
	"testing"
)

func TestPalindrom(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"Sakthi", false},   // Test case 1
		{"RAR", true},       // Test case 2
		{"Hello", false},    // Test case 3
		{"MalayalaM", true}, // Test case 4
		{"eye", true},       // Test case 5
		{"madam", true},     // Test case 6
		{"racer", true},     // Test case 7
		{"level", true},     // Test case 8
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("IsPalindrom(%s)", testCase.input), func(t *testing.T) {
			result := practice.IsPalindrome(testCase.input)
			if result != testCase.expected {
				t.Errorf("For input %s, expected %v but got %v", testCase.input, testCase.expected, result)
			}
		})
	}
}
