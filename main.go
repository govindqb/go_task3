package main

import (
	"fmt"
	"go_task3/scrapper"
	"go_task3/utils"
	"log"
	"os"

	"github.com/joho/godotenv"
)

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
		utils.SendRequest(endpoint.Method, path, params)
	}
}
