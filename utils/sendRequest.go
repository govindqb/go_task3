package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/joho/godotenv"
)

// sendRequest sends an HTTP request based on the method, URL, and optional parameters
func SendRequest(method, baseURL string, parameters map[string]string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	// If there are parameters, append them to the URL
	if len(parameters) > 0 {
		queryParams := url.Values{}
		for key, value := range parameters {
			queryParams.Add(key, value)
		}
		baseURL = fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		log.Printf("Failed to create request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to get response: %v", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return
	}
	fmt.Printf("Response from %s %s:\nStatus: %s\nBody: %s\n\n", method, baseURL, resp.Status, string(respBody))
}
