package main

import "fmt"

type MockDatabase struct {
	users map[int]*User
}

func (mdb *MockDatabase) GetUserByID(userID int) (*User, error) {
	if user, ok := mdb.users[userID]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (mdb *MockDatabase) SaveUser(user *User) error {
	if mdb.users == nil {
		mdb.users = make(map[int]*User)
	}
	mdb.users[user.ID] = user
	return nil
}

// Initialize and use MockDatabase in tests...
