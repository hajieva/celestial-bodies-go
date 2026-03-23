package loader

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"

	"github.com/hajieva/celestial-bodies-go/internal/data"
)

//go:embed planets.json moons.json
var celestialFiles embed.FS

// LoadPlanets reads the planets.json file and decodes it into a slice of Planet structs.
func LoadPlanets() ([]data.Planet, error) {
	file, err := celestialFiles.ReadFile("planets.json")
	if err != nil {
		return nil, fmt.Errorf("reading planets.json: %w", err)
	}
	var planets []data.Planet
	if err := json.NewDecoder(bytes.NewReader(file)).Decode(&planets); err != nil {
		return nil, fmt.Errorf("decoding planets.json: %w", err)
	}
	return planets, nil
}

// LoadMoons reads the moons.json file and decodes it into a slice of Moon structs.
func LoadMoons() ([]data.Moon, error) {
	file, err := celestialFiles.ReadFile("moons.json")
	if err != nil {
		return nil, fmt.Errorf("reading moons.json: %w", err)
	}
	var moons []data.Moon
	if err := json.NewDecoder(bytes.NewReader(file)).Decode(&moons); err != nil {
		return nil, fmt.Errorf("decoding moons.json: %w", err)
	}
	return moons, nil
}
