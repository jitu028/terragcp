package pkg

import (
	"fmt"
	"os"
)

// AnalyzeTerraform reads and analyzes a Terraform file.
func AnalyzeTerraform(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read Terraform file: %v", err)
	}

	// Process the file content (e.g., analyze, generate explanations)
	analysisResult, err := GenerateContent(string(content))
	if err != nil {
		return "", fmt.Errorf("analysis failed: %v", err)
	}

	return analysisResult, nil
}

// AnalyzeTerraformAndSave reads, analyzes, and saves the analysis of a Terraform file.
func AnalyzeTerraformAndSave(filePath, savePath string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read Terraform file: %v", err)
	}

	err = GenerateContentAndSave(string(content), savePath)
	if err != nil {
		return fmt.Errorf("failed to analyze and save: %v", err)
	}

	return nil
}
