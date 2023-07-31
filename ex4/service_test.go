package main

import (
	"reflect"
	"testing"
)

func TestGetUserByID(t *testing.T) {
	// Create a mock database
	mockDB := &MockDatabase{
		users: map[int]*User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}

	// Create the UserService with the mockDB
	userService := NewUserService(mockDB)

	// Test the GetUserByID method
	user, err := userService.GetUserByID(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedUser := &User{ID: 1, Name: "Alice"}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("Expected user %+v, but got %+v", expectedUser, user)
	}
}
