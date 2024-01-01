package main

import (
	"fmt"
	"os"

	"github.com/jitu028/terragcp/cmd" // Corrected import path
)

func main() {
	if err := cmd.Execute(); err != nil { // Corrected method call
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
