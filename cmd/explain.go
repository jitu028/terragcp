package cmd

import (
	"fmt"

	"github.com/jitu028/terragcp/pkg" // Updated import path
	"github.com/spf13/cobra"
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Provide an explanation for Terraform GCP code",
	Long: `The explain command is used to provide detailed explanations of Terraform GCP code.
It allows you to analyze and explain various aspects of your Terraform configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("filepath")
		saveFilePath, _ := cmd.Flags().GetString("save")

		if filePath == "" {
			fmt.Println("Error: A file path must be provided using the -f flag")
			return
		}

		if saveFilePath != "" {
			fmt.Printf("Explain called with file path: %s and save file path: %s\n", filePath, saveFilePath)
			err := pkg.AnalyzeTerraformAndSave(filePath, saveFilePath) // Updated function call
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Explanation saved to:", saveFilePath)
			}
		} else {
			fmt.Printf("Explain called with file path: %s\n", filePath)
			data, err := pkg.AnalyzeTerraform(filePath) // Updated function call
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Explanation:", data)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)

	// Flags for the explain command
	explainCmd.Flags().StringP("filepath", "f", "", "Specify the file path for explanation (required)")
	explainCmd.Flags().StringP("save", "s", "", "Specify the file path to save the explanation (optional)")
}
