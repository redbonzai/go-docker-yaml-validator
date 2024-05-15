package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "docker-compose-validate",
	Short: "Docker Compose Validator",
	Long:  "A tool to validate Docker Compose files.",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
