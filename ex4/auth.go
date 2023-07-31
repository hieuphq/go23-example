package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type auth struct {
	BaseURL string
}

// Function to test
func (a auth) Login(username, password string) (bool, error) {
	// Prepare the request payload
	data := map[string]string{
		"username": username,
		"password": password,
	}

	// Convert data to JSON
	requestBody, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	// Send the POST request to the login endpoint
	resp, err := http.Post(a.BaseURL+"/auth", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("failed")
	}

	// Parse the response JSON
	var response map[string]bool
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return false, err
	}

	return response["authenticated"], nil
}
