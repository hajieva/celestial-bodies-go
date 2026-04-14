package server

import (
	"net/http"

	"github.com/hajieva/celestial-bodies-go/internal/data"
)

// Server represents an HTTP server for celestial bodies data.
type Server struct {
	planets []data.Planet
	moons   []data.Moon
	router  *http.ServeMux
}

// NewServer creates and returns a new Server instance with the provided planets and moons.
func NewServer(planets []data.Planet, moons []data.Moon) *Server {
	s := &Server{
		planets: planets,
		moons:   moons,
		router:  http.NewServeMux(),
	}
	s.routes()
	return s
}

// routes registers all HTTP route handlers for the server.
func (s *Server) routes() {
	s.router.HandleFunc("GET /bodies/planets", s.getPlanets)
	s.router.HandleFunc("GET /bodies/planets/{name}", s.getPlanetbyName)
	s.router.HandleFunc("GET /bodies/planets/{name}/moons", s.getPlanetMoons)
	s.router.HandleFunc("GET /bodies/moons/{name}", s.getMoon)

}

// Handler returns the HTTP handler for the server.
func (s *Server) Handler() http.Handler {
	return (s.router)
}
