package main

import (
	"log"
	"net/http"

	"github.com/hajieva/celestial-bodies-go/internal/loader"
	"github.com/hajieva/celestial-bodies-go/internal/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listenCmd)
}

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Start the API server to listen for incoming requests",
	Long:  "Start the API server to listen for incoming requests on a specified port.",
	Run: func(cmd *cobra.Command, args []string) {
		startServer()
	},
}

func startServer() {
	planets, err := loader.LoadPlanets()
	if err != nil {
		log.Fatalf("Error loading planet and moon data: %v", err)
	}
	moons, err := loader.LoadMoons()
	if err != nil {
		log.Fatalf("Error loading moons: %v", err)
	}
	log.Printf("Loaded %d planets and %d moons\n", len(planets), len(moons))
	srv := server.NewServer(planets, moons) // You can pass actual data if needed
	if err := http.ListenAndServe(":8080", srv.Handler()); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
