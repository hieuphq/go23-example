package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test function using httptest
func TestLogin(t *testing.T) {
	// Create a fake HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate the response from the external API
		response := map[string]bool{
			"authenticated": true,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Use the fake server URL in the login function
	loginURL := server.URL + "/auth"

	a := auth{
		BaseURL: loginURL,
	}
	// Perform the test
	authenticated, err := a.Login("testuser", "testpassword")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !authenticated {
		t.Errorf("Expected authentication to be true, but got false")
	}
}

func Test_auth_Login(t *testing.T) {
	validServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate the response from the external API
		response := map[string]bool{
			"authenticated": true,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}))
	defer validServer.Close()

	invalidServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate the response from the external API
		response := map[string]bool{}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	}))
	defer invalidServer.Close()

	type fields struct {
		BaseURL string
	}
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "auth success",
			fields: fields{
				BaseURL: validServer.URL,
			},
			args: args{
				username: "hieu",
				password: "123",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "auth failed",
			fields: fields{
				BaseURL: invalidServer.URL,
			},
			args: args{
				username: "",
				password: "",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := auth{
				BaseURL: tt.fields.BaseURL,
			}
			got, err := a.Login(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("auth.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("auth.Login() = %v, want %v", got, tt.want)
			}
		})
	}
}
