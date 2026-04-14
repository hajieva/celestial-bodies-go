package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hajieva/celestial-bodies-go/internal/data"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(planetsCmd)
}

var planetsCmd = &cobra.Command{
	Use:   "planets",
	Short: "Get information about planets in the solar system",
	Long:  "Retrieve detailed information about the planets in our solar system, including their names, diameters, and distances from the sun.",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := fmt.Sprintf("%s/bodies/planets", serverAddr)
		resp, err := http.Get(url)
		if err != nil {
			return fmt.Errorf("failed to fetch planets from %s: %w", url, err)
		}
		defer resp.Body.Close()

		var planets []data.Planet

		if err := json.NewDecoder(resp.Body).Decode(&planets); err != nil {
			return err
		}
		for _, planet := range planets {
			fmt.Printf("Name: %s\n", planet.Name)
		}
		return nil
	},
}
