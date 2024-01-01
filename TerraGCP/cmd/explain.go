package cmd

import (
	"TerraGCP/pkg/analysis"
	"fmt"

	"github.com/spf13/cobra"
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain Terraform code",
	Long: `Explain command provides detailed explanations of Terraform code 
           configurations for Google Cloud Platform.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Explain command invoked")

		// Example: Assuming args[0] is the path to the Terraform file
		if len(args) < 1 {
			fmt.Println("Error: Terraform file path is required")
			return
		}
		terraformFilePath := args[0]

		// Call the explain logic from the analysis package
		explanation, err := analysis.ExplainTerraform(terraformFilePath)
		if err != nil {
			fmt.Printf("Error explaining Terraform file: %s\n", err)
			return
		}

		fmt.Println("Terraform Explanation:")
		fmt.Println(explanation)
	},
}

func init() {
	rootCmd.AddCommand(explainCmd)

	// Here you can define specific flags for the explain command
}
