package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the initialization command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the TerraGCP tool",
	Long: `Initialize command sets up the necessary configurations and environment 
           for TerraGCP to interact with Terraform and Google Cloud Platform.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initializing TerraGCP...")

		// Add your initialization logic here
		// This could include setting up config files, environment variables, etc.

		fmt.Println("TerraGCP initialized successfully.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you can define specific flags for the init command
}
