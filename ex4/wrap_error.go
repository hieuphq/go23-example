package main

import (
	"errors"
	"fmt"
)

var errNotFound = errors.New("not found")

func CheckData() error {
	dt := []int{}

	if len(dt) <= 0 {
		return fmt.Errorf("data error: %w", errNotFound)
	}

	return nil
}

func ValidateData() {
	err := CheckData()
	if err != nil {
		fmt.Printf("Origin error: %v\n", err)
		if errors.Is(err, errNotFound) {
			fmt.Println("not found the data")
		}
	}
}
