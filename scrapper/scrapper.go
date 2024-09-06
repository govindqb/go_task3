package scrapper

import (
	"encoding/json"
	"log"
	"net/http"
)

// Structure of a parameter in an HTTP operation
type Parameter struct {
	In   string `json:"in"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Structure of an HTTP operation in a path
type Operation struct {
	Parameters []Parameter `json:"parameters"`
}

// Structure of paths
type SwaggerPath struct {
	Get    *Operation `json:"get,omitempty"`
	Post   *Operation `json:"post,omitempty"`
	Put    *Operation `json:"put,omitempty"`
	Delete *Operation `json:"delete,omitempty"`
	Patch  *Operation `json:"patch,omitempty"`
}

// Structure of Swagger JSON response
type SwaggerDoc struct {
	Paths map[string]SwaggerPath `json:"paths"`
}

// Structure of an endpoint
type Endpoint struct {
	Method     string
	Path       string
	Parameters []Parameter
}

// GetEndpoints fetches the Swagger JSON and returns a list of endpoints
func GetEndpoints(url string) []Endpoint {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch JSON: %v", err)
	}
	defer res.Body.Close()

	var swaggerDoc SwaggerDoc
	if err := json.NewDecoder(res.Body).Decode(&swaggerDoc); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	var endpoints []Endpoint

	for path, pathItem := range swaggerDoc.Paths {
		if pathItem.Get != nil {
			endpoints = append(endpoints, Endpoint{
				Method:     "GET",
				Path:       path,
				Parameters: pathItem.Get.Parameters,
			})
		}
		if pathItem.Post != nil {
			endpoints = append(endpoints, Endpoint{
				Method:     "POST",
				Path:       path,
				Parameters: pathItem.Post.Parameters,
			})
		}
		if pathItem.Put != nil {
			endpoints = append(endpoints, Endpoint{
				Method:     "PUT",
				Path:       path,
				Parameters: pathItem.Put.Parameters,
			})
		}
		if pathItem.Delete != nil {
			endpoints = append(endpoints, Endpoint{
				Method:     "DELETE",
				Path:       path,
				Parameters: pathItem.Delete.Parameters,
			})
		}
		if pathItem.Patch != nil {
			endpoints = append(endpoints, Endpoint{
				Method:     "PATCH",
				Path:       path,
				Parameters: pathItem.Patch.Parameters,
			})
		}
	}

	return endpoints
}
