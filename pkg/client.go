package pkg

import (
	"fmt"
	"os"
)

// GCPClient struct holds the client configuration and context
type GCPClient struct {
	// Add fields as needed for GCP client
}

// NewGCPClient creates a new client to interact with GCP
func NewGCPClient() (*GCPClient, error) {
	// Initialize GCP client
	// ...

	return &GCPClient{
		// Initialize fields
	}, nil
}

// SetGeminiAPIKey sets the Gemini API key as an environment variable
func SetGeminiAPIKey(apiKey string) error {
	key := "GEMINI_API_KEY"
	err := os.Setenv(key, apiKey)
	if err != nil {
		return fmt.Errorf("failed to set Gemini API key: %v", err)
	}
	return nil
}

// GetGeminiAPIKey retrieves the Gemini API key from an environment variable
func GetGeminiAPIKey() (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("gemini api key not set")
	}
	return apiKey, nil
}
