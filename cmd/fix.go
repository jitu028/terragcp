package cmd

import (
	"fmt"

	"github.com/jitu028/terragcp/pkg" // Updated import path

	"github.com/spf13/cobra"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix Terraform code issues",
	Long: `The fix command is used to automatically correct known issues in Terraform code 
specific to Google Cloud Platform configurations using the Gemini Pro API.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("filepath")

		if filePath == "" {
			fmt.Println("Error: Terraform file path is required")
			return
		}

		fmt.Printf("Fixing Terraform file: %s\n", filePath)
		fixes, err := pkg.ApplyFixes(filePath) // Updated function call
		if err != nil {
			fmt.Printf("Error occurred while fixing Terraform file: %s\n", err)
			return
		}

		if len(fixes) == 0 {
			fmt.Println("No fixes were necessary.")
		} else {
			fmt.Println("Applied fixes:")
			for _, f := range fixes {
				fmt.Println(f)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)

	// Flags for the fix command
	fixCmd.Flags().StringP("filepath", "f", "", "Specify the file path of the Terraform file to fix")
}
