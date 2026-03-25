package main

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "celes",
	Short: "API to retrieve celestial bodies information",
	Long:  "A CLI tool to query planetary and lunar data from a local API.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
