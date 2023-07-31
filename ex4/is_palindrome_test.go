package main

import (
	"fmt"
	"testing"
)

// Test function with table-driven tests and subtests
func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"radar", true},   // Palindrome
		{"level", true},   // Palindrome
		{"hello", false},  // Not a palindrome
		{"deified", true}, // Palindrome
		{"golang", false}, // Not a palindrome
		{"", true},        // Empty string (considered a palindrome)
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %s", tc.input), func(t *testing.T) {
			result := IsPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("For input '%s', expected %t, but got %t", tc.input, tc.expected, result)
			}
		})
	}
}
