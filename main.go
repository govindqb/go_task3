package main

import (
	"fmt"
	"go_task3/scrapper"
	"io"
	"log"
	"net/http"
	"net/url"
)

// sendRequest sends an HTTP request based on the method, URL, and optional parameters
func sendRequest(method, baseURL string, parameters map[string]string) {
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
	url := "https://httpbin.org/spec.json"
	endpoints := scrapper.GetEndpoints(url)

	for _, endpoint := range endpoints {
		path := fmt.Sprintf("https://httpbin.org%s", endpoint.Path)

		// Create a parameters map if the endpoint has parameters
		params := make(map[string]string)
		for _, param := range endpoint.Parameters {
			paramType := param.Type
			switch paramType {
			case "string":
				params[param.Name] = "test_value"
			case "int":
				params[param.Name] = "1234"
			case "bool":
				params[param.Name] = "true"
			default:
				params[param.Name] = "test_value"

			}
		}

		// Call sendRequest with method, path, and parameters
		sendRequest(endpoint.Method, path, params)
	}
}
