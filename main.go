package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var planets = []Planets{
	{Name: "Mercury", Mass: 0.9889, Satellites: []string{}},
	{Name: "Mars", Mass: 0.64717, Satellites: []string{"Phobos", "Deimos"}},
	{Name: "Earth", Mass: 1.0, Satellites: []string{"Moon"}},
}

func main() {
	fmt.Println("Celestial API")
	http.HandleFunc("/planets", getPlanets)

	http.ListenAndServe(":8080", nil)
}

type Planets struct {
	Name       string
	Mass       float64
	Satellites []string
}

func getPlanets(w http.ResponseWriter, r *http.Request) {
	//mars := Planets{Name: "Mars",
	//Mass:       0.64171,
	//Satellites: []string{"Phobos", "Deimos"}}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(planets)
}
