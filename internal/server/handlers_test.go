package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hajieva/celestial-bodies-go/internal/data"
)

func TestGetPlanets(t *testing.T) {
	planets := []data.Planet{{Name: "Earth"}, {Name: "Mars"}}
	srv := NewServer(planets, nil)

	req := httptest.NewRequest("GET", "/bodies/planets", nil)
	w := httptest.NewRecorder()
	srv.getPlanets(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	var got []data.Planet
	if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	if len(got) != 2 {
		t.Errorf("Expected 2 planets, got %d", len(got))
	}

}
func TestGetPlanetbyName(t *testing.T) {
	planets := []data.Planet{{Name: "Earth"}}
	srv := NewServer(planets, nil)

	req := httptest.NewRequest("GET", "/bodies/planets/Earth", nil)
	req.SetPathValue("name", "Earth")

	w := httptest.NewRecorder()
	srv.getPlanetbyName(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
func TestGetPlanetbyName_NotFound(t *testing.T) {
	srv := NewServer([]data.Planet{}, nil)

	req := httptest.NewRequest("GET", "/bodies/planets/Mercury", nil)
	req.SetPathValue("name", "Mercury")

	w := httptest.NewRecorder()
	srv.getPlanetbyName(w, req)
	if w.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Result().StatusCode)
	}
}
func TestGetPlanetMoons_noPlanets_emptyarray(t *testing.T) {
	srv := NewServer([]data.Planet{{Name: "Venus"}}, []data.Moon{})
	req := httptest.NewRequest("GET", "/bodies/planets/Venus/moons", nil)
	req.SetPathValue("name", "Venus")
	w := httptest.NewRecorder()
	srv.getPlanetMoons(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
	if strings.TrimSpace(w.Body.String()) == "null" {
		t.Errorf("Expected empty array, got null")
	}
}
func TestMethodControl(t *testing.T) {
	srv := NewServer(nil, nil)

	t.Run("Invalid method", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/bodies/planets", nil)
		w := httptest.NewRecorder()

		srv.Handler().ServeHTTP(w, req)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status 405, got %d", w.Code)
		}
	})
}
