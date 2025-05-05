package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

// Structure for JSON response
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Website string `json:"website"`
}

func main() {
	url := "http://localhost:50055/users/30"

	// Send HTTP GET request
		response, err := http.Get( url )
		if err != nil {
			log.Fatalf("HTTP request failed: %v", err )
		}
		defer response.Body.Close()

	// Read response body
		responseBody, err := io.ReadAll( response.Body )
		if err != nil {
			log.Fatalf("Reading response failed: %v", err )
		}

	// Decode JSON into struct
		var user User
		err = json.Unmarshal( responseBody, &user )
		if err != nil {
			log.Fatal(err)
		}

	// Print decoded data
		fmt.Printf("User =>\nID:%s\nName:%s\nEmail:%s\nWebsite:%s",
								user.ID,
								user.Name,
								user.Email,
								user.Website )
}