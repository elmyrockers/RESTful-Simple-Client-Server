package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
)

type UserResponse struct {
	ID  string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Website string `json:"website"`
}

func main() {
	http.HandleFunc("/users/", handleUserRequest)
	fmt.Println( "HTTP server is listening on port 50055" )
	http.ListenAndServe(":50055", nil)
}

// handleUserRequest processes GET requests
func handleUserRequest( responseWriter http.ResponseWriter, request *http.Request) {
	// Allow only GET method
		if request.Method != http.MethodGet {
			http.Error( responseWriter, "Method Not Allowed", http.StatusMethodNotAllowed )
			return
		}

	// Request data
		userID := strings.TrimPrefix(request.URL.Path, "/users/")

	// Send response back
		user := UserResponse{
			ID:       userID,
			Name:     "Helmi Aziz",
			Email: "helmi@xeno.com.my",
			Website:   "https://elmyrockers.github.io",
		}

		responseWriter.Header().Set("Content-Type", "application/json")
		json.NewEncoder( responseWriter ).Encode( user )
}