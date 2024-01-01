package cmd

import (
	"TerraGCP/pkg/gcp"
	"fmt"

	"github.com/spf13/cobra"
)

// listResourcesCmd represents the command to list GCP resources
var listResourcesCmd = &cobra.Command{
	Use:   "list-resources",
	Short: "List GCP resources in Terraform",
	Long: `List resources command scans the Terraform configuration files and 
           lists the GCP resources defined in them.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing GCP resources in Terraform configuration...")

		// Assuming args[0] is the directory path of Terraform files
		if len(args) < 1 {
			fmt.Println("Error: Terraform configuration directory path is required")
			return
		}
		terraformConfigDir := args[0]

		// Call the function to list resources from the gcp package
		resources, err := gcp.ListGCPResources(terraformConfigDir)
		if err != nil {
			fmt.Printf("Error listing GCP resources: %s\n", err)
			return
		}

		fmt.Println("GCP Resources found in Terraform configuration:")
		for _, resource := range resources {
			fmt.Println(resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(listResourcesCmd)

	// Here you can define specific flags for the list-resources command
}
