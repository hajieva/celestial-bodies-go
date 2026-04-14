package main

import (
	"os"

	"github.com/spf13/cobra"
)

var serverAddr string
var rootCmd = &cobra.Command{
	Use:   "celes",
	Short: "API to retrieve celestial bodies information",
	Long:  "A CLI tool to query planetary and lunar data from a local API.",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&serverAddr, "server", "http://localhost:8080", "Address of the API server")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
