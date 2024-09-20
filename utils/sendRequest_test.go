package utils

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func setupEnv() {
	// Create a mock .env file for testing
	err := os.WriteFile(".env", []byte("API_BASE_URL=http://example.com\nDEFAULT_STRING=example\nDEFAULT_INT=123\nDEFAULT_BOOL=true"), 0644)
	if err != nil {
		panic("Failed to create mock .env file")
	}
}

func teardownEnv() {
	// Clean up .env file after testing
	os.Remove(".env")
}

func TestSendRequest(t *testing.T) {
	setupEnv()
	defer teardownEnv()

	// Create a mock server to simulate API responses
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "success"}`))
	}))
	defer mockServer.Close()

	// Define test parameters
	params := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}

	// Capture log output
	var logBuf bytes.Buffer
	log.SetOutput(&logBuf)
	defer func() { log.SetOutput(os.Stderr) }() // Restore default log output

	// Call SendRequest function
	SendRequest("GET", mockServer.URL, params)

	// Print captured log output for debugging
	logOutput := logBuf.String()
	t.Logf("Captured log output:\n%s", logOutput)

	// Check if the log output contains expected content
	if !contains(logOutput, "Response from GET") {
		t.Errorf("Expected log to contain 'Response from GET', but got: %s", logOutput)
	}

	if !contains(logOutput, "Status: 200") {
		t.Errorf("Expected log to contain 'Status: 200', but got: %s", logOutput)
	}

	if !contains(logOutput, `"message": "success"`) {
		t.Errorf("Expected log to contain '\"message\": \"success\"', but got: %s", logOutput)
	}
}

// Helper function to check if a substring exists in a string
func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
