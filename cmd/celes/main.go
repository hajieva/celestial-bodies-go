package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hajieva/celestial-bodies-go/internal/data"
	"github.com/hajieva/celestial-bodies-go/internal/server"
	//"github.com/spf13/cobra"
)

//go:embed planets.json moons.json
var celestialFiles embed.FS

func loadMoons() ([]data.Moon, error) {
	file, err := celestialFiles.ReadFile("moons.json")
	if err != nil {
		return nil, fmt.Errorf("reading moons.json: %w", err)
	}

	var moons []data.Moon
	decoder := json.NewDecoder(bytes.NewReader(file))
	err = decoder.Decode(&moons)
	if err != nil {
		return nil, fmt.Errorf("decoding moons.json: %w", err)
	}
	log.Printf("Moons loaded: %d", len(moons))
	return moons, nil

}
func loadPlanets() ([]data.Planet, error) {
	file, err := celestialFiles.ReadFile("planets.json")
	if err != nil {
		return nil, fmt.Errorf("reading planets.json: %w", err)
	}

	var planets []data.Planet
	decoder := json.NewDecoder(bytes.NewReader(file))

	err = decoder.Decode(&planets)
	if err != nil {
		return nil, fmt.Errorf("decoding planets.json: %w", err)
	}
	log.Printf("Planets loaded: %d\n", len(planets))
	return planets, nil
}

func main() {
	log.Println("Celestial API")
	planets, err := loadPlanets()
	if err != nil {
		log.Fatalf("Error loading planet and moon data: %v", err)
	}
	moons, err := loadMoons()
	if err != nil {
		log.Fatalf("Error loading moons: %v", err)
	}
	log.Printf("Loaded %d planets and %d moons\n", len(planets), len(moons))
	srv := server.NewServer(planets, moons)
	if err := http.ListenAndServe(":8080", srv.Handler()); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
