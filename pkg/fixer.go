package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const gptAPIURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent"

type RequestBody struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"Parts"`
	} `json:"contents"`
}

type ResponseBody struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
		FinishReason  string `json:"finishReason"`
		Index         int    `json:"index"`
		SafetyRatings []struct {
			Category    string `json:"category"`
			Probability string `json:"probability"`
		} `json:"safetyRatings"`
	} `json:"candidates"`
	PromptFeedback struct {
		SafetyRatings []struct {
			Category    string `json:"category"`
			Probability string `json:"probability"`
		} `json:"safetyRatings"`
	} `json:"promptFeedback"`
}

// ApplyFixes suggests fixes for the given Terraform file using Gemini Pro API.
func ApplyFixes(filePath string) ([]string, error) {
	// Read the Terraform file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read Terraform file: %v", err)
	}

	// Call the Gemini Pro API to get suggestions
	suggestions, err := callGeminiProAPI(string(fileContent))
	if err != nil {
		return nil, err
	}

	// Process the suggestions and return them
	// This part needs to be implemented based on how you want to handle the API response
	return processSuggestions(suggestions), nil
}

func callGeminiProAPI(inputText string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("Gemini API key not set")
	}

	// Prepare the request body
	requestBody := RequestBody{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"Parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{
						Text: inputText,
					},
				},
			},
		},
	}

	// Convert the request body to JSON
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	// Make the API request
	resp, err := http.Post(fmt.Sprintf(gptAPIURL, apiKey), "application/json", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", fmt.Errorf("failed to make API request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	// Parse the response JSON
	var response ResponseBody
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	// Extract and return the generated text
	if len(response.Candidates) > 0 && len(response.Candidates[0].Content.Parts) > 0 {
		return response.Candidates[0].Content.Parts[0].Text, nil
	}

	return "", fmt.Errorf("no generated text found in response")
}

func processSuggestions(apiResponse string) []string {
	// Assuming the API response is a JSON string that contains an array of suggestions
	// First, unmarshal the response into a suitable Go struct
	var response struct {
		Suggestions []struct {
			Text string `json:"text"`
		} `json:"suggestions"`
	}

	err := json.Unmarshal([]byte(apiResponse), &response)
	if err != nil {
		fmt.Printf("Error processing suggestions: %v\n", err)
		return nil
	}

	// Extract suggestions from the response
	var suggestions []string
	for _, suggestion := range response.Suggestions {
		suggestions = append(suggestions, suggestion.Text)
	}

	return suggestions
}
