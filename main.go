package main

import (
	"fmt"
	"go_task3/scrapper"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// sendRequest sends an HTTP request based on the method, URL, and optional parameters
func sendRequest(method, baseURL string, parameters map[string]string) {
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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	apiBaseURL := os.Getenv("API_BASE_URL")
	defaultString := os.Getenv("DEFAULT_STRING")
	defaultInt := os.Getenv("DEFAULT_INT")
	defaultBool := os.Getenv("DEFAULT_BOOL")

	url := fmt.Sprintf("%s/spec.json", apiBaseURL)
	endpoints := scrapper.GetEndpoints(url)

	for _, endpoint := range endpoints {
		path := fmt.Sprintf("%s%s", apiBaseURL, endpoint.Path)

		// Create a parameters map if the endpoint has parameters
		params := make(map[string]string)
		for _, param := range endpoint.Parameters {
			switch param.Type {
			case "string":
				params[param.Name] = defaultString
			case "int":
				params[param.Name] = defaultInt
			case "bool":
				params[param.Name] = defaultBool
			default:
				params[param.Name] = defaultString
			}
		}

		// Call sendRequest with method, path, and parameters
		sendRequest(endpoint.Method, path, params)
	}
}
