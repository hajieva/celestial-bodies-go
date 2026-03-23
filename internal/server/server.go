package server

import (
	"net/http"

	"github.com/hajieva/celestial-bodies-go/internal/data"
)

// Server object holds planet moon data
// and mux is a router bridges incoming requests to handlers
type Server struct {
	router  *http.ServeMux
	planets []data.Planet
	moons   []data.Moon
}

// constructor for server
func NewServer(planets []data.Planet, moons []data.Moon) *Server {
	s := &Server{
		planets: planets,
		moons:   moons,
		router:  http.NewServeMux(),
	}
	s.routes()
	return s
}
func (s *Server) routes() {
	s.router.HandleFunc("GET /bodies/planets", s.getPlanets)
	s.router.HandleFunc("GET /bodies/planets/{name}", s.getPlanetbyName)
	s.router.HandleFunc("GET /bodies/planets/{name}/moons", s.getPlanetMoons)
	s.router.HandleFunc("GET /bodies/moons/{name}", s.getMoon)

}
func (s *Server) Handler() http.Handler {
	return (s.router)
}
