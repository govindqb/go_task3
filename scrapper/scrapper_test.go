package scrapper_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go_task3/scrapper"
)

// Helper function to check if an endpoint is present in the result set
func endpointExists(endpoints []scrapper.Endpoint, method, path string) bool {
	for _, e := range endpoints {
		if e.Method == method && e.Path == path {
			return true
		}
	}
	return false
}

func TestGetEndpoints_HttpBin(t *testing.T) {
	// Mock Swagger JSON response based on httpbin.org API
	mockSwagger := scrapper.SwaggerDoc{
		Paths: map[string]scrapper.SwaggerPath{
			"/get": {
				Get: &scrapper.Operation{
					Parameters: []scrapper.Parameter{
						{In: "query", Name: "param1", Type: "string"},
					},
				},
			},
			"/post": {
				Post: &scrapper.Operation{
					Parameters: []scrapper.Parameter{
						{In: "body", Name: "bodyParam", Type: "string"},
					},
				},
			},
			"/put": {
				Put: &scrapper.Operation{
					Parameters: []scrapper.Parameter{
						{In: "body", Name: "updateData", Type: "string"},
					},
				},
			},
			"/delete": {
				Delete: &scrapper.Operation{},
			},
			"/patch": {
				Patch: &scrapper.Operation{
					Parameters: []scrapper.Parameter{
						{In: "body", Name: "patchParam", Type: "string"},
					},
				},
			},
		},
	}

	// Mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(mockSwagger)
	}))
	defer server.Close()

	// Call the function with the mock server URL
	endpoints := scrapper.GetEndpoints(server.URL)

	// Expected endpoints from httpbin.org
	expectedEndpoints := []struct {
		Method string
		Path   string
	}{
		{"GET", "/get"},
		{"POST", "/post"},
		{"PUT", "/put"},
		{"DELETE", "/delete"},
		{"PATCH", "/patch"},
	}

	// Check if all expected endpoints exist in the result
	for _, expected := range expectedEndpoints {
		if !endpointExists(endpoints, expected.Method, expected.Path) {
			t.Errorf("Expected endpoint %s %s not found", expected.Method, expected.Path)
		}
	}

}
