package main

import "testing"

// Code to test
func Add(a, b int) int {
	return a + b
}

// Test function
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
