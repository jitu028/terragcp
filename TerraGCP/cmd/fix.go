package cmd

import (
	"TerraGCP/pkg/fix"
	"fmt"

	"github.com/spf13/cobra"
)

// fixCmd represents the fix command
var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix Terraform code issues",
	Long: `Fix command automatically corrects known issues in Terraform code 
           specific to Google Cloud Platform configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fix command invoked")

		// Example: Assuming args[0] is the path to the Terraform file
		if len(args) < 1 {
			fmt.Println("Error: Terraform file path is required")
			return
		}
		terraformFilePath := args[0]

		// Call the fix logic from the fix package
		err := fix.TerraformFix(terraformFilePath)
		if err != nil {
			fmt.Printf("Error fixing Terraform file: %s\n", err)
			return
		}

		fmt.Println("Terraform file fixed successfully")
	},
}

func init() {
	rootCmd.AddCommand(fixCmd)

	// Here you can define specific flags for the fix command
}
