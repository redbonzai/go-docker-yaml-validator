package main

import (
	"fmt"
	"log"
	"os"

	"github.com/redbonzai/docker-compose-validate-go/internal/validator"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate [file]",
	Short: "Validate a Docker Compose file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}
		errors, err := validator.ValidateComposeFile(data)
		if err != nil {
			log.Fatalf("Validation error: %v", err)
		}
		if len(errors) > 0 {
			fmt.Println("Validation errors found:")
			for _, e := range errors {
				fmt.Println("-", e)
			}
		} else {
			fmt.Println("Docker Compose file is valid.")
		}
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
}
