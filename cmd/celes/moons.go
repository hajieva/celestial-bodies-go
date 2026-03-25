package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hajieva/celestial-bodies-go/internal/data"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(moonsCmd)
}

var moonsCmd = &cobra.Command{
	Use:   "moons <planet>",
	Short: "List moons of a planet",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		planet := args[0]
		resp, err := http.Get("http://localhost:8080/bodies/planets/" + planet + "/moons")
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var moons []data.Moon
		if err := json.NewDecoder(resp.Body).Decode(&moons); err != nil {
			return err
		}
		if len(moons) == 0 {
			fmt.Printf("%s has no moons, \n", planet)
			return nil
		}
		for _, m := range moons {
			fmt.Println(m.Name)
		}
		return nil
	},
}
