package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/hajieva/celestial-bodies-go/internal/data"
)

var planets []data.Planet
var moons []data.Moon

func loadMoons() []data.Moon {
	file, err := os.Open("moons.json")
	if err != nil {
		fmt.Println("Error opening moons.json:", err)
		return nil
	}

	defer file.Close()
	var moons []data.Moon
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&moons)
	if err != nil {
		fmt.Println("Error decoding moons.json:", err)
		return nil
	}
	fmt.Printf("Moons loaded: %d\n", len(moons))
	return moons

}
func loadPlanets() []data.Planet {
	file, err := os.Open("planets.json")
	if err != nil {
		fmt.Println("Error opening planets.json", err)
		return nil
	}
	defer file.Close()

	var planets []data.Planet
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&planets) // why store in memory?
	if err != nil {
		fmt.Println("Error decoding planets.json", err)
		return nil
	}
	fmt.Printf("Planets loaded: %d\n", len(planets))
	return planets
}

func main() {
	fmt.Println("Celestial API")
	planets = loadPlanets()
	moons = loadMoons()

	fmt.Printf("Loaded %d planets and %d moons\n", len(planets), len(moons))

	http.HandleFunc("/bodies/planets", getPlanets)
	http.HandleFunc("/bodies/planets/{name}", getPlanet)
	http.HandleFunc("/bodies/planets/{name}/moons", getPlanetMoons)
	http.HandleFunc("/bodies/moons/{name}", getMoon)

	http.ListenAndServe(":8080", nil)
}

func getPlanets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(planets)
}
func getPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// how to get planet name from URL?
	planetName := r.PathValue("name")

	for _, planet := range planets {
		if strings.EqualFold(planet.Name, planetName) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(planet)
			return
		}
	}
	http.Error(w, "Planet not found", http.StatusNotFound)

}

var moonCount int
var counter int

func getPlanetMoons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// specific planet's moon details
	// find the planet first
	// find that planet's moons
	planetName := r.PathValue("name")

	for _, planet := range planets {

		if strings.EqualFold(planetName, planet.Name) {

			moonCount = planet.MoonsCount

			for _, moons := range moons {
				if strings.EqualFold(moons.ParentPlanet, planet.Name) {
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(moons)
					if counter == moonCount {
						return
					}
					counter++
				}
			}
		}
	}
	http.Error(w, "Planet or moons not found", http.StatusNotFound)

}
func getMoon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	moonName := r.PathValue("name")
	for _, moon := range moons {
		if strings.EqualFold(moon.Name, moonName) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(moon)
			return
		}
	}
	http.Error(w, "Moon not found", http.StatusNotFound)
}
