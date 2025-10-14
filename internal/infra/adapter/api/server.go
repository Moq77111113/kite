package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moq77111113/kite/internal/application/template"
)

// Server handles HTTP requests for the Kite API and web UI
type Server struct {
	service *template.Service
	router  *mux.Router
}

// New creates a new Server instance
func New(svc *template.Service) *Server {
	s := &Server{
		service: svc,
	}
	s.setupRoutes()
	return s
}

// setupRoutes configures all HTTP routes
func (s *Server) setupRoutes() {
	s.router = mux.NewRouter()

	api := s.router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/templates", s.listTemplates).Methods("GET")
	api.HandleFunc("/templates/{name}", s.getTemplate).Methods("GET")

	s.router.HandleFunc("/", s.healthCheck).Methods("GET")
}

func (s *Server) Start(port string) error {
	addr := ":" + port
	return http.ListenAndServe(addr, s.router)
}
