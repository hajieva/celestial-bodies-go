package main

import (
	"log"
	"net/http"

	"github.com/hajieva/celestial-bodies-go/internal/loader"
	"github.com/hajieva/celestial-bodies-go/internal/server"
	"github.com/spf13/cobra"
)

var port string

func init() {
	listenCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to listen on")
	rootCmd.AddCommand(listenCmd)
}

var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Start the API server to listen for incoming requests",
	Long:  "Start the API server to listen for incoming requests on a specified port.",
	RunE: func(cmd *cobra.Command, args []string) error {
		planets, err := loader.LoadPlanets()
		if err != nil {
			return err
		}
		moons, err := loader.LoadMoons()
		if err != nil {
			return err
		}
		log.Printf("Loaded %d planets and %d moons\n", len(planets), len(moons))
		srv := server.NewServer(planets, moons)
		if err := http.ListenAndServe(":"+port, srv.Handler()); err != nil {
			return err
		}
		return nil
	},
}
