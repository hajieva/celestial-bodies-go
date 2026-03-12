package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"github.com/hajieva/celestial-bodies-go/internal/data"
)
func loadMoons() []data.Moon{
	file, err := os.Open("moons.json")
	if err != nil{
		fmt.Println("Error opening moons.json:", err)
		return nil
	}

	defer file.Close()

	var moons [] data.Moon
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&moons)
	if err != nil{
		fmt.Println("Error decoding moons.json:", err)
		return nil
	}
	fmt.Printf("Moons loaded: %d\n", len(moons))
	return moons

}
func loadPlanets() []data.Planet {
	file, err := os.Open("planets.json")
	if err != nil{
		fmt.Println("Error opening planets.json", err)
		return nil
	}
	defer file.Close()
	
	var planets []data.Planet
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&planets)
	if err != nil{
		fmt.Println("Error decoding planets.json", err)
		return nil
	}
	fmt.Printf("Planets loaded: %d\n", len(planets))
	return planets
}

func main() {
	fmt.Println("Celestial API")
	
	moons := loadMoons()
	planets := loadPlanets()
	fmt.Printf("Loaded %d planets and %d moons\n", len(planets), len(moons))

	
	//http.HandleFunc("/planets", get Planets)

	//http.ListenAndServe(":8080", nil)
}



func getPlanets(w http.ResponseWriter, r *http.Request) {
	//mars := Planets{Name: "Mars",
	//Mass:       0.64171,
	//Satellites: []string{"Phobos", "Deimos"}}

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(planets)
}
