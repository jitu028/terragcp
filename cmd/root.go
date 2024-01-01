package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "terragcp-cli",
	Short: "CLI tool for managing Terraform configurations in GCP",
	Long: `TerraGCP is a Command Line Interface tool designed to facilitate the management,
analysis, and modification of Terraform configurations specifically for Google Cloud Platform.
For example:

This application can analyze, explain, and suggest fixes for your Terraform GCP configurations.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	// Here you can define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Example of a global flag
	// rootCmd.PersistentFlags().StringVarP(&someConfigVar, "config", "c", "", "config file (default is $HOME/.TerraGCP.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
