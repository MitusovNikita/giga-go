package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// User represents a simple user struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-memory store for users (slice)
var users []User
var nextID int = 1        // Incremental ID for new users
var mutex = &sync.Mutex{} // To ensure thread-safe operations on users slice

// handler function for the /users route
func usersHandler(w http.ResponseWriter, r *http.Request) {
	// Handle GET request for creating a new user with query params
	// http://localhost:61635/users?create=true&name=John%20Doe&email=john.doe@example.com
	if r.Method == "GET" && r.URL.Query().Get("create") != "" {
		// Parse query parameters for name and email
		name := r.URL.Query().Get("name")
		email := r.URL.Query().Get("email")

		// Validate that Name and Email fields are provided
		if name == "" || email == "" {
			http.Error(w, "Name and Email are required", http.StatusBadRequest)
			return
		}

		// Ensure thread-safe access to users slice using a mutex
		mutex.Lock()

		// Create a new user and append to the users slice
		newUser := User{
			ID:    nextID,
			Name:  name,
			Email: email,
		}
		nextID++
		users = append(users, newUser)

		mutex.Unlock()

		// Respond with the newly created user in JSON format
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
		return
	}

	// Handle GET request for fetching users
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	}

	// If the method is not GET, return a method not allowed response
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	// Define the /users route and its handler function
	http.HandleFunc("/users", usersHandler)

	// Start the server on port 8085
	log.Println("Server started on :8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
