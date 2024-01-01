package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// geminiConfCmd represents the command to configure Gemini API settings
var geminiConfCmd = &cobra.Command{
	Use:   "gemini-conf",
	Short: "Configure Gemini API settings",
	Long: `The gemini-conf command is used to manage configuration settings for the Gemini API.
It allows you to set and modify the Gemini API key used for accessing Gemini services.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, _ := cmd.Flags().GetString("apikey")

		if apiKey != "" {
			err := setGeminiAPIKey(apiKey)
			if err != nil {
				fmt.Printf("Failed to set Gemini API key: %v\n", err)
			} else {
				fmt.Println("Gemini API key configured successfully")
			}
		} else {
			fmt.Println("Please provide the Gemini API key using the --apikey flag")
		}
	},
}

func init() {
	rootCmd.AddCommand(geminiConfCmd)

	// Flags for the gemini-conf command
	geminiConfCmd.Flags().StringP("apikey", "k", "", "Specify the Gemini API key (required)")
}

// setGeminiAPIKey sets the Gemini API key as an environment variable
func setGeminiAPIKey(apiKey string) error {
	key := "GEMINI_API_KEY"
	err := os.Setenv(key, apiKey)
	if err != nil {
		return fmt.Errorf("failed to set environment variable: %v", err)
	}
	return nil
}
