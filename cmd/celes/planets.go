package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	Run: func(cmd *cobra.Command, args []string) {
		fetchAllPlanets()
	},
}

func fetchAllPlanets() {
	resp, err := http.Get("http://localhost:8080/bodies/planets")
	if err != nil {
		log.Fatalf("Error fetching planet data: %v", err)
	}
	defer resp.Body.Close() //why this line is needed?

	//log.Printf("Response status: %s", req.Status)
	var planets []data.Planet

	if err := json.NewDecoder(resp.Body).Decode(&planets); err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	for _, planet := range planets {

		fmt.Printf("Name: %s\n", planet.Name)

	}

}
