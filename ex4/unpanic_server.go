package main

import (
	"fmt"
	"net/http"
)

// panicMiddleware recovers from panics and prevents the server from crashing
func panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("Recovered from panic:", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func StartServer() {
	// Create a simple HTTP server
	mux := http.NewServeMux()

	// Attach the panicMiddleware to the default handler
	mux.Handle("/", panicMiddleware(http.HandlerFunc(handleRequest)))
	http.HandleFunc("/auth/login", AuthHandler)

	// Start the server on port 8080
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", mux)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Simulate a panic for demonstration purposes
	if r.URL.Path == "/panic" {
		panic("Something went wrong!")
	}

	// Your normal request handling logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, this is a normal response."))
}

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	// Handle the request for Google OAuth2 login
	// Redirect the user to Google login page
	// (Code for redirecting to Google login is not shown in this example)
}
