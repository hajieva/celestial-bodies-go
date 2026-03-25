package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hajieva/celestial-bodies-go/internal/data"
	"github.com/spf13/cobra"
)

func init() {
	describeCmd.AddCommand(describePlanetCmd)
	describeCmd.AddCommand(describeMoonCmd)
	rootCmd.AddCommand(describeCmd)
}

var describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Print full detials for a planet and moon",
}
var describePlanetCmd = &cobra.Command{
	Use:   "planet <name>",
	Short: "Provide details for a planet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		resp, err := http.Get("http://localhost:8080/bodies/planets/" + name)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("planet %q not found", name)
		}

		var p data.Planet
		if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
			return err
		}
		printPlanet(p)
		return nil
	},
}
var describeMoonCmd = &cobra.Command{
	Use:   "moon <name>",
	Short: "Provide details for a moon",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		resp, err := http.Get("http://localhost:8080/bodies/moons/" + name)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("moon %q not found", name)
		}
		var m data.Moon
		if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
			return err
		}
		printMoon(m)
		return nil

	},
}

func printPlanet(p data.Planet) {
	fmt.Printf("Name:              %s\n", p.Name)
	fmt.Printf("Type:              %s\n", p.Type)
	fmt.Printf("Mass:              %.2e kg\n", p.MassKg)
	fmt.Printf("Diameter:          %g km\n", p.DiameterKm)
	fmt.Printf("Distance from Sun: %.0f km\n", p.DistanceFromSunKm)
	fmt.Printf("Orbital Period:    %.2f days\n", p.OrbitalPeriodDays)
	fmt.Printf("Rotation Period:   %.2f hours\n", p.RotationPeriodHours)
	fmt.Printf("Moons:             %d\n", p.MoonsCount)
	fmt.Println("Atmosphere:")
	for _, a := range p.Atmosphere {
		fmt.Printf("  %-16s %.2f%%\n", a.Gas, a.Percentage)
	}
	fmt.Printf("Temperature:\n  Min: %.0f°C   Mean: %.0f°C   Max: %.0f°C\n",
		p.Temperature.MinCelsius, p.Temperature.MeanCelsius, p.Temperature.MaxCelsius)
	fmt.Printf("Description:\n  %s\n", p.Description)
}
func printMoon(m data.Moon) {
	fmt.Printf("Name:          %s\n", m.Name)
	fmt.Printf("Parent Planet: %s\n", m.ParentPlanet)
	fmt.Printf("Mass:          %.2e kg\n", m.MassKg)
	fmt.Printf("Diameter:      %.0f km\n", m.DiameterKm)
	fmt.Printf("Orbital Period: %.2f days\n", m.OrbitalPeriodDays)
	fmt.Printf("Description:\n  %s\n", m.Description)
}
