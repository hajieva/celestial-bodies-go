package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hajieva/celestial-bodies-go/internal/data"
)

// getPlanets retrieves all planets and returns them as JSON.
func (s *Server) getPlanets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(s.planets); err != nil {
		http.Error(w, "Failed to encode planets", http.StatusInternalServerError)
	}
}

// getPlanetbyName retrieves a planet by name and returns it as JSON.
func (s *Server) getPlanetbyName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.PathValue("name")

	for _, p := range s.planets {
		if strings.EqualFold(p.Name, name) {
			if err := json.NewEncoder(w).Encode(p); err != nil {
				http.Error(w, "Failed to encode planet", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Planet not found", http.StatusNotFound)
}

// getPlanetMoons retrieves all moons for a specific planet and returns them as JSON.
func (s *Server) getPlanetMoons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	planetName := r.PathValue("name")
	for _, planet := range s.planets {
		if strings.EqualFold(planetName, planet.Name) {
			moons := []data.Moon{}
			for _, moon := range s.moons {
				if strings.EqualFold(moon.ParentPlanet, planet.Name) {
					moons = append(moons, moon)
				}
			}
			if err := json.NewEncoder(w).Encode(moons); err != nil {
				http.Error(w, "Failed to encode moons", http.StatusInternalServerError)
			}
			return
		}

	}

	http.Error(w, "Planet  not found", http.StatusNotFound)
}

// getMoon retrieves a moon by name and returns it as JSON.
func (s *Server) getMoon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	moonName := r.PathValue("name")
	for _, moon := range s.moons {
		if strings.EqualFold(moon.Name, moonName) {
			if err := json.NewEncoder(w).Encode(moon); err != nil {
				http.Error(w, "Failed to encode moon", http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Moon not found", http.StatusNotFound)
}
